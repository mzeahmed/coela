// Package project defines Project, the data model produced by a stack's
// interactive wizard and consumed by internal/scaffold. It holds no
// business logic: it is the plain vocabulary shared between the wizard and
// the scaffold phase.
package project

import (
	"fmt"
	"strings"
)

// Project holds the answers collected by a stack's Wizard func. It is the
// only input the scaffold engine receives to render a stack's templates.
type Project struct {
	// Name is the project name, used both as the target directory and as
	// template data (e.g. in README.md.tmpl).
	Name string
	// Stack identifies which framework/CMS this project targets.
	Stack Stack
	// PHPVersion is the PHP version selected in the wizard (e.g. "8.4").
	PHPVersion string
	// Database is the database engine selected in the wizard.
	Database Database
	// Redis reports whether the Redis service should be enabled.
	Redis bool
	// Mailpit reports whether the Mailpit service should be enabled.
	Mailpit bool
	// Traefik reports whether the Traefik reverse proxy should be enabled.
	Traefik bool
}

// String returns a human-readable summary of the project, e.g. for display
// right after the wizard completes.
func (p Project) String() string {
	var b strings.Builder

	b.WriteString("Project\n\n")
	fmt.Fprintf(&b, "Name       : %s\n", p.Name)
	fmt.Fprintf(&b, "Stack      : %s\n", p.Stack)
	fmt.Fprintf(&b, "PHP        : %s\n", p.PHPVersion)
	fmt.Fprintf(&b, "Database   : %s\n", p.Database)
	fmt.Fprintf(&b, "Redis      : %s\n", yesNo(p.Redis))
	fmt.Fprintf(&b, "Mailpit    : %s\n", yesNo(p.Mailpit))
	fmt.Fprintf(&b, "Traefik    : %s\n", yesNo(p.Traefik))

	return b.String()
}

// yesNo renders a boolean as the "Yes"/"No" label used in Project.String.
func yesNo(v bool) string {
	if v {
		return "Yes"
	}

	return "No"
}
