# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Generate test coverage reports (one per test group) given by two additional configuration options:
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

