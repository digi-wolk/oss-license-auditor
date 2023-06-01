## OSS License Auditor

[![Test status](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-tests.yml/badge.svg?branch=master)](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-tests.yml)
[![Code quality status](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-code-quality.yml/badge.svg?branch=master)](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-code-quality.yml)
[![Security status](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-security-checks.yml/badge.svg?branch=master)](https://github.com/digi-wolk/oss-license-auditor/actions/workflows/run-security-checks.yml)

[![License](https://img.shields.io/github/license/digi-wolk/oss-license-auditor?color=blue&label=License&style=flat-square)](https://github.com/digi-wolk/oss-license-auditor/blob/master/LICENSE.md)
[![Last commit](https://img.shields.io/github/last-commit/digi-wolk/oss-license-auditor.svg?color=blue&style=flat-square)](https://github.com/digi-wolk/oss-license-auditor/commits/master)
[![Contributors](https://img.shields.io/github/contributors/digi-wolk/oss-license-auditor?color=blue&style=flat-square)](https://github.com/digi-wolk/oss-license-auditor/graphs/contributors)


# About

⚡ FAST and COMPREHENSIVE OSS license audit tool ⚡

OSS license analyzer helps you track and validate licenses of all your dependencies in multiple programming languages.
Thanks to the speed of Golang, this is one of the **FASTEST** tools you can use to audit licenses used by the dependencies
used in your code. That means it will not be expensive to run it in your CI/CD pipelines.

Currently, this tool supports:
- **JavaScript/TypeScript**: NPM - `package-lock.json`
- **Go**: Go modules - `go.mod`

Later, it will support:

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
go run main.go --path /path/to/your/project
```

Alternatively, you can build the project and run the binary file:

```bash
go build -o oss-license-auditor main.go
./oss-license-auditor --path /path/to/your/project
```

## GitHub Action

```yaml
TBD
```

# Test

To run unit test for this project, run the following command:

```bash
go test ./...
```

# How to contribute

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull
requests to us.

# Release a new version

### GitHub Release

TBD

### Docker Hub

TBD

### GitHub Action

TBD

# Documentation and resources

- [Contributing](CONTRIBUTING.md)
- [Code of Conduct](CODE_OF_CONDUCT.md)
- [License](LICENSE.md)
- [Roadmap](ROADMAP.md)
- [Support](SUPPORT.md)
- [Security](SECURITY.md)
- [Changelog](CHANGELOG.md)

# Maintainers
- 
- Pooyan Razian: [github.com/prazian](https://github.com/prazian) - [pooyan.info](https://pooyan.info)
- TBD (can be you!)

# Sponsors
Support this project by becoming a sponsor. Your logo will show up here with a link to your website
