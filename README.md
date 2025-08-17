### Go Recipes

A curated collection of idiomatic Go demos and patterns. The goal is to showcase common language features, standard library techniques, and production-ready snippets in small, runnable programs.

### Requirements

- Go 1.22 installed locally. On macOS (Homebrew), this project expects the Go binary at:
  `/opt/homebrew/opt/go@1.22/bin/go`

### Quick Start

- Print toolchain info:

```bash
make env
```

- Run any single-file demo:

```bash
make run FILE=concurrency/ping-pong.go
```

- Test and build everything:

```bash
make test
make build
```

You can override the Go binary by exporting `GO` or passing it to make:

```bash
GO=$(which go) make test
```

### Project Layout (selected)

- `algorithm/`: classic algorithm problems with tests
- `concurrency/`: goroutines, channels, fan-in/out, cancellation patterns
- `network/`: HTTP client/server basics
- `cron/`: scheduling tasks with robfig/cron
- `circuit_breaker/`: Hystrix-style resilience patterns
- `encoding/`: runes, whitespace handling
- `examples/`: focused, runnable demos (Go 1.18+ features)

### New Demos (Go 1.18+ features and best practices)

- Context cancellation and timeouts: `examples/context/cancellation/main.go`
- Error wrapping with `errors.Is/As`: `examples/errors/wrap_is_as/main.go`
- Generics helpers (`Map`, `Filter`): `examples/generics/slices_filter_map/main.go`
- Worker pool pattern: `examples/sync/worker_pool/main.go`
- `sync.Once` and `sync.Pool`: `examples/sync/once_pool/main.go`
- Simple rate limiter via ticker: `examples/rate_limiter/ticker/main.go`
- Static assets with `go:embed`: `examples/embed/hello/main.go`
- JSON `omitempty` and (un)marshal: `examples/json/marshal_unmarshal/main.go`
- Select with timeout: `examples/time/timeout_select/main.go`

Each file is a self-contained `main` package so you can run it directly with `make run-example NAME=...`.

### Contributing

Contributions are welcome! Please:

- Keep demos small, focused, and runnable
- Prefer standard library solutions where possible
- Include comments explaining why the pattern is useful
- Add tests if the demo benefits from verification

### License

This project is licensed under the MIT License. See `LICENSE` for details.

### Acknowledgements

This repository vendors a few dependencies (see `vendor/`) with their original licenses preserved in-place. We are grateful to the authors and communities behind those projects.

