# Changelog

All notable changes to **pxconv** are documented in this file.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).
This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] — 0000-00-00

### Added

- Property-based tests added using `pgregory.net/rapid` (v1.1.0):
    - `TestPropDpToPxRoundtrip` — verifies `DpToPx → PxToDp` stability within rounding tolerance `0.5/pxPerDp`
    - `TestPropSpToPxRoundtrip` — verifies `SpToPx → PxToSp` roundtrip stability with equivalent tolerance
    - `TestPropInchToPxEqualsDpi` — ensures `InchToPx(1) == round(Dpi)` for DPI values in range `[72, 600]`
    - `TestPropMmPerInchEqualsDpi` — validates physical invariant `25.4 mm == 1 inch`
    - `TestPropPtPerInchEqualsDpi` — validates typographic invariant `72 pt == 1 inch`
    - `TestPropDpToSpIdentity` — checks identity conversion when `PxPerDp == PxPerSp`
    - `TestPropDpToSpRoundtrip` — verifies roundtrip correctness for arbitrary density values

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

### Changed

- `go.mod`:
    - added dependency `pgregory.net/rapid` v1.1.0 for property-based testing
    - updated minimum Go version to `1.21` to align with CI matrix requirements
