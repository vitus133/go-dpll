# Copilot instructions for go-dpll

Purpose: help AI coding agents be productive in this repository.

- Quick start
  - Build: `make build` (produces `dpll-cli` via `go build -o dpll-cli main.go`).
  - Test: `make test` (runs tests in `pkg/dpll-ynl`).
  - Tidy/vendor: `make tidy` (runs `go mod tidy && go mod vendor`).
  - Container: `make next` / `make image` — the Containerfile requires `.rh_user` and `.rh_pass` files and uses `podman --secret` to build and push.

- High-level architecture
  - CLI: `cmd/` (Cobra); `main.go` just calls `cmd.Execute()`.
  - Core netlink logic: `pkg/dpll-ynl/` — contains UAPI constants (`dpll-uapi.go`), encoder/decoder logic (`dpll.go`) and tests (`dpll_test.go`).
  - Binary encodings: sample binary blobs used for tests live in `pkg/dpll-ynl/testdata/`.

- Important patterns and conventions (do not change without care)
  - Netlink/genetlink usage: Dial with `genetlink.Dial`, then `GetFamily("dpll")`. Message headers set `Command` from constants in `dpll-uapi.go` and `Version` from the family. See `pkg/dpll-ynl/dpll.go` for examples.
  - Attribute encoding/decoding: use `netlink.NewAttributeEncoder()` to build requests and `netlink.NewAttributeDecoder()` to parse replies. Nested attributes are handled via `ad.Nested(...)` callbacks.
  - Many Do* methods expect exactly one reply (e.g., `DoDeviceGet`); callers and tests rely on that. Preserve error semantics when changing these flows.
  - Human-readable helpers live in `dpll-uapi.go` (e.g., `GetLockStatus`, `GetPinType`, `GetDpllStatusHR`). Use them when producing JSON/human output.
  - Tests validate precise binary encodings (see `pkg/dpll-ynl/dpll_test.go`). If you change encoding logic, update or regenerate testdata files accordingly.

- How to add a new netlink attribute/command
  1. Add the new constant(s) to `pkg/dpll-ynl/dpll-uapi.go` (attributes or commands block).
  2. Update parsing/encoding in `pkg/dpll-ynl/dpll.go` (switch statements in `ParseDeviceReplies`, `ParsePinReplies`, or encode paths using `netlink.NewAttributeEncoder`).
  3. Add/adjust a Cobra subcommand under `cmd/` to expose CLI functionality.
  4. Add tests in `pkg/dpll-ynl/` and, if binary encoding is involved, update `testdata/` blobs.
  5. Run `make test` and `make tidy` before submitting changes.

- Developer workflows and tips
  - Use `make run` for quick local runs; it uses `sudo /usr/local/go/bin/go run main.go` in this Makefile.
  - Container builds push to `quay.io/vgrinber/tools:dpll` by default — change `Makefile` or `IMG` in README if you push to a different repo.
  - The repo vendors dependencies; prefer `make tidy` to refresh `vendor/`.

- Integration and external dependencies
  - Uses `github.com/mdlayher/genetlink` and `github.com/mdlayher/netlink` for netlink interactions.
  - CLI framework: `github.com/spf13/cobra` (see `cmd/`).

If anything here is unclear or you want instructions expanded (examples for editing a parsing switch or a short walkthrough for adding a CLI command), tell me which section to expand.
