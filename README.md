# go-mono

An opinionated Go monorepo scaffold supporting multiple services, CLI tools, and shared libraries.

## Dependencies

| Tool | Purpose | Install (macOS) |
| ---- | ------- | --------------- |
| [Go](https://go.dev/doc/install) | Language toolchain | `brew install go` |
| [Docker](https://docs.docker.com/engine/install/) | Container builds | `brew install --cask docker` |
| [Task](https://taskfile.dev/installation/) | Task runner | `brew install go-task/tap/go-task` |
| [golangci-lint](https://golangci-lint.run/docs/welcome/install/local/) | Linter | `brew install golangci-lint` |

## Layout

```sh
cmd/
  <name>/             # service binary — thin main.go (config + logger + service.Run)
  <name>-cli/         # CLI companion tool for the same service
internal/
  service/<name>/     # service business logic (imported by cmd/<name> and cmd/<name>-cli)
  pkg/<name>/         # shared business libraries (cross-service)
go.work               # workspace tying all modules together for local development
```

## Module structure

The repo uses a **Go workspace** (`go.work`) with one module per binary:

- The **root `go.mod`** (`github.com/dyptan-io/go-mono`) owns everything under `internal/`. It is also the place for `go tool` directives.
- Each **`cmd/<name>/go.mod`** is its own module. It declares `require github.com/dyptan-io/go-mono v0.0.0` and a matching `replace` directive pointing back to the repo root, so `go mod tidy` resolves cross-module imports without hitting the network.
- `go.work` lists all modules so the toolchain resolves everything locally during development.
- `go.work.sum` is gitignored — it is auto-generated and can be recreated with `go work sync`.

There are **no `go.mod` files inside `internal/`** — all packages there belong to the root module.

## Tasks

```sh
task build              # build all apps
task build:<name>       # build one app
task run:<name>         # run one app
task test               # test all packages (go test ./... from root)
task test:<name>        # test internal/service/<name> + cmd/<name>
task lint               # lint entire codebase
task tidy               # go mod tidy every module + go work sync
task generate           # go generate entire codebase
task docker:<name>      # build Docker image for one app
```

Run `task` (no args) to list all available tasks.

## Build policy

`task build:<name>` uses source-file checksums to skip unchanged apps:

- Change in `cmd/<X>/` → only `<X>` rebuilds.
- Change in `internal/service/<X>/` → only `<X>` and `<X>-cli` rebuild.
- Change in `internal/pkg/**` or `internal/platform/**` → all apps rebuild.

## Adding a new service

1. Create `internal/service/<name>/` with your business logic (no `go.mod` needed).
2. Create `cmd/<name>/` with `main.go` and a `go.mod`:

   ```go
   module github.com/dyptan-io/go-mono/cmd/<name>

   require github.com/dyptan-io/go-mono v0.0.0
   replace github.com/dyptan-io/go-mono => ../..
   ```

3. Add `./cmd/<name>` to `go.work` and run `task tidy`.
4. If a CLI tool is needed, repeat steps 2–3 for `cmd/<name>-cli`.
5. Add per-app path filters to `.github/workflows/docker-publish.yaml` under the `Generate path filters` step — the new app is picked up automatically by all tasks.

## Docker

```sh
task docker:<name>      # local image build
```

The Dockerfile is parameterised via `--build-arg APP=<name>`. CI builds and pushes per-app images to `ghcr.io/<owner>/go-mono/<name>` on every push. Changed-app detection uses `dorny/paths-filter` — only apps whose source paths were touched are rebuilt.
