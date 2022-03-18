# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Support additional arguments on the go test command. These are configured by an `extraArgs` field at the test group version level.

### Changed

- `coverProfileFilePath` and `coverPackages` options have been removed, they must be replaced with entries in the `extraArgs` field. Example:
   ```
   extraArgs:
      - "-coverpkg=github.com/your/module/..."
      - "-coverprofile=./coverage.out"
   ```

## [0.1.3] - 2022-03-17

### Added

- Generate test coverage reports (one per test group version) given by two additional configuration options:
  - coverPackages: maps to go test `-coverpkg`.
  - coverProfileFilePath: maps to go test `-coverprofile`.

## [0.1.2] - 2022-03-08

### Added

- Support `exposed_ports` in container dependencies configuration


## [0.1.1] - 2022-02-21

### Added
- Additional config options for Docker container dependencies
  - Mounts
  - Entrypoint
  - CMD
  - ExtraHosts
  - Labels

## [0.1.0] - 2022-02-11

First Beta version

