// Package scaffold renders a stack's template filesystem into a real
// project on disk. It knows nothing about any specific stack (Symfony,
// WordPress, ...): it only receives a *project.Project and an fs.FS to
// walk.
package scaffold

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mzeahmed/coela/internal/project"
)

// Generate walks templates, renders every *.tmpl file it finds against p
// (using text/template, with p itself as the template data), and writes the
// result under a directory named after p.Name. Directory structure is
// mirrored from templates, minus the .tmpl suffix on file names.
//
// templates is an fs.FS rather than a disk path because a stack's templates
// are embedded into the coela binary (see each stack's TemplatesDir), so
// `coela new` works standalone without the source repo present on disk.
func Generate(p *project.Project, templates fs.FS) error {
	root := p.Name

	return fs.WalkDir(templates, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Only *.tmpl files produce output; directories and any other
		// file (e.g. .gitkeep) are skipped.
		if d.IsDir() || !strings.HasSuffix(path, ".tmpl") {
			return nil
		}

		// fs.FS paths are always "/"-separated; convert to the OS's
		// separator before joining with a real filesystem path.
		// e.g. "docker/php/Dockerfile.tmpl" -> "<project>/docker/php/Dockerfile"
		target := filepath.Join(root, filepath.FromSlash(strings.TrimSuffix(path, ".tmpl")))

		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}

		tmpl, err := template.ParseFS(templates, path)
		if err != nil {
			return err
		}

		out, err := os.Create(target)
		if err != nil {
			return err
		}
		defer out.Close()

		return tmpl.Execute(out, p)
	})
}
