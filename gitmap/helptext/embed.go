package helptext

import "embed"

//go:embed all:../help/*.md
var Files embed.FS
