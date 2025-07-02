# Baristeuer

Baristeuer is a proof-of-concept desktop application for managing **Vereinssteuererklärung** (club tax returns). The project combines a React front end using Material UI with a Go back end.

## Features

- Desktop client built with [Wails](https://wails.io/)
- React + Material UI interface located under `internal/ui`
- Go back end located under `cmd`

## Installation

The project requires [Go](https://go.dev/), [Node.js](https://nodejs.org/) and the [Wails](https://wails.io/) CLI:

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## Development

Change into the `cmd` directory and start the application in development mode:

```bash
wails dev
```

This command watches both the backend and the frontend and reloads the application automatically.
