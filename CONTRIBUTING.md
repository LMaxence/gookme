# Contributing to Gookme

Thank you for your interest in contributing to this project! This guide will help you get started by explaining the project's architecture, how to set it up, and how to submit your contributions. Whether you want to add features, fix bugs, or improve documentation, all contributions are welcome.

## Getting Started

### Prerequisites

To contribute to this project, you will need:

- [Go](https://golang.org/dl/) (v1.22 or higher)
- Git
- Any text editor (e.g., VSCode, GoLand)


### Cloning the Repository

To get started, clone the repository and set up your local environment:

```bash
git clone https://github.com/LMaxence/gookme.git
cd gookme
go mod download
```

This will download all the dependencies and get the project ready to run.

## Project Architecture

Here's a high-level overview of the components that make up Gookme:

- **cmd/cli**: This folder contains the entry point for the `gookme` command.
- **cmd/schemas**: This folder contains the command utility used to generate the Gookme hook files JSON schemas.
- **docs**: This folder contains the documentation for the project. Powered by [MkDocs](https://www.mkdocs.org/) and [Material for MkDocs](https://squidfunk.github.io/mkdocs-material/).
- **hooks**: This folder contains definition of the hooks that can be used in Gookme. You should not use gookme init to install Gookme within Gookme. This will work, but you will use the released version. If you want to run the development version, you should use the makefile command `make hooks` instead.
- **packages**: This folder contains the core logic of Gookme. It is divided into several sub-packages:
  - **cli**: The scripts being executed by the CLI for each `gookme` command. Each command has its own file
  - **configuration**: This is where the different data structures related to hooks, steps, and global configuration of the Gookme CLI are defined. This is also where the logic for walking across the repositorie's files and look for hooks is defined.
  - **executor**: This is where the logic for wrapping the execution of the hooks is defined.
  - **filters**: This is where the logic for filtering the hooks is defined.
  - **git-client**: This is where the logic for interacting with the git repository is defined.
  - **hooks-scripts**: This is where the logic for managing hook files in the .git repository is defined.
  - **logging**: This is where Gookme's logger is defined.
  - **meta**: This very small package contains the version of Gookme. Changing this version will change the version of Gookme and trigger a release.
  - **test-helpers**: This is where the helpers for the tests are defined.
- **scripts**: This folder contains the scripts used to manage the project. It is divided into several sub-packages:
  - **check-formtat.sh**: This script is used to check the format of the files in the project.
  - **commit-msg**: This script is used to check the format of the commit messages.
  - **generate-dependabot-config.sh**: This script is used to generate the dependabot configuration file.
  - **install.ps1**: This script is used to install Gookme on Windows.
  - **install.sh**: This script is used to install Gookme on Unix systems.
  - **pre-commit.sh**: This script is used to run the pre-commit checks. This is what is copied to the .git hooks folder when running `make hooks`.