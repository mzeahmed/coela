# Contributing to Coela

First of all, thank you for contributing to Coela ❤️

Coela is an opinionated CLI that scaffolds complete PHP development environments.

Our goal is not to build the most flexible tool.

Our goal is to build the simplest tool that solves a real problem.

This document explains the development principles used throughout the project.

---

# Philosophy

Coela follows one simple philosophy:

> Simplicity beats cleverness.

When multiple solutions exist, always prefer the one that is:

- easier to read
- easier to maintain
- easier to explain
- easier to remove

Good software grows through small, understandable steps.

---

# Core Principles

Coela follows these principles:

- KISS (Keep It Simple, Stupid)
- YAGNI (You Aren't Gonna Need It)
- Convention over Configuration
- Single Responsibility Principle

We intentionally avoid over-engineering.

If a problem does not exist today, we do not solve it today.

---

# Project Structure

```
cmd/

internal/

    project/

    scaffold/

    stacks/

        symfony/assets/

        wordpress/assets/

    ui/

docs/
```

Every package has one responsibility.

Nothing more.

---

# Package Responsibilities

## cmd

Contains CLI commands only.

Responsibilities:

- parse CLI arguments
- orchestrate the application

Must never contain:

- business logic
- filesystem operations
- Docker logic
- template generation

---

## project

Contains domain models.

Example:

- Project

A model contains data.

A model does not contain business logic.

---

## scaffold

Responsible for creating projects.

Responsibilities:

- create folders
- execute templates
- generate files

It should not know anything about Symfony or WordPress.

---

## stacks

Each stack is completely independent.

Examples:

```
symfony/

wordpress/
```

A stack is responsible for:

- its wizard
- its templates
- installing its framework

A stack must never know another stack.

---

## ui

Contains reusable terminal components.

Examples:

- Input
- Select
- Confirm

Only terminal interaction belongs here.

---

# Code Style

## Keep functions small

Prefer functions between:

20–30 lines.

Avoid functions larger than:

50 lines.

If a function becomes difficult to understand,

split it.

---

## Keep packages focused

One package.

One responsibility.

Avoid "god packages".

---

## Prefer composition

Compose small components.

Avoid large objects that know too much.

---

## Avoid unnecessary abstractions

Do not create:

- interfaces
- factories
- strategies
- builders
- managers

unless they solve a real problem.

Every abstraction has a maintenance cost.

---

## Avoid generic names

Do not create packages or types called:

- Helper
- Utils
- Common
- Manager
- Service
- Base

Names should describe intent.

Prefer:

- Scaffold
- Project
- Stack
- Wizard
- Install

---

# Dependencies

Prefer the Go standard library.

Before adding a dependency, ask:

> Can the standard library solve this?

If yes,

use the standard library.

---

# Templates

All templates belong inside their stack's own directory:

```
internal/stacks/<stack>/assets/
```

They are embedded into the binary at build time via `go:embed`, so `coela new` works from a standalone downloaded binary.

Templates are resources.

They must never be generated from Go code.

---

# Error Handling

Never panic for expected errors.

Always return errors.

Prefer:

```go
if err != nil {
    return err
}
```

Errors should be explicit.

---

# Comments

Comments explain **why**.

Code explains **what**.

Avoid comments like:

```go
// Increment i
i++
```

Useful comments explain design decisions.

---

# Logging

Do not print debug information in production code.

Use logging only when it provides value to the user.

---

# Testing

Keep business logic independent from the filesystem whenever possible.

Prefer code that can be tested without requiring Docker.

Avoid hidden dependencies.

---

# Keep the MVP Small

Current scope:

- Symfony
- WordPress (Bedrock)
- Docker
- Traefik
- HTTPS
- Docker Compose
- Makefile

Nothing more.

If a feature is not required for the MVP,

do not implement it.

---

# The 6-Month Rule

Write code as if you were going to read it again six months from now.

If future-you cannot understand it in a few minutes,

it is probably too complex.

Always prefer:

- readability over cleverness
- explicit code over magic
- simple solutions over flexibility

The goal is not to write impressive code.

The goal is to write code that remains easy to maintain over time.

---

# When in Doubt

If you hesitate between two implementations,

choose the one that:

- introduces fewer abstractions
- introduces fewer concepts
- has fewer files
- is easier to explain
- is easier to remove
- is easier to maintain

Simple code wins.

---

# Before Opening a Pull Request

Ask yourself:

- Is this simpler than before?
- Does this introduce a new abstraction?
- Is this abstraction really needed today?
- Would a new contributor understand this code in five minutes?
- Am I solving today's problem or tomorrow's imaginary problem?

If the answer is "no",

consider simplifying the implementation.

---

# Final Thought

Software is read far more often than it is written.

Optimize for the next developer.

Most of the time,

that next developer is you.