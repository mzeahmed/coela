package wordpress

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mzeahmed/coela/internal/project"
)

// ConfigureEnv turns the app/.env.example that `composer create-project
// roots/bedrock` ships into a working app/.env: it fills in the database
// DSN matching the "database" service docker-compose.yml renders, the site
// URL matching whether Traefik is enabled, and a fresh set of WordPress
// secret keys/salts (Bedrock ships every project with the same
// "generateme" placeholders, so leaving them as-is would mean every
// generated project shares the same secrets).
func ConfigureEnv(p *project.Project) error {
	appDir := filepath.Join(p.Name, "app")

	example, err := os.ReadFile(filepath.Join(appDir, ".env.example"))
	if err != nil {
		return err
	}

	salts, err := saltOverrides()
	if err != nil {
		return err
	}

	overrides := append([]envOverride{
		{"DATABASE_URL", "'" + databaseURL(p) + "'"},
		{"WP_ENV", "development"},
		{"WP_HOME", wpHome(p)},
		{"WP_SITEURL", `"${WP_HOME}/wp"`},
	}, salts...)

	env := applyOverrides(string(example), overrides)

	// DATABASE_URL supersedes DB_NAME/DB_USER/DB_PASSWORD: comment them out
	// rather than leaving them set alongside the DSN, per .env.example's own
	// "When using a DSN, you can remove the DB_NAME, DB_USER, DB_PASSWORD,
	// and DB_HOST variables" note.
	env = commentOutLines(env, []string{"DB_NAME", "DB_USER", "DB_PASSWORD"})

	return os.WriteFile(filepath.Join(appDir, ".env"), []byte(env), 0644)
}

// commentOutLines prefixes with "# " every not-already-commented line
// assigning one of keys.
func commentOutLines(env string, keys []string) string {
	lines := strings.Split(env, "\n")

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		for _, key := range keys {
			if strings.HasPrefix(trimmed, key+"=") {
				lines[i] = "# " + line

				break
			}
		}
	}

	return strings.Join(lines, "\n")
}

func wpHome(p *project.Project) string {
	if p.Traefik {
		return fmt.Sprintf("https://%s.local", p.Name)
	}

	return "http://localhost"
}

// databaseURL builds the DSN matching the "database" service
// docker-compose.yml.tmpl renders for p.Database, using the same
// credentials as the project root's .env (DB_DATABASE/DB_USER/DB_PASSWORD,
// all p.Name) that docker-compose.yml reads for the database container.
func databaseURL(p *project.Project) string {
	if p.Database == project.DatabasePostgres {
		return fmt.Sprintf("pgsql://%s:%s@database:5432/%s", p.Name, p.Name, p.Name)
	}

	return fmt.Sprintf("mysql://%s:%s@database:3306/%s", p.Name, p.Name, p.Name)
}

// envOverride is a single KEY=value pair to force into a .env file.
type envOverride struct {
	key   string
	value string
}

// saltOverrides generates a fresh, unique value for each of Bedrock's
// WordPress secret keys/salts, in place of the "generateme" placeholders
// app/.env.example ships with.
func saltOverrides() ([]envOverride, error) {
	keys := []string{
		"AUTH_KEY", "SECURE_AUTH_KEY", "LOGGED_IN_KEY", "NONCE_KEY",
		"AUTH_SALT", "SECURE_AUTH_SALT", "LOGGED_IN_SALT", "NONCE_SALT",
	}

	overrides := make([]envOverride, len(keys))

	for i, key := range keys {
		salt, err := randomSalt()
		if err != nil {
			return nil, err
		}

		overrides[i] = envOverride{key, "'" + salt + "'"}
	}

	return overrides, nil
}

func randomSalt() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// applyOverrides walks env line by line and, for every line assigning one
// of overrides' keys (commented out or not, as .env.example ships them),
// replaces it with "KEY=value". Keys with no matching line are appended at
// the end.
func applyOverrides(env string, overrides []envOverride) string {
	lines := strings.Split(env, "\n")
	remaining := make(map[string]string, len(overrides))
	for _, o := range overrides {
		remaining[o.key] = o.value
	}

	for i, line := range lines {
		trimmed := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "#"))

		for key, value := range remaining {
			if trimmed == key || strings.HasPrefix(trimmed, key+"=") {
				lines[i] = key + "=" + value
				delete(remaining, key)

				break
			}
		}
	}

	var b strings.Builder
	b.WriteString(strings.Join(lines, "\n"))

	for _, o := range overrides {
		if value, ok := remaining[o.key]; ok {
			fmt.Fprintf(&b, "\n%s=%s\n", o.key, value)
		}
	}

	return b.String()
}
