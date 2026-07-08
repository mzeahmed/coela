// Package scaffold renders a stack's template directory into a real project
// on disk. It knows nothing about any specific stack (Symfony, WordPress,
// ...): it only receives a *project.Project and a path to walk.
package scaffold

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mzeahmed/coela/internal/project"
)

// Generate walks templatesDir, renders every *.tmpl file it finds against p
// (using text/template, with p itself as the template data), and writes the
// result under a directory named after p.Name. Directory structure is
// mirrored from templatesDir, minus the .tmpl suffix on file names.
func Generate(p *project.Project, templatesDir string) error {
	root := p.Name

	return filepath.WalkDir(templatesDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Only *.tmpl files produce output; directories and any other
		// file (e.g. .gitkeep) are skipped.
		if d.IsDir() || !strings.HasSuffix(path, ".tmpl") {
			return nil
		}

		relPath, err := filepath.Rel(templatesDir, path)
		if err != nil {
			return err
		}

		// e.g. "docker/php/Dockerfile.tmpl" -> "<project>/docker/php/Dockerfile"
		target := filepath.Join(root, strings.TrimSuffix(relPath, ".tmpl"))

		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}

		tmpl, err := template.ParseFiles(path)
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
