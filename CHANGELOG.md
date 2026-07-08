# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- Auto-configuration of generated Symfony projects: `app/.env.local` is now populated with `DATABASE_URL` (MariaDB, MySQL, or PostgreSQL), `MAILER_DSN`, `DEFAULT_URI`, and `NOREPLY_EMAIL`, matching the scaffolded Docker services.
- `docs/` reference documentation: architecture, development workflow, versioning, release process, templates, roadmap, and generated project structure.

### Changed

- README trimmed: Project Structure, Development, and Architecture sections now link to `docs/` instead of duplicating it. The Roadmap section was removed in favor of `docs/roadmap.md`.

### Removed

- Root `ARCHITECTURE.md`, superseded by `docs/architecture.md`.

### Fixed

- Generated `app/.env` could end up with a duplicate `DEFAULT_URI` block, depending on the installed `symfony/skeleton` version. Symfony-specific overrides are now written only to `app/.env.local`, never appended to `app/.env`.

## [0.1.1] - 2026-07-08

### Changed

- Project renamed from StackForge to Coela (CLI command, module path, and documentation).

## [0.1.0] - 2026-07-08

### Added

- Initial release
- Interactive CLI
- Symfony scaffolding
- Docker environment generation
- Traefik configuration
- Mailpit support
- Redis support
- Automatic Symfony installation