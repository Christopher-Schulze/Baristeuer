<div align="center">
  <img src="logo.png" alt="Bari$teuer Logo" width="300">
</div>

# Bari$teuer

Bari$teuer is a cross-platform desktop application that assists German non-profit associations with their tax reporting. A Go backend and a Svelte frontend styled with DaisyUI are bundled together using Wails.

## Highlights

- Calculates corporate and sales tax for non-profit projects
- Stores data in SQLite and manages member lists
- Generates PDF reports and CSV exports
- Runs on macOS and Windows

## Building

Run the following commands to compile the application:

```bash
go work sync
npm ci --prefix internal/ui
wails build
```

Binaries for macOS and Windows are built automatically via GitHub Actions.

## Tests

UI tests require Node.js dependencies located in `internal/ui`. Run `make deps`
once to install them before executing `make ui-test` or `make test`.

For installation and detailed usage instructions see [docs/DOCUMENTATION.md](docs/DOCUMENTATION.md).
Place the official PDF forms in `internal/pdf/templates/` as described in [docs/Todo-fuer-User.md](docs/Todo-fuer-User.md).

---

_This project is for internal use and is not open for contributions._
