# Coela

> Pronounced "séla".

> From zero to a ready-to-code PHP project.

[![CI](https://github.com/mzeahmed/coela/actions/workflows/ci.yml/badge.svg)](https://github.com/mzeahmed/coela/actions/workflows/ci.yml)
[![Release](https://github.com/mzeahmed/coela/actions/workflows/release.yml/badge.svg)](https://github.com/mzeahmed/coela/actions/workflows/release.yml)
[![License](https://img.shields.io/github/license/mzeahmed/coela)](https://github.com/mzeahmed/coela/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/mzeahmed/coela)](https://github.com/mzeahmed/coela/blob/main/go.mod)
[![Latest Release](https://img.shields.io/github/v/release/mzeahmed/coela)](https://github.com/mzeahmed/coela/releases/latest)
[![Go](https://img.shields.io/badge/Go-1.26+-00ADD8?logo=go)](https://go.dev/)

Coela is a CLI written in Go that scaffolds complete, Docker-based PHP development environments. Instead of hand-writing Docker Compose files, Dockerfiles, Traefik configuration, and project structure every time, Coela generates all of it — and installs your framework of choice — with a single command.

---

## Quick Start

### 1. Download

Get the latest release: **[github.com/mzeahmed/coela/releases/latest](https://github.com/mzeahmed/coela/releases/latest)**

Download the archive for your OS, extract the `coela` binary, and put it on your `PATH`. See [Installation](#installation) below for the exact commands per platform.

### 2. Verify

```bash
coela --help
```

### 3. Create a project

```bash
coela new
```

Coela generates a complete PHP development environment, ready to use.

---

## Features

- **Frameworks** — Symfony, WordPress (Bedrock)
- **Infrastructure** — Docker Compose, Nginx, PHP-FPM
- **Local HTTPS** — Traefik reverse proxy with automatically generated certificates
- **Databases** — MariaDB, MySQL, or PostgreSQL
- **Mailpit** — catches outgoing email locally
- **Redis** — optional cache and session store
- **Automatic installation** — the selected framework is installed for you

---

## Supported Stacks

| Framework | Status |
|---|---|
| Symfony | ✅ Supported |
| WordPress (Bedrock) | ✅ Supported |
| Laravel | 🚧 Planned |

---

## Installation

### Download a binary

Prebuilt binaries for Linux, macOS, and Windows are available on the [GitHub Releases](https://github.com/mzeahmed/coela/releases/latest) page. Pick the archive matching your OS and architecture (`amd64` or `arm64`).

#### Linux

```bash
curl -LO https://github.com/mzeahmed/coela/releases/latest/download/coela_Linux_x86_64.tar.gz
tar -xzf coela_Linux_x86_64.tar.gz
sudo mv coela /usr/local/bin/
coela --help
```

Use `coela_Linux_arm64.tar.gz` on ARM64 (Raspberry Pi, AWS Graviton, ...).

#### macOS

```bash
curl -LO https://github.com/mzeahmed/coela/releases/latest/download/coela_Darwin_arm64.tar.gz
tar -xzf coela_Darwin_arm64.tar.gz
sudo mv coela /usr/local/bin/
coela --help
```

Use `coela_Darwin_x86_64.tar.gz` on Intel Macs.

Coela isn't notarized by Apple. If macOS refuses to run it ("cannot be opened because the developer cannot be verified"), clear the quarantine flag:

```bash
xattr -d com.apple.quarantine /usr/local/bin/coela
```

#### Windows (PowerShell)

```powershell
Invoke-WebRequest -Uri https://github.com/mzeahmed/coela/releases/latest/download/coela_Windows_x86_64.zip -OutFile coela.zip
Expand-Archive coela.zip -DestinationPath .
.\coela.exe --help
```

Use `coela_Windows_arm64.zip` on ARM64 devices. Move `coela.exe` to a folder of your choice and add that folder to your `PATH` (Settings → System → About → Advanced system settings → Environment Variables) to run `coela` from any terminal.

### Build from source

```bash
git clone https://github.com/mzeahmed/coela.git
cd coela
go install
```

---

## Usage

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

Coela will automatically:

- create the project
- generate Docker configuration
- configure Traefik
- install the selected framework

---

## Generated Project Structure

Every generated project ships with Docker Compose, Nginx, PHP-FPM, optional Traefik, and the installed framework under `app/` — ready to run.

See [Project Structure](docs/project-structure.md) for the full layout.

---

## Requirements

- Docker
- Docker Compose
- Composer
- Go (only if building from source)

---

## Documentation

| Guide | Description |
|---|---|
| [Architecture](docs/architecture.md) | Packages, responsibilities, and the `coela new` workflow |
| [Development](docs/development.md) | Local Go workflow — format, vet, test, build, install |
| [Versioning](docs/versioning.md) | How version numbers are chosen |
| [Release Process](docs/release.md) | Steps to cut a new release |
| [Templates](docs/templates.md) | How templates are structured and rendered |
| [Roadmap](docs/roadmap.md) | What's planned for upcoming versions |
| [Project Structure](docs/project-structure.md) | What a generated project looks like |

---

## Architecture

Coela follows a small, single-responsibility package layout: `cmd/`, `internal/project`, `internal/scaffold`, `internal/stacks/*`, `internal/ui`, and `internal/traefik`.

See [Architecture](docs/architecture.md) for the full breakdown.

---

## Philosophy

- **KISS** — keep the project simple
- **YAGNI** — don't implement features before they're needed
- **Convention over Configuration** — a generated project works immediately, no manual setup
- **Developer Experience First** — save time, reduce boilerplate

---

## Roadmap

**Current**
- ✅ Symfony
- ✅ WordPress (Bedrock)

**Next**
- 🚧 Laravel

**Later**
- Additional integrations (Mercure, RabbitMQ, MongoDB, Elasticsearch, and more)

See the full [Roadmap](docs/roadmap.md).

---

## License

This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file for details.