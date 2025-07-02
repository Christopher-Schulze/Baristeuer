# Baristeuer

Baristeuer is a proof-of-concept desktop application for managing **Vereinssteuererklärung** (club tax returns). The project combines a React front end with a Go back end and uses [Wails](https://wails.io/) to produce a desktop application.

## Features

- Desktop client built with [Wails](https://wails.io/)
- React + Material UI interface located under `internal/ui`
- Go back end located under `cmd`
- SQLite persistence layer

## Installation

The application requires **Go**, **Node.js** and **Wails**. The following commands show a minimal setup on Debian/Ubuntu. You can also install the tools via the official download pages for other platforms.

### Go

```bash
sudo apt-get install golang-go
```

### Node.js

```bash
sudo apt-get install nodejs npm
```

### Wails

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Database

The persistence layer relies on SQLite via the `github.com/mattn/go-sqlite3` driver. Ensure SQLite development headers are available (e.g. `libsqlite3-dev` on Debian/Ubuntu).

## Building the Front End

The front end resides in `internal/ui`. To install dependencies and create a production build run:

```bash
cd internal/ui
npm install
npm run build
```

During development you can start the application with:

```bash
wails dev
```

This command will build the Go backend and start the Vite development server.

## Project Structure

- `cmd/` \u2013 entry point of the Go application
- `internal/data/` \u2013 data models and persistence code
- `internal/pdf/` \u2013 PDF generation utilities
- `internal/taxlogic/` \u2013 tax calculation logic
- `internal/ui/` \u2013 React + Vite front end source
- `docs/` \u2013 additional documentation
- `.github/workflows/` \u2013 CI configuration

## Continuous Integration

A GitHub Actions workflow (`.github/workflows/ci.yml`) runs `go vet`, `go test`, and builds the front end on every pull request:

```
- name: Run go vet
  run: go vet ./...
- name: Run go test
  run: go test ./...
- name: Install frontend dependencies
  run: npm install
  working-directory: internal/ui
- name: Build frontend
  run: npm run build
  working-directory: internal/ui
```

This ensures that both Go and Node.js code are validated automatically.
