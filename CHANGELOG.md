# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2025-12-24

### Changed
- Restructured repository to follow professional Go project layout
- Moved `cmd/` to `cmd/ascii-art/` for better package organization
- Reorganized `internal/` to `internal/ascii/` with modular file structure
- Migrated banner assets from `text/` to `internal/assets/`
- Separated code into focused modules: ascii.go, filter.go, args.go, read-file.go, printer.go
- Enhanced test coverage with dedicated test files

### Added
- CHANGELOG.md for tracking project changes
- New test files: art_test.go, banner_test.go for unit testing
- Placeholder test files: main_test.go, edge_cases_test.go for future integration and edge case tests

## [0.1.0] - Previous Release

### Added
- Initial ASCII art generator implementation
- Support for three banner styles: standard, shadow, thinkertoy
- Command-line argument parsing
- Character filtering for non-printable characters
