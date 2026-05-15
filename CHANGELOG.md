# Changelog

All notable changes to **pxconv** are documented in this file.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).
This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] — 0000-00-00

### Added

- `.golangci.yml` — enabled full required linter set plus a solid baseline configuration:
    - `errcheck` — detects ignored errors
    - `exhaustive` — ensures completeness of enum `switch` statements
    - `revive` — naming rules, exported types, receiver naming conventions
    - `godot` — enforces punctuation in documentation comments (scope limited to declarations)
    - `govet` (with `shadow`), `staticcheck`, `unused`, `gosimple`, `ineffassign`, `misspell`, `gofmt`, `whitespace`

- Test files are excluded from `errcheck` and `exhaustive`, following common Go testing practices.

- `ci.yml` — CI pipeline split into independent jobs:
    - **test**
        - Go version matrix: `stable`, `oldstable`, `1.21`
        - runs `go vet`, unit tests, and `-race` checks
    - **lint**
        - `golangci-lint-action@v6`
    - **staticcheck**
        - separate job as required by issue specifications
    - **docs**
        - runs `go test -run=^Example ./...` to validate all documentation examples
    - **benchmark (smoke)**
        - ensures benchmarks compile and do not break during pull requests

- `pxconv.go` improvements for `godot` and `revive` compliance:
    - fixed documentation comments:
        - `// PxPerDp - number of pixels...`
        - → `// PxPerDp is the number of pixels...`
        - same updates applied to `PxPerSp` and `Dpi`
    - added `Deprecated:` annotation to `ScaleByDpi`
        - aligns with linting rules and prepares API for future cleanup
