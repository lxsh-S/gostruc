package templates

import "embed"

//go:embed go/*.tmpl cpp/*.tmpl ts/*.tmpl
var FS embed.FS
