# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

u2json is a Go CLI tool that parses URLs and outputs their components as JSON. It uses Go's standard `flag` and `net/url` packages with no external dependencies.

## Commands

```bash
# Build
go build .

# Run all tests
go test -v .

# Run a single test
go test -v -run TestBasic .

# Lint (CI uses golangci-lint v1.29 — needs updating for Go 1.26)
# No local config file; CI runs with defaults
```

## Architecture

Single-package (`main`) app with two core files:

- **main.go** - CLI entry point with two flags: `--query-array` (parse duplicate query params as arrays) and `--use-ParseRequestURI` (strict URI validation). The `run()` function takes args/stdout/stderr for testability.
- **convert.go** - `convert(input string, opt *convertOpt)` parses a URL and returns JSON bytes. The `convertOpt` struct carries flag state. Output is a `urlResult` struct with fields in alphabetical order (matching JSON key order) and `omitempty` tags (port is numeric, not string).

Tests live in `convert_test.go` (unit tests for `convert()`) and `main_test.go` (integration tests for `run()`).
