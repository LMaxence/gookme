# Default recipe to show help
default:
    @just --list

# Install dependencies and setup the project
install:
    go mod download
    go mod tidy
    go install ./cmd/cli

alias i := install

# Run linters
lint:
    golangci-lint run

# Generate all assets including schemas, dependabot config, and git hooks
assets: schemas dependabot hooks

alias a := assets

# Prepare the project for development
prepare: install assets

alias p := prepare

# Generate all schemas
schemas:
    make assets/schemas/global.schema.json
    make assets/schemas/hooks.schema.json
    make assets/schemas/steps.schema.json

alias s := schemas

# Generate dependabot configuration
dependabot:
    make .github/dependabot.yml

alias d := dependabot

# Install all git hooks
hooks:
    make .git/hooks/pre-commit
    make .git/hooks/commit-msg

alias h := hooks

# Build for all platforms
build: build-darwin build-linux build-windows

alias b := build

# Build for Darwin (macOS)
build-darwin:
    make build/gookme-darwin-amd64
    make build/gookme-darwin-arm64

alias bd := build-darwin

# Build for Linux
build-linux:
    make build/gookme-linux-amd64
    make build/gookme-linux-arm64

alias bl := build-linux

# Build for Windows
build-windows:
    make build/gookme-windows-amd64
    make build/gookme-windows-arm64 

alias bw := build-windows

test:
    go test ./...

alias t := test

# Run all tests
check: lint test

alias c := check