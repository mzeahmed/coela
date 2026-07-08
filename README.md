# Coela

> From zero to a ready-to-code PHP project.

Coela is a CLI written in Go that scaffolds complete Docker-based development environments for modern PHP applications.

Instead of manually creating Docker Compose files, Dockerfiles, Traefik configuration, Makefiles and project structure every time, Coela generates everything for you.

The goal is simple:

> **Create a fully working development environment with a single command.**

---

## Features

Current features:

- Symfony support
- WordPress (Bedrock) support
- Docker Compose
- Traefik
- HTTPS local development
- Nginx
- PHP
- MariaDB / MySQL / PostgreSQL
- Mailpit
- Redis (optional)
- Ready-to-use project structure
- Automatic framework installation

---

## Project Structure

Coela generates a ready-to-use project: Docker Compose, Nginx, PHP-FPM, Traefik (optional), and the installed framework itself under `app/`.

See [Project Structure](docs/project-structure.md) for the full generated layout.

---

## Philosophy

Coela follows a few simple principles.

### KISS

Keep the project simple.

### YAGNI

Do not implement features before they are needed.

### Convention over Configuration

A generated project should work immediately without requiring manual configuration.

### Developer Experience First

The CLI should save time, reduce boilerplate and provide a consistent project structure.

---

## Usage

Create a new project:

```bash
coela new
```

The interactive wizard will ask:

```text
? Project name
> bookingapp

? Stack
❯ Symfony
  WordPress

? PHP Version
❯ 8.4

? Database
❯ MariaDB

? Use Redis?
Yes

? Use Mailpit?
Yes

? Use Traefik?
Yes
```

Coela will then:

- Create the project structure
- Generate Docker configuration
- Install the selected framework
- Produce a ready-to-use development environment

---

## Requirements

- Docker
- Docker Compose
- Go (only for development)
- Composer

---

## Development

Clone the repository and run it from source:

```bash
git clone https://github.com/mzeahmed/coela.git
cd coela
go run . new
```

See [Development](docs/development.md) for the full local workflow (format, vet, test, build, install).

---

## Architecture

Coela follows a small, single-responsibility package layout — `cmd/`, `internal/project`, `internal/scaffold`, `internal/stacks/*`, `internal/ui`, and `internal/traefik`.

See [Architecture](docs/architecture.md) for the full breakdown and the `coela new` workflow.

---

## Why Coela?

Because creating the same Docker configuration over and over again is boring.

Developers should spend time building applications, not copying boilerplate.

Coela automates the repetitive work while keeping the generated project clean, understandable and customizable.

---

## Documentation

Detailed, technical documentation lives in [`docs/`](docs/):

- [Architecture](docs/architecture.md) — packages, responsibilities, and the `coela new` workflow.
- [Development](docs/development.md) — the local Go workflow (format, vet, test, build, install).
- [Versioning](docs/versioning.md) — how Coela's version numbers are chosen.
- [Release Process](docs/release.md) — the steps to cut a new release.
- [Templates](docs/templates.md) — how templates are structured and rendered, and how to add one.
- [Roadmap](docs/roadmap.md) — what's planned for upcoming versions.
- [Project Structure](docs/project-structure.md) — what a generated project looks like, and why.

---

## License

This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file for details.