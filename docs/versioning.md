# Versioning

Coela follows [Semantic Versioning](https://semver.org/): `MAJOR.MINOR.PATCH`.

## PATCH

Incremented for backward-compatible bug fixes. No new behavior, no breaking change.

Example: `0.1.0` → `0.1.1`.

## MINOR

Incremented for backward-compatible additions: a new stack, a new supported database engine, a new wizard question that does not change existing behavior.

Example: `0.1.1` → `0.2.0`.

## MAJOR

Incremented for breaking changes: a change to the generated project structure, a renamed or removed wizard question, a change to the CLI's command surface.

Example: `0.9.0` → `1.0.0`.

## Why Coela Is Currently on 0.x

A `0.x` version means the project has not yet made any stability guarantee. The CLI's wizard questions, the generated project layout, and the template variables exposed to `.tmpl` files may still change without a MAJOR bump. Version `1.0.0` will mark the first release where this surface is considered stable.

## Git Tags

All tags use the `v` prefix:

```bash
git tag v0.1.0
```

This matches the convention expected by Go's own tooling (`go install module@version`) and by release tooling such as GoReleaser (see [release.md](release.md)).

> Note: the very first tags (`0.1.0`, `0.1.1`) were pushed without the `v` prefix, before this convention was written down. All tags going forward use it.