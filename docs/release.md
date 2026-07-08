# Release Process

This document describes how a new version of Coela is released. See [versioning.md](versioning.md) for how version numbers are chosen.

## Steps

1. **Run the checks**

   ```bash
   go fmt ./...
   go vet ./...
   go test ./...
   ```

   See [development.md](development.md) for what each command does. Coela does not have a `make check` target yet — until it does, run these three commands individually.

2. **Update the README if necessary**

   If the release changes user-facing behavior (a new stack, a new flag, a new requirement), update [README.md](../README.md) accordingly.

3. **Update the CHANGELOG**

   Add a new entry to `CHANGELOG.md` describing what changed, following the existing format.

4. **Commit**

   ```bash
   git commit -m "chore: prepare release vX.Y.Z"
   ```

5. **Tag**

   ```bash
   git tag vX.Y.Z
   ```

6. **Push the branch**

   ```bash
   git push origin main
   ```

7. **Push the tag**

   ```bash
   git push origin vX.Y.Z
   ```

## Future: GoReleaser

Coela does not use [GoReleaser](https://goreleaser.com/) yet. Releases are currently produced manually with the steps above. GoReleaser is expected to automate binary builds and GitHub releases once the project's release surface stabilizes — this document will be updated when that happens.