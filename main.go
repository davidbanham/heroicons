package heroicons

import (
	"fmt"
	"io/fs"
	"strings"
	"text/template"
)

var pathMap = map[string]string{
	"outline": "upstream/src/24/outline",
	"solid":   "upstream/src/24/solid",
	"mini":    "upstream/src/20/solid",
}

func Tmpl() (*template.Template, error) {
	files := Files()

	t := template.New("heroicons")

	for prefix, dir := range pathMap {
		if err := slurp(files, dir, "heroicons/"+prefix, t); err != nil {
			return nil, err
		}
	}

	return t, nil
}

type Icons struct {
	cache fs.FS
}

func (this *Icons) ByName(name string) (string, error) {
	if this.cache == nil {
		this.cache = Files()
	}

	parts := strings.Split(name, "/")
	if len(parts) != 2 {
		return "", fmt.Errorf("Invalid name")
	}

	d, err := fs.ReadFile(this.cache, pathMap[parts[0]]+"/"+parts[1]+".svg")
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func slurp(files fs.FS, dir, prefix string, t *template.Template) error {
	icons, err := fs.ReadDir(files, dir)
	if err != nil {
		return err
	}
	for _, file := range icons {
		if file.IsDir() {
			continue
		}
		d, err := fs.ReadFile(files, dir+"/"+file.Name())
		if err != nil {
			return err
		}
		name := prefix + "/" + strings.ReplaceAll(file.Name(), ".svg", "")
		tmpl := fmt.Sprintf(`{{define "%s"}}%s{{end}}`, name, string(d))
		if _, err := t.New(name).Parse(tmpl); err != nil {
			return err
		}
	}
	return nil
}

func Files() fs.FS {
	return icons
}
