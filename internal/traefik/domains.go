// Package traefik automates the local HTTPS setup (mkcert certificates and
// /etc/hosts entries) that a Traefik-fronted project needs. It knows
// nothing about any specific stack: it only receives a *project.Project.
package traefik

import "github.com/mzeahmed/coela/internal/project"

// Domains returns the local domains routed through Traefik for p, mirroring
// the Host() rules rendered in traefik/dynamic.yml.tmpl: the app itself,
// Mailpit if enabled, and phpMyAdmin (always present).
func Domains(p *project.Project) []string {
	domains := []string{p.Name + ".local"}

	if p.Mailpit {
		domains = append(domains, "mail."+p.Name+".local")
	}

	domains = append(domains, "db."+p.Name+".local")

	return domains
}
