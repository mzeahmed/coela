# Coela Architecture

## Overview

Coela is a Go CLI that scaffolds complete PHP development environments.

Its goal is simple:

> Generate a fully working development environment with a single command.

The project currently supports:

- Symfony
- WordPress (Bedrock)

Laravel will be added in a future release.

---

# Design Principles

Coela follows a pragmatic architecture.

We intentionally avoid unnecessary abstractions.

The project follows:

- KISS
- YAGNI
- Convention over Configuration
- Single Responsibility Principle

We only introduce abstractions when a real need appears.

---

# Architecture

```
cmd/

internal/

    scaffold/

    project/

    stacks/

        symfony/

        wordpress/

    ui/

assets/

docs/
```

---

# Package Responsibilities

## cmd

Contains all CLI commands.

Example:

- new

The commands orchestrate the application.

They contain no business logic.

---

## project

Contains the Project model.

Project represents the answers collected during the interactive wizard.

Example:

```go
type Project struct {
    Name string

    Stack string

    PHPVersion string

    Database string

    Redis bool

    Mailpit bool

    Traefik bool
}
```

No business logic belongs here.

---

## ui

Contains reusable terminal components.

Current components:

- Input
- Select
- Confirm

Only terminal interaction.

---

## stacks

Each stack is self-contained.

Example:

```
stacks/

    symfony/

        wizard.go

        generate.go

        install.go

        templates/

    wordpress/

        wizard.go

        generate.go

        install.go

        templates/
```

Each stack is responsible for:

- collecting its own configuration
- exposing its templates
- installing the framework

A stack knows nothing about another stack.

---

## scaffold

Responsible for scaffolding the project.

Responsibilities:

- create directories
- execute Go templates
- generate configuration files

The scaffold package does not know Symfony.

The scaffold package only receives:

- a Project
- a template directory

---

# Assets

Project templates are stored under:

```
assets/

    symfony/

    wordpress/
```

Each stack owns its templates.

Example:

```
assets/

    symfony/

        docker-compose.yml.tmpl

        docker/

            php/

                Dockerfile.tmpl

            nginx/

                default.conf.tmpl

        Makefile.tmpl

        README.md.tmpl
```

---

# Project Generation Flow

```
coela new

        │

        ▼

Select Stack

        │

        ▼

Run Wizard

        │

        ▼

Project

        │

        ▼

Scaffold

        │

        ▼

Generate project structure

        │

        ▼

Install framework

        │

        ▼

Ready to use
```

---

# Generated Project

Coela generates the following structure.

```
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

The application itself is installed inside:

```
app/
```

Examples:

Symfony

```
app/

    bin/

    config/

    src/

    templates/

    vendor/
```

WordPress Bedrock

```
app/

    config/

    web/

    vendor/
```

---

# Current Scope

Version 1 focuses on:

- Symfony
- WordPress Bedrock
- Docker
- Traefik
- HTTPS
- Docker Compose
- Makefile
- Automatic framework installation

Nothing more.

---

# Future Evolution

Future versions may add:

- Laravel
- Meilisearch
- Mercure
- RabbitMQ
- PostgreSQL
- MongoDB

Those features will only be introduced when they solve a real problem.

The architecture should remain simple.

If a future feature requires a new abstraction, it will be introduced at that time.

Never before.

---

# Philosophy

Coela is not a generic template engine.

Coela is a developer tool.

Its mission is to eliminate repetitive project setup while producing clean, understandable and maintainable development environments.

Whenever there is a choice between flexibility and simplicity,

**simplicity wins.**