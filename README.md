# Baristeuer

Baristeuer is a proof-of-concept desktop application for managing **Vereinssteuererklärung** (club tax returns). The project combines a React front end with a Go back end and uses [Wails](https://wails.io/) to produce a desktop application.

## Features

- Desktop client built with [Wails](https://wails.io/)
- React + Material UI interface located under `internal/ui`
- Go back end located under `cmd`
- SQLite persistence layer

## Installation

The application requires **Go**, **Node.js** and **Wails**. The following commands show a minimal setup on Debian/Ubuntu. For other OSes use the official download pages.

### Go

```bash
sudo apt-get install golang-go