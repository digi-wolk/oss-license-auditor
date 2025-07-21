
<div align="center">

<img src="assets/logo/logo-txt.svg" width="250">
<br /><br />

[![Test status](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-tests.yml/badge.svg?branch=master)](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-tests.yml)
[![Code quality status](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-code-quality.yml/badge.svg?branch=master)](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-code-quality.yml)
[![goSec security status](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-security-checks.yml/badge.svg?branch=master)](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-security-checks.yml)

[![License](https://img.shields.io/github/license/digi-wolk/oss-license-auditor?color=blue&label=License&style=flat-square)](https://github.com/digi-wolk/oss-license-auditor/blob/master/LICENSE.md)
[![Last commit](https://img.shields.io/github/last-commit/digi-wolk/oss-license-auditor.svg?color=blue&style=flat-square)](https://github.com/digi-wolk/oss-license-auditor/commits/master)
[![Contributors](https://img.shields.io/github/contributors/digi-wolk/oss-license-auditor?color=blue&style=flat-square)](https://github.com/digi-wolk/oss-license-auditor/graphs/contributors)

</div>

# OSS License Auditor

⚡ FAST and COMPREHENSIVE OSS license audit tool ⚡

OSS license auditor helps you track and validate licenses of all your dependencies in multiple programming languages.
Thanks to the speed of Golang, this is one of the **FASTEST** tools you can use to audit licenses used by the
dependencies used in your code. That means it will not be expensive to run it in your CI/CD pipelines.

Currently, this tool supports:
- **JavaScript/TypeScript**: NPM - `package-lock.json`
- **JavaScript/TypeScript**: PNPM - `pnpm-lock.yaml`
- **JavaScript/TypeScript**: Yarn - `yarn.lock` (only from NPM)
- **Go**: Go modules - `go.mod`

Later, it will support:

- **JavaScript/TypeScript**: Yarn - `yarn.lock` (from other sources)
- **Python**: pip - `requirements.txt`
- **Java**: Maven - `pom.xml`
- **PHP**: Composer - `composer.json`

If you like this project, remember to give us a star ⭐ on GitHub. (Spoiler: It motivates a lot!)
<br />You can also support the creator of this project by [buying me a coffee](https://bmc.link/pooyan).

# Project layout
Project layout is based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

# Related projects

- Docker image on Hub: TBD
- GitHub Action on Marketplace: TBD

# How to use

## CLI

To run this project, you need to have Go installed on your machine.
Then, run the following command inside the root directory of this project

```bash
go run cmd/olaudit/olaudit.go --path /path/to/your/project
```

Alternatively, you can build the project and run the binary file:

```bash
go build -o olaudit cmd/olaudit/olaudit.go
./olaudit --path /path/to/your/project
```

## GitHub Action

This project is also available as a GitHub Action. You can use it in your CI/CD pipelines.
- [GitHub Marketplace](https://github.com/marketplace/actions/oss-license-auditor)
- [GitHub repository](https://github.com/digi-wolk/olaudit-action)

Example usage:
```yaml
name: OSS License Auditor

on:
  pull_request:
    branches: [ master ]

jobs:
  oss-license-auditor:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: OSS License Auditor
        uses: digi-wolk/olaudit@v1
        with:
          path: .
```

# Test

To run unit test for this project, run the following command:

```bash
go test ./...
```

## Code Coverage

To run code coverage tests and ensure coverage is above 70%, use the provided script:

```bash
./scripts/run-coverage.sh
```

This script reads package exclusions from `.codecov.yml` and generates a coverage report. You can configure which packages to exclude from coverage calculations by editing the `.codecov.yml` file:

```yaml
# Configuration file for code coverage settings
# Packages listed here will be excluded from coverage calculations

exclude_packages:
  - cmd/olaudit
  - internal/cli
  - internal/print-dependencies
  - test/fixtures
```

To view the detailed coverage report:

```bash
go tool cover -func=coverage.out
```

Or to view the coverage report in HTML format:

```bash
go tool cover -html=coverage.out -o coverage.html
```

# How to contribute

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull
requests to us.

# Check the latest version
You can get the latest version number by:
- Checkin the [CHANGELOG.md](CHANGELOG.md) file
- Running this command:
```bash
git describe --abbrev=0 --tags --match "v*"
```

# Release a new version
You can create a release by tagging the version as follows:
```bash
# Update the version before running the command
RELEASE_VERSION="v0.2.0"
git tag "${RELEASE_VERSION}" -m "Release version ${RELEASE_VERSION}"
git push origin "${RELEASE_VERSION}"
```
The GitHub Action release workflow triggers immediately after the tag is pushed upstream. It will build and publish 
the docker image into DockerHub.

### Docker Hub

Docker image is available on [Docker Hub](https://hub.docker.com/r/prazian/oss-license-audit).

You can pull the latest version by running the following command:
```bash
docker pull prazian/oss-license-audit:latest
```

You can pull a specific version by running the following command:
```bash
docker pull prazian/oss-license-audit:v0.0.1
```

# Documentation and resources

- [Contributing](CONTRIBUTING.md)
- [Code of Conduct](CODE_OF_CONDUCT.md)
- [License](LICENSE.md)
- [Roadmap](ROADMAP.md)
- [Support](SUPPORT.md)
- [Security](SECURITY.md)
- [Changelog](CHANGELOG.md)

# Maintainers
- Pooyan Razian: [github.com/prazian](https://github.com/prazian) - [pooyan.info](https://pooyan.info)
- _< Become the first co-maintainer and add your name here! >_

# Sponsors
Support this project by becoming a sponsor. Your logo will show up here with a link to your website
