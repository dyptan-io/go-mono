# Go-Go

An opinionated Go project scaffold for [mono-repository](https://en.wikipedia.org/wiki/Monorepo). Structure, tooling, and conventions for Go applications, services, and CLIs.

## Dependencies

| Tool | Purpose | Install (macOS) |
| ---- | ------- | --------------- |
| [Go](https://go.dev/doc/install) | Language toolchain | `brew install go` |
| [Docker](https://docs.docker.com/desktop/install/mac-install/) | Container builds | `brew install --cask docker` |
| [Task](https://taskfile.dev/installation/) | Task runner | `brew install go-task/tap/go-task` |
| [golangci-lint](https://golangci-lint.run/welcome/install/) | Linter | `brew install golangci-lint` |

## Usage

```sh
task build    # compile binary to bin/app
task run      # build and run the application
task test     # run tests with race detector
task lint     # run golangci-lint
task generate # run go generate
```

Run `task --list` to see all available tasks.
