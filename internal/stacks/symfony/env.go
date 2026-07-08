package symfony

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mzeahmed/coela/internal/project"
)

// ConfigureEnv points the freshly installed Symfony application at the
// Docker services Coela just scaffolded: the database connection, the mail
// transport, and the default URI used to generate URLs outside an HTTP
// request (e.g. CLI commands).
//
// Every value goes into app/.env.local, never app/.env. app/.env comes from
// `composer create-project` and its content depends on whatever recipes the
// installed symfony/skeleton version ships (which may already define its
// own DEFAULT_URI, for example) — appending to it risks duplicating a block
// Symfony already wrote. app/.env.local is gitignored, always safe to
// create fresh, and — per Symfony's own convention documented at the top of
// app/.env — takes precedence over app/.env, so these values are guaranteed
// to win regardless of what the installed skeleton already defines.
func ConfigureEnv(p *project.Project) error {
	appDir := filepath.Join(p.Name, "app")

	return appendToFile(filepath.Join(appDir, ".env.local"), envBlock(p)+databaseURLLine(p))
}

func envBlock(p *project.Project) string {
	mailerDSN := "null://null"
	if p.Mailpit {
		mailerDSN = "smtp://mailpit:1025"
	}

	defaultURI := "http://localhost"
	if p.Traefik {
		defaultURI = fmt.Sprintf("https://%s.local", p.Name)
	}

	return fmt.Sprintf(`###> symfony/routing ###
# Configure how to generate URLs in non-HTTP contexts, such as CLI commands.
# See https://symfony.com/doc/current/routing.html#generating-urls-in-commands
DEFAULT_URI=%s
###< symfony/routing ###

###> symfony/mailer ###
MAILER_DSN=%s
###< symfony/mailer ###

NOREPLY_EMAIL=noreply@%s.local

`, defaultURI, mailerDSN, p.Name)
}

// databaseURLLine builds the DATABASE_URL matching the "database" service
// docker-compose.yml.tmpl renders for p.Database, using the same
// credentials as the project root's .env (DB_DATABASE/DB_USER/DB_PASSWORD,
// all p.Name) that docker-compose.yml reads for the database container.
func databaseURLLine(p *project.Project) string {
	switch p.Database {
	case project.DatabasePostgres:
		return fmt.Sprintf(
			"DATABASE_URL=\"postgresql://%s:%s@database:5432/%s?serverVersion=17&charset=utf8\"\n",
			p.Name, p.Name, p.Name,
		)
	case project.DatabaseMySQL:
		return fmt.Sprintf(
			"DATABASE_URL=\"mysql://%s:%s@database:3306/%s?serverVersion=8.4.0&charset=utf8mb4\"\n",
			p.Name, p.Name, p.Name,
		)
	default:
		return fmt.Sprintf(
			"DATABASE_URL=\"mysql://%s:%s@database:3306/%s?serverVersion=11.8.8-MariaDB&charset=utf8mb4\"\n",
			p.Name, p.Name, p.Name,
		)
	}
}

func appendToFile(path, content string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)

	return err
}