# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- Running `coela new` from source (`go run .`, a plain `go build`) now generates projects under `tmp/` (gitignored) instead of the repo root, so iterating on the CLI doesn't scatter generated projects across the coela source tree. Release binaries are unaffected: they always generate in the current working directory (`cmd.devOutputDir` is forced back to `""` via `.goreleaser.yaml`'s ldflags).

## [v0.1.6] 2026-07-09

### Added

- WordPress (Bedrock) projects now get `app/.env` auto-configured after install, mirroring what Symfony projects already had: `DATABASE_URL` set to a DSN matching the scaffolded database service, `WP_HOME`/`WP_SITEURL`/`WP_ENV` set from the wizard's answers (Traefik on/off), and a fresh, unique set of WordPress secret keys/salts generated per project instead of Bedrock's shared `generateme` placeholders.
- A root `.env.tmpl` and `.gitignore.tmpl` for the WordPress stack, so `DB_DATABASE`/`DB_USER`/`DB_PASSWORD`/`DB_ROOT_PASSWORD` are available to `docker-compose.yml` and the generated root `.env` isn't tracked — WordPress projects only had these for the Symfony stack until now.

### Changed

- `coela` and `coela new` now show a real description and usage instead of the Cobra boilerplate (`Short`/`Long` help text).

## [v0.1.5] 2026-07-08

### Fixed

- `coela new` failed with `Error: lstat assets/symfony: no such file or directory` when run from a binary downloaded off GitHub Releases, because stack templates were read from a path (`assets/<stack>/`) relative to the current working directory instead of being bundled into the binary. Templates now live under `internal/stacks/<stack>/assets/` and are embedded at build time via `go:embed`, so `coela new` works standalone regardless of the working directory.

### Changed

- Release CI now pins `goreleaser-action` to GoReleaser `v2.17.0` instead of `latest`, so a release build can't silently pick up a new GoReleaser version and behave differently.

## [v0.1.4] 2026-07-08

### Changed

- README rewritten as a landing page: Quick Start, Supported Stacks table, per-OS installation steps (Linux, macOS — including the Gatekeeper quarantine fix — and Windows), and a reorganized Documentation, Roadmap, and Philosophy layout.

## [v0.1.1] - 2026-07-08

### Fixed

- release CI has uncorrect env secret name

## [v0.1.0] - 2026-07-08

### Added

- Auto-configuration of generated Symfony projects: `app/.env.local` is now populated with `DATABASE_URL` (MariaDB, MySQL, or PostgreSQL), `MAILER_DSN`, `DEFAULT_URI`, and `NOREPLY_EMAIL`, matching the scaffolded Docker services.
- `docs/` reference documentation: architecture, development workflow, versioning, release process, templates, roadmap, and generated project structure.
- Continuous integration via GitHub Actions (`.github/workflows/ci.yml`): runs on every push and pull request, checks formatting (`go fmt`), `go vet`, `go test`, and `go build`.

### Changed

- README trimmed: Project Structure, Development, and Architecture sections now link to `docs/` instead of duplicating it. The Roadmap section was removed in favor of `docs/roadmap.md`.
- `.goreleaser.yaml` cleaned up: builds now target only linux/darwin/windows on amd64 and arm64 (386 dropped), the binary is always named `coela`, version/commit/date are injected via ldflags, and the release footer was removed.

### Removed

- Root `ARCHITECTURE.md`, superseded by `docs/architecture.md`.

### Fixed

- Generated `app/.env` could end up with a duplicate `DEFAULT_URI` block, depending on the installed `symfony/skeleton` version. Symfony-specific overrides are now written only to `app/.env.local`, never appended to `app/.env`.

### Changed

- Project renamed from StackForge to Coela (CLI command, module path, and documentation).

### Added

- Initial release
- Interactive CLI
- Symfony scaffolding
- Docker environment generation
- Traefik configuration
- Mailpit support
- Redis support
- Automatic Symfony installation