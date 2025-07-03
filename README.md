# Baristeuer

Baristeuer is a proof-of-concept desktop application for managing **Vereinssteuererkl\u00e4rung** (club tax returns). It combines a React front end styled with Tailwind CSS and a Go back end. The application is bundled with [Wails](https://wails.io/) so it can run as a native program.

## Features

- Desktop client built with [Wails](https://wails.io/)
- React + Tailwind CSS interface located under `internal/ui`
- Go back end located under `cmd`
- SQLite persistence layer

## Installation

The application requires **Go**, **Node.js** and the **Wails CLI**. The following example shows a minimal setup on Debian/Ubuntu. For other operating systems please refer to the official download pages.

### Go

Install Go 1.x from your package manager or from <https://go.dev/dl/>:

```bash
sudo apt-get update
sudo apt-get install -y golang
```

### Node.js

Install Node.js (which also installs `npm`):

```bash
sudo apt-get install -y nodejs npm
```

### Wails CLI

Install the Wails command line interface globally and verify the setup:

```bash
npm install -g @wailsapp/cli
wails doctor
```

## Project Structure

```
cmd/         Application entry point and Wails setup
docs/        Documentation and changelog
internal/
  data/      Data models and SQLite storage
  pdf/       PDF generation
  taxlogic/  Business logic for tax calculations
  ui/        React/Tailwind frontend
```

## Usage

During development start the application with:

```bash
wails dev
```

To build the UI for production manually run:

```bash
npm run build
```

`wails build` will bundle the compiled UI from `internal/ui/dist` into the final executable.

## License

This project is licensed under the [MIT License](LICENSE).
