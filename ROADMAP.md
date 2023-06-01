# Project Roadmap

This document outlines the planned features, enhancements, and milestones for the future development of the Dependency
License Analyzer project. Please note that the roadmap is subject to change as priorities may shift based on community
feedback, emerging technologies, or other factors.

## 2023
- [x] **License detection from package-lock.json**: Add support for detecting licenses from `package-lock.json` files.
- [x] **License detection from go.mod**: Add support for detecting licenses from `go.mod` files.
- [x] **GitHub Templates**: Create issue and pull request templates to provide a standardized format for submitting bug
  reports, feature requests, and code contributions.
- [x] **Quality Assurance**: Implement automated testing to ensure the quality of the project.
- [x] **Continuous Integration (CI)**: Set up a CI pipelines to automatically check builds, tests, and code quality.
- [x] **Code Coverage**: Add automated code coverage reports to ensure the stability and reliability of the project.
- [ ] **Custom config files for definition of risky licenses**: Instead of relying on the default configurations, allow
  users to define their own configuration files to specify the licenses that they consider risky.
- [ ] **Docker Image**: Create a Docker image for the project to allow users to run the Dependency License Analyzer in a
  containerized environment.
- [ ] **Publish Actions in GitHub Marketplace**: Publish the OSS License Auditor as
  a [GitHub Action in the GitHub Marketplace](https://docs.github.com/en/actions/creating-actions/publishing-actions-in-github-marketplace)
  to allow users to easily integrate it into their workflows.
  automated dependency analysis as part of the build and testing process.
- [ ] **Documentation and Examples**: Provide detailed documentation and usage examples to help users get started with the
  Dependency License Analyzer.
- [ ] **Refactor the code to improve maintainability**: Before adding new features, the code should be refactored 
  to improve the maintainability of the project.
- [ ] **License detection from requirements.txt**: Add support for detecting licenses from `requirements.txt` files.

## 2024
- [ ] **GitHub Sponsors**: Enable GitHub Sponsors to allow users to financially support the project and its maintainers.
- [ ] **License detection from pom.xml**: Add support for detecting licenses from `pom.xml` files.
- [ ] **License detection from build.gradle**: Add support for detecting licenses from `build.gradle` files.
- [ ] **License detection from composer.json**: Add support for detecting licenses from `composer.json` files.
- [ ] **API and Web Interface**: Develop a web-based interface and API endpoints to provide a user-friendly way to interact
  with the Dependency License Analyzer.
- [ ] **External dashboards**: Create a dashboard to visualize the results of the Dependency License Analyzer.

## Backlog
- [ ] **Dependency Graph Visualization**: Implement a visualization feature to generate a dependency graph that visualizes
  the relationships between different dependencies.

This roadmap provides an overview of our vision for the future development of the Dependency License Analyzer. It aims
to address the needs of our users and create a robust and versatile tool for managing dependencies and ensuring license
compliance. As we progress, we welcome community feedback, contributions, and suggestions to shape the project's
direction and make it even more valuable to the open-source community.