package heroicons

import (
	"html/template"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmpl(t *testing.T) {
	tmpl, err := Tmpl()
	assert.Nil(t, err)
	assert.Nil(t, tmpl.ExecuteTemplate(os.Stdout, "outline/arrow-right", nil))
}

func TestExtend(t *testing.T) {
	parent, err := template.New("parent").Parse(`{{define "hai"}}hai{{end}}`)
	assert.Nil(t, err)
	assert.Nil(t, Extend(parent))
	assert.Nil(t, parent.ExecuteTemplate(os.Stdout, "heroicons/outline/arrow-right", nil))
	assert.Nil(t, parent.ExecuteTemplate(os.Stdout, "hai", nil))
}

func TestIconsByName(t *testing.T) {
	icons := Icons{}
	_, err := icons.ByName("outline/arrow-right")
	assert.Nil(t, err)
}
