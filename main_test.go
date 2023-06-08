package heroicons

import (
	"bytes"
	"html/template"
	"io/fs"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestTmpl(t *testing.T) {
	tmpl, err := Tmpl()
	assert.Nil(t, err)
	buf := bytes.Buffer{}
	assert.Nil(t, tmpl.ExecuteTemplate(&buf, "outline/arrow-right", nil))
	t.Log(buf.String())

	uniq := uuid.NewV4().String()

	assert.Nil(t, tmpl.ExecuteTemplate(&buf, "outline/arrow-right", map[string]string{
		"Classes": uniq,
	}))
	assert.Contains(t, buf.String(), ` class="`+uniq+`" `)
}

func TestExtend(t *testing.T) {
	parent, err := template.New("parent").Parse(`{{define "hai"}}hai{{end}}`)
	assert.Nil(t, err)
	assert.Nil(t, Extend(parent))
	buf := bytes.Buffer{}
	assert.Nil(t, parent.ExecuteTemplate(&buf, "heroicons/outline/arrow-right", nil))
	assert.Nil(t, parent.ExecuteTemplate(&buf, "hai", nil))
	t.Log(buf.String())
}

func TestIconsByName(t *testing.T) {
	icons := Icons{}
	_, err := icons.ByName("outline/arrow-right")
	assert.Nil(t, err)
}

func TestFixedColor(t *testing.T) {
	icons := Icons{}
	data, err := icons.ByName("outline/arrow-right")
	assert.Nil(t, err)
	assert.Contains(t, data, "currentColor")

	data2, err := icons.ByName("outline/arrow-right", Options{FixedColor: true})
	assert.Nil(t, err)
	assert.NotContains(t, data2, "currentColor")

	tmpl, err := Tmpl(Options{FixedColor: true})
	assert.Nil(t, err)
	buf := bytes.Buffer{}
	assert.Nil(t, tmpl.ExecuteTemplate(&buf, "outline/arrow-right", nil))
	assert.NotContains(t, buf.String(), "currentColor")

	tmpl2, err := Tmpl(Options{FixedColor: true})
	assert.Nil(t, err)
	buf2 := bytes.Buffer{}
	assert.Nil(t, tmpl2.ExecuteTemplate(&buf2, "outline/arrow-right", nil))
	assert.NotContains(t, buf2.String(), "currentColor")
}

func TestDimensions(t *testing.T) {
	files := Files()
	target := "upstream/src/24/outline/academic-cap.svg"
	d, err := fs.ReadFile(files, target)
	assert.Nil(t, err)
	assert.True(t, dimensions.Match(d))
	assert.Equal(t, 1, len(dimensions.FindAll(d, -1)))
}

func TestFixedDimensions(t *testing.T) {
	icons := Icons{}
	data, err := icons.ByName("outline/arrow-right")
	assert.Nil(t, err)
	assert.Contains(t, data, `class="w-6 h-6"`)
	assert.NotContains(t, data, `width="24"`)
	assert.NotContains(t, data, `height="24"`)

	data2, err := icons.ByName("outline/arrow-right", Options{FixedDimensions: true})
	assert.Nil(t, err)
	assert.Contains(t, data2, `width="24"`)
	assert.Contains(t, data2, `height="24"`)
	assert.NotContains(t, data2, `class="w-6 h-6"`)

	data3, err := icons.ByName("mini/arrow-right")
	assert.Nil(t, err)
	assert.Contains(t, data3, `class="w-5 h-5"`)
	assert.NotContains(t, data3, `width="24"`)
	assert.NotContains(t, data3, `height="24"`)

	data4, err := icons.ByName("mini/arrow-right", Options{FixedDimensions: true})
	assert.Nil(t, err)
	assert.NotContains(t, data4, `class="w-5 h-5"`)
	assert.Contains(t, data4, `width="20"`)
	assert.Contains(t, data4, `height="20"`)
}
