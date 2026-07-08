package symfony

import (
	"embed"
	"io/fs"
)

// templatesFS embeds this stack's templates into the coela binary so `coela
// new` works from a plain downloaded binary, without the source repo's
// assets/ directory present on disk.
//
//go:embed all:assets
var templatesFS embed.FS

// TemplatesDir returns the filesystem scaffold.Generate should walk to
// render a Symfony project.
func TemplatesDir() fs.FS {
	sub, err := fs.Sub(templatesFS, "assets")
	if err != nil {
		panic(err)
	}

	return sub
}
