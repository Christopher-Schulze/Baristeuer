# Baristeuer

Baristeuer is a proof-of-concept desktop application for managing **Vereinssteuererklärung** (club tax returns). The project combines a React front end using Material UI with a Go back end.

## Features

- Desktop client built with [Wails](https://wails.io/) (planned)
- React + Material UI interface located under `internal/ui`
- Go back end located under `cmd`

## Installation Prerequisites

- **Go** 1.20 or later
- **Node.js** 18 or later
- **Wails** (install via `go install github.com/wailsapp/wails/v2/cmd/wails@latest`)
- **SQLite** development headers (required by `github.com/mattn/go-sqlite3`)

Ensure these tools are available in your `PATH` before building the project.

## Basic Usage

1. Install dependencies:
   ```bash
   cd internal/ui && npm install
   ```
2. Start the React dev server (placeholder):
   ```bash
   npm start
   ```
3. Build and run the Go back end:
   ```bash
   cd ../../cmd && go run .
   ```

Integration with Wails will later allow packaging everything as a desktop app.

## Future Plans

- Implement the tax declaration forms and local data storage
- Package the application via Wails for Windows, macOS and Linux
- Add PDF export and optional cloud sync

Contributions and feedback are welcome!
