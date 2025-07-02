# Baristeuer

Baristeuer is a proof-of-concept desktop application for managing **Vereinssteuererklärung** (club tax returns). The project combines a React front end using Material UI with a Go back end.

## Features

- Desktop client built with [Wails](https://wails.io/)
- React + Material UI interface located under `internal/ui`
- Go back end located under `cmd`

## Installation

The project requires **Go**, **Node.js** and **Wails**. The following commands show a minimal setup on Debian/Ubuntu. Other platforms can use the official download pages.

### Go

```bash
sudo apt-get install golang-go

### Database

The persistence layer relies on SQLite via the
`github.com/mattn/go-sqlite3` driver. Ensure SQLite development headers are
available (e.g. `libsqlite3-dev` on Debian/Ubuntu).
