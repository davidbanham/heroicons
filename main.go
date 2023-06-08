package heroicons

import (
	"fmt"
	"html/template"
	"io/fs"
	"regexp"
	"strings"
)

var pathMap = map[string]string{
	"outline": "upstream/src/24/outline",
	"solid":   "upstream/src/24/solid",
	"mini":    "upstream/src/20/solid",
}

type Options struct {
	FixedColor      bool
	FixedDimensions bool
}

func Tmpl(opts ...Options) (*template.Template, error) {
	t := template.New("heroicons")

	opt := Options{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	err := extend(t, "", opt)

	return t, err
}

func extend(t *template.Template, preface string, opt Options) error {
	files := Files()

	for prefix, dir := range pathMap {
		if err := slurp(files, dir, preface+prefix, t, opt); err != nil {
			return err
		}
	}

	return nil
}

func Extend(t *template.Template, opts ...Options) error {
	opt := Options{}
	if len(opts) > 0 {
		opt = opts[0]
	}
	return extend(t, "heroicons/", opt)
}

type Icons struct {
	cache fs.FS
}

func (this *Icons) ByName(name string, opts ...Options) (string, error) {
	if this.cache == nil {
		this.cache = Files()
	}

	opt := Options{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	parts := strings.Split(name, "/")
	if len(parts) != 2 {
		return "", fmt.Errorf("Invalid name")
	}

	d, err := fs.ReadFile(this.cache, pathMap[parts[0]]+"/"+parts[1]+".svg")
	if err != nil {
		return "", err
	}

	icon := applyOptions(d, opt)

	return icon, nil
}

func slurp(files fs.FS, dir, prefix string, t *template.Template, opt Options) error {
	icons, err := fs.ReadDir(files, dir)
	if err != nil {
		return err
	}
	for _, file := range icons {
		if file.IsDir() {
			continue
		}
		b, err := fs.ReadFile(files, dir+"/"+file.Name())
		if err != nil {
			return err
		}
		name := prefix + "/" + strings.ReplaceAll(file.Name(), ".svg", "")

		data := applyOptions(b, opt)

		tmpl := fmt.Sprintf(`{{define "%s"}}%s{{end}}`, name, data)
		if _, err := t.New(name).Parse(tmpl); err != nil {
			return err
		}
	}
	return nil
}

var dimensions = regexp.MustCompile(`((width|height)\="\d\d" ){2}`)
var dimensionNumber = regexp.MustCompile(`\d\d`)

var fixedColorCode = regexp.MustCompile(`#0F172A`)

var dimensionReplacements = map[string][]byte{
	"20": []byte(`class="w-5 h-5" `),
	"24": []byte(`class="w-6 h-6" `),
}

func applyOptions(icon []byte, opt Options) string {
	if !opt.FixedColor {
		icon = fixedColorCode.ReplaceAll(icon, []byte("currentColor"))
	}
	if !opt.FixedDimensions {
		dimension := dimensionNumber.Find(dimensions.Find(icon))
		classes := dimensionReplacements[string(dimension)]
		icon = dimensions.ReplaceAll(icon, classes)
	}
	return string(icon)
}

func Files() fs.FS {
	return icons
}
