// Package symfony implements the Symfony stack: it asks the Symfony-specific
// wizard questions, exposes its template directory, and installs the
// framework via Composer. It knows nothing about any other stack.
package symfony

import (
	"fmt"

	"github.com/mzeahmed/coela/internal/project"
	"github.com/mzeahmed/coela/internal/ui"
)

// Wizard interactively asks the user for a project name, PHP version,
// database engine, and which optional services (Redis, Mailpit, Traefik)
// to enable, then returns a fully-populated *project.Project for a Symfony
// project.
func Wizard() (*project.Project, error) {
	name, err := ui.Input("Project name")
	if err != nil {
		return nil, err
	}

	phpVersion, err := ui.Select("PHP Version", []string{"8.4", "8.3"})
	if err != nil {
		return nil, err
	}

	databases := map[string]project.Database{
		"MariaDB":    project.DatabaseMariaDB,
		"MySQL":      project.DatabaseMySQL,
		"PostgreSQL": project.DatabasePostgres,
	}

	databaseLabel, err := ui.Select("Database", []string{"MariaDB", "MySQL", "PostgreSQL"})
	if err != nil {
		return nil, err
	}

	database, ok := databases[databaseLabel]
	if !ok {
		return nil, fmt.Errorf("unknown database: %s", databaseLabel)
	}

	redis, err := ui.Confirm("Use Redis")
	if err != nil {
		return nil, err
	}

	mailpit, err := ui.Confirm("Use Mailpit")
	if err != nil {
		return nil, err
	}

	traefik, err := ui.Confirm("Use Traefik")
	if err != nil {
		return nil, err
	}

	return &project.Project{
		Name:       name,
		Stack:      project.StackSymfony,
		PHPVersion: phpVersion,
		Database:   database,
		Redis:      redis,
		Mailpit:    mailpit,
		Traefik:    traefik,
	}, nil
}
