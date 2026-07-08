# Changelog

All notable changes to this project will be documented in this file.

## [v0.1.4]

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