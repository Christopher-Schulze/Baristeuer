# Baristeuer

Baristeuer is a proof-of-concept desktop application for managing **Vereinssteuererklärung** (club tax returns). The project combines a React front end using Material UI with a Go back end.

## Features

- Desktop client built with [Wails](https://wails.io/)
- React + Material UI interface located under `internal/ui`
- Go back end located under `cmd`

## Installation Prerequisites

- **Go** 1.20 or later
- **Node.js** 18 or later
- **Wails** (install via `go install github.com/wailsapp/wails/v2/cmd/wails@latest`)

Ensure these tools are available in your `PATH` before building the project.

## Basic Usage

1. Install dependencies:
   ```bash
   cd internal/ui && npm install
   ```
2. Run the application in development mode:
   ```bash
   wails dev
   ```

This launches the Wails runtime with hot reloading for the React front end and the Go back end.

## Future Plans

- Implement the tax declaration forms and local data storage
- Package the application via Wails for Windows, macOS and Linux
- Add PDF export and optional cloud sync

Contributions and feedback are welcome!
