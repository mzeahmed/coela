# StackForge

> From zero to a ready-to-code PHP project.

StackForge is a CLI written in Go that scaffolds complete Docker-based development environments for modern PHP applications.

Instead of manually creating Docker Compose files, Dockerfiles, Traefik configuration, Makefiles and project structure every time, StackForge generates everything for you.

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
- MariaDB
- Mailpit
- Redis (optional)
- Ready-to-use project structure
- Automatic framework installation

---

## Project Structure

A generated project looks like this:

```text
bookingapp/

├── app/
├── certs/
├── docker/
│   ├── nginx/
│   └── php/
├── docs/
├── traefik/
├── docker-compose.yml
├── Makefile
├── README.md
└── .gitignore
```

The `app/` directory contains the selected framework:

- Symfony
- WordPress Bedrock

---

## Philosophy

StackForge follows a few simple principles.

### KISS

Keep the project simple.

### YAGNI

Do not implement features before they are needed.

### Convention over Configuration

A generated project should work immediately without requiring manual configuration.

### Developer Experience First

The CLI should save time, reduce boilerplate and provide a consistent project structure.

---

## Roadmap

### Version 1

- [x] Symfony
- [x] WordPress (Bedrock)
- [x] Docker
- [x] Traefik
- [x] HTTPS
- [x] Docker Compose
- [x] Makefile
- [x] Automatic installation

### Version 2

- [ ] Laravel
- [ ] Meilisearch
- [ ] Mercure
- [ ] RabbitMQ
- [ ] PostgreSQL
- [ ] MongoDB
- [ ] Elasticsearch

### Version 3

- [ ] Custom stacks
- [ ] Plugin system
- [ ] Configuration file
- [ ] Update existing projects

---

## Usage

Create a new project:

```bash
stackforge new
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

StackForge will then:

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

Clone the repository:

```bash
git clone https://github.com/mzeahmed/stackforge.git

cd stackforge
```

Run the CLI:

```bash
go run . new
```

Build:

```bash
go build -o stackforge
```

---

## Architecture

The project intentionally follows a simple architecture.

```
cmd/

internal/

    scaffold/

    project/

    stacks/

        symfony/

        wordpress/

    ui/
```

Every package has a single responsibility.

No unnecessary abstractions.

The project evolves only when a real need appears.

---

## Why StackForge?

Because creating the same Docker configuration over and over again is boring.

Developers should spend time building applications, not copying boilerplate.

StackForge automates the repetitive work while keeping the generated project clean, understandable and customizable.

---

## License

This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file for details.
