<!-- markdownlint-disable MD024 -->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Added initial `CHANGELOG.md` following Keep a Changelog format.

### Changed

- Updated CLI usage to `doit [task]` to reflect that running without arguments lists available tasks.
- Clarified CLI argument validation message for too many positional arguments.

### Removed

- Removed unused French docs translation file while localization is not implemented.

## [0.2.1] - 2026-02-08

### Changed

- Synced init template and installation notes across docs and setup flow.
- Improved project description and docs-site links.
- Standardized dependency installation instructions.

### Fixed

- Removed redundant `go mod download` step from task automation.
- Removed redundant `dev-quiet` task check in default config tests.
- Added clear prerequisites and Go version requirements in `README.md`.

## [0.2.0] - 2026-02-07

### Added

- Added GitHub Pages deployment workflow for docs.

### Changed

- Performed release readiness updates for `v0.2.0`.

## [0.1.0] - 2026-01-29

### Added

- Initial CLI application structure and task runner foundation.
- Environment variable support in task commands.
- Project license (`MIT`) and initial documentation.

### Changed

- Refactored project layout to the current command/package structure.
- Improved task categories and task metadata/schema descriptions.

### Fixed

- Corrected tasks file naming (`tasks.yml`).
- Updated author field to consistently use array format across config and docs.
- Improved init command error handling and user feedback.

[Unreleased]: https://github.com/tidjee-dev/doit/compare/v0.2.1...HEAD
[0.2.1]: https://github.com/tidjee-dev/doit/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/tidjee-dev/doit/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/tidjee-dev/doit/releases/tag/v0.1.0

<!-- markdownlint-enable MD024 -->
