package heroicons

import (
	"embed"
	_ "embed"
)

//go:embed upstream/src/**/**/*.svg
var icons embed.FS
