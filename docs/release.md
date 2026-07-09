# Release Process

This document describes how a new version of Coela is released. See [versioning.md](versioning.md) for how version numbers are chosen.

Pushing a `vX.Y.Z` tag is what triggers a release — everything after that is automated by `.github/workflows/release.yml` and [GoReleaser](https://goreleaser.com/) (see `.goreleaser.yaml`).

## Steps

1. **Run the checks**

   ```bash
   make check
   ```

   Runs `go fmt`, `go vet`, and `go test` — see [development.md](development.md) for what each does. `make release-snapshot` additionally simulates a full GoReleaser release locally (build, archive, checksums), without publishing anything.

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

   Pushing the tag triggers the `Release` workflow, which runs `goreleaser release --clean`. It builds the Linux, macOS, and Windows binaries, packages them, generates the checksums and changelog, and publishes the GitHub Release — automatically. Nothing else to do.

## What Happens in CI

The `Release` workflow (`.github/workflows/release.yml`) only runs on tag pushes matching `v*`, never on branches. It checks out the full history (GoReleaser needs it to generate the changelog), installs Go, and runs GoReleaser via the official `goreleaser/goreleaser-action`, authenticated with the repository's built-in `GITHUB_TOKEN` — no custom secret is required.