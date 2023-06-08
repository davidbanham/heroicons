package heroicons

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmpl(t *testing.T) {
	tmpl, err := Tmpl()
	assert.Nil(t, err)
	buf := bytes.Buffer{}
	assert.Nil(t, tmpl.ExecuteTemplate(&buf, "outline/arrow-right", nil))
	t.Log(buf.String())
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
