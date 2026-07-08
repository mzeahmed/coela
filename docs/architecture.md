# Architecture

Coela is a Go CLI that scaffolds complete Docker-based development environments for PHP projects. This document describes how the codebase is organized and how a single `coela new` invocation flows from user input to a ready-to-use project.

## Design Principles

### KISS (Keep It Simple, Stupid)

Every package does one thing. There is no plugin system, no dependency injection container, and no configuration file. The codebase should be readable end to end in a single sitting.

### YAGNI (You Aren't Gonna Need It)

Abstractions are introduced only when a real, current need justifies them. Support for a new stack, database engine, or feature is added when it is actually built, not in anticipation of future ones.

### Convention over Configuration

A stack's templates always live at `internal/stacks/<stack-id>/assets/`, embedded into the binary at build time via `go:embed`. A generated project is always written to a directory named after the project. There is nothing to configure beyond answering the wizard's questions.

## Packages

| Package | Responsibility |
|---|---|
| `cmd/` | Cobra CLI commands. Parses input and orchestrates the flow below. Contains no business logic. |
| `internal/project` | Defines `Project`, the data model holding the answers collected by a wizard (name, stack, PHP version, database, and optional services). Holds no logic of its own. |
| `internal/scaffold` | Renders a stack's template filesystem into real files on disk. Knows nothing about any specific stack — it only receives a `*project.Project` and an `fs.FS` to walk. |
| `internal/stacks/symfony`, `internal/stacks/wordpress` | One package per supported stack. Each owns its own wizard questions, its own embedded template directory, and its own framework installation command. A stack never references another stack. |
| `internal/ui` | Terminal interaction primitives (`Input`, `Select`, `Confirm`) built on top of promptui. Knows nothing about `Project` or any stack. |
| `internal/traefik` | Automates the local HTTPS setup for Traefik-fronted projects: generating an mkcert certificate and registering the project's domains in `/etc/hosts`. Knows nothing about any specific stack. |

Each stack package exposes the same three functions, called by `cmd/new.go`:

- `Wizard() (*project.Project, error)` — asks the stack-specific questions.
- `TemplatesDir() fs.FS` — returns the stack's embedded template filesystem.
- `Install(dir string) error` — installs the framework itself (via Composer).

There is no shared interface between stacks; `cmd/new.go` selects the right set of functions with a plain `switch` on the user's choice.

## Workflow

```
coela new
    │
    ▼
Select Stack
    │
    ▼
Wizard
    │
    ▼
Project
    │
    ▼
Scaffold
    │
    ▼
Framework Installation
    │
    ▼
Local HTTPS Provisioning (only if Traefik is enabled)
    │
    ▼
Ready
```

1. **Select Stack** — the user picks a stack (Symfony or WordPress/Bedrock) from an interactive list.
2. **Wizard** — the chosen stack asks its own questions (project name, PHP version, database engine, optional services) using `internal/ui`.
3. **Project** — the wizard's answers are assembled into a `*project.Project`. This is the only data structure that crosses from "asking questions" to "generating files".
4. **Scaffold** — `internal/scaffold.Generate` walks the stack's template directory and renders every `.tmpl` file it finds against the `Project`, writing the result into a directory named after the project.
5. **Framework Installation** — the stack's `Install` function shells out to Composer to install the actual framework (e.g. `symfony/skeleton`) into the project's `app/` directory.
6. **Local HTTPS Provisioning** — if the wizard's Traefik question was answered yes, `internal/traefik` generates a local TLS certificate with mkcert and adds the project's domains to `/etc/hosts`.
7. **Ready** — the project is on disk, its framework is installed, and (if applicable) its local domains resolve and are HTTPS-ready.

See [templates.md](templates.md) for how templates are structured and rendered, and [project-structure.md](project-structure.md) for what a generated project looks like.