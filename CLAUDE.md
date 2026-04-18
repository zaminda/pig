# pig

Simple Bubbletea TUI that displays a list. Spike for Homebrew distribution.

## Commands

- `go build -o pig .` — build
- `go test ./...` — run tests
- `go vet ./...` — lint

## Structure

- `main.go` — entrypoint
- `tui/` — Bubbletea model, view, update
- `.goreleaser.yaml` — release config
- `.github/workflows/release.yml` — tag-triggered release

## Distribution

Released via GoReleaser. Installed via `brew tap zaminda/tap && brew install pig`.
Homebrew formula lives in `zaminda/homebrew-tap`.
