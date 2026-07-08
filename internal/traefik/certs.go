package traefik

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mzeahmed/coela/internal/project"
)

// CertFile and KeyFile are the filenames GenerateCert writes into
// <project>/certs, and the exact paths traefik/dynamic.yml.tmpl expects
// under /certs inside the Traefik container. Keep the two in sync.
func CertFile(p *project.Project) string { return p.Name + ".local.pem" }
func KeyFile(p *project.Project) string  { return p.Name + ".local-key.pem" }

// GenerateCert trusts the local mkcert CA (idempotent) and generates a TLS
// certificate covering p.Name+".local" and its subdomains into
// <p.Name>/certs, ready for traefik/dynamic.yml.tmpl to pick up.
func GenerateCert(p *project.Project) error {
	if _, err := exec.LookPath("mkcert"); err != nil {
		return fmt.Errorf("mkcert is required to generate local HTTPS certificates but was not found in PATH: %w", err)
	}

	install := exec.Command("mkcert", "-install")
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr

	if err := install.Run(); err != nil {
		return err
	}

	certsDir := filepath.Join(p.Name, "certs")
	if err := os.MkdirAll(certsDir, 0755); err != nil {
		return err
	}

	generate := exec.Command("mkcert",
		"-cert-file", filepath.Join(certsDir, CertFile(p)),
		"-key-file", filepath.Join(certsDir, KeyFile(p)),
		p.Name+".local",
		"*."+p.Name+".local",
	)
	generate.Stdout = os.Stdout
	generate.Stderr = os.Stderr

	return generate.Run()
}
