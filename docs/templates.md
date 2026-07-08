# Templates

Every file a stack generates comes from a template rendered by `internal/scaffold`. This document explains where templates live, how they are rendered, and how to add one.

## Location

Templates are stored under `assets/`, one directory per stack, matching the stack's ID:

```
assets/
├── symfony/
│   ├── docker-compose.yml.tmpl
│   ├── .env.tmpl
│   ├── .gitignore.tmpl
│   ├── Makefile.tmpl
│   ├── README.md.tmpl
│   ├── docker/
│   │   ├── nginx/
│   │   │   └── default.conf.tmpl
│   │   └── php/
│   │       └── Dockerfile.tmpl
│   └── traefik/
│       ├── traefik.yml.tmpl
│       └── dynamic.yml.tmpl
└── wordpress/
```

A stack only ever reads from its own subdirectory.

> `assets/wordpress/` currently has no templates. The WordPress stack's `Wizard` and `Install` are implemented, but running it through `coela new` produces no scaffolded files until templates are added here — only the framework installed by Composer.

## How Templates Are Rendered

`internal/scaffold.Generate` receives a `*project.Project` and a template directory. It:

1. Walks the directory recursively.
2. For every file ending in `.tmpl`, parses it with the standard library's `text/template`.
3. Executes it with the `*project.Project` itself as the template data.
4. Writes the result to the generated project, at the same relative path, with the `.tmpl` suffix stripped.

For example, `assets/symfony/docker/php/Dockerfile.tmpl` becomes `<project>/docker/php/Dockerfile`.

Files that do not end in `.tmpl` (for example a placeholder `.gitkeep`) are skipped.

Because the `Project` struct itself is the template data, any exported field is available directly, for example:

```
{{ .Name }}
{{ .PHPVersion }}
{{ .Database }}
```

`.Database` and `.Stack` render as their human-readable label (e.g. `MariaDB`) because both types implement `String()`. Boolean fields (`.Redis`, `.Mailpit`, `.Traefik`) can drive conditional blocks:

```
{{ if .Traefik }}
...
{{ end }}
```

`assets/symfony/docker-compose.yml.tmpl` is the most complete example: it conditionally includes the Traefik, Mailpit, and Redis services, and switches the database service's image and environment variables based on `{{ .Database }}` (MariaDB, MySQL, or PostgreSQL).

Some values are intentionally not embedded directly in the template output. `docker-compose.yml.tmpl` references database credentials as `${DB_DATABASE}`, `${DB_USER}`, and so on — real shell-style variables resolved by Docker Compose from the sibling `.env` file at container start, not by Go's template engine. `assets/symfony/.env.tmpl` is what actually renders those values.

## Adding a New Template

1. Create the file under `assets/<stack>/`, ending in `.tmpl`, at the path you want it to appear in the generated project (minus the `.tmpl` suffix).
2. Use `{{ .FieldName }}` to reference any field on `project.Project`.
3. If the file should only exist for certain wizard answers, wrap the relevant section in `{{ if ... }} ... {{ end }}`. `scaffold.Generate` has no concept of skipping a whole file conditionally — every `.tmpl` file found is always rendered and written; conditionals only affect content within a file.
4. No Go code changes are required. `scaffold.Generate` discovers the new file automatically the next time it walks the stack's template directory.