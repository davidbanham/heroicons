package heroicons

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmpl(t *testing.T) {
	tmpl, err := Tmpl()
	assert.Nil(t, err)
	assert.Nil(t, tmpl.ExecuteTemplate(os.Stdout, "heroicons/outline/arrow-right", nil))
}

func TestIconsByName(t *testing.T) {
	icons := Icons{}
	_, err := icons.ByName("outline/arrow-right")
	assert.Nil(t, err)
}
