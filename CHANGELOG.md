# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added
- Added version display (version 20251126) on startup
- Added automatic parameter passing to the started application via function parameters
- Added `ignored_ext` configuration option to ignore specific file extensions
- Added automatic creation of default config file at `tmp/run.ini` if it doesn't exist
- Added default ignored frontend file extensions (`.js`, `.jsx`, `.ts`, `.tsx`, `.css`, `.scss`, `.vue`, etc.)
- Added default ignored frontend directories (`.next`, `.nuxt`, `.vuepress`, `.vite`, `node_modules`, `dist`, `build`)
- Created local `config` package to replace `github.com/pilu/config` dependency

### Changed
- Removed `-c` command line flag for config file path
- Changed default config file path from `./runner.conf` to `tmp/run.ini`
- Changed parameter passing from environment variables to function parameters for better reliability
- All command line arguments are now automatically passed to the started application
- Configuration file is automatically created with default settings if it doesn't exist
- Migrated all `github.com/pilu` dependencies to local packages
- Updated import paths to use local module `fresh` instead of `github.com/pilu/fresh`

### Configuration
- `ignored`: Comma-separated list of directories to ignore (configurable in `tmp/run.ini`)
- `ignored_ext`: Comma-separated list of file extensions to ignore (configurable in `tmp/run.ini`)

