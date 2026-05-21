# Coding Agent Guidelines

You are a Go programming assistant following the conventions of this repository.

## Getting Started

- Read README.md for project overview, layout, and setup.

## Code Style

- **Error wrapping**: Use NASA notation: `fmt.Errorf("parsing data: %w", err)`.
- **Comments**: Minimal; prefer self-documenting code.
- **Interfaces**: Define only when needed; avoid premature abstraction.
- **Type Semantics**: Prefer value types; use pointers to signal mutation.
- **Packages**: Domain-oriented structure; sub-packages for implementation (e.g., `search/searxng`).
  Shared packages → `internal/pkg/`. Transport & config → `cmd/<app-name>`.
  Avoid deep nesting and unnecessary small packages.

## Testing

Group all tests for a type or function under a single top-level `TestFoo` function with `t.Run` subtests.

### Table Tests

Prefer table-driven tests if there are multiple cases to cover.
Use a `map[string]struct{}` table when cases share the same inputs/outputs shape. Keep cases focused — don't over-inflate.

- Inputs: `give*` prefix (e.g., `giveUserID`).
- Outputs: `want*` prefix (e.g., `wantErr`).
- Import `github.com/stretchr/testify/{assert,require}`.

See example at [cmd/app/example_test.go](cmd/app/example_test.go) for reference.

## Before You Finish

Run `task lint` and fix all warnings before considering the task done.
