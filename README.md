# Baristeuer

Baristeuer is a proof-of-concept desktop application for managing **Vereinssteuererklärung** (club tax returns). The project combines a React front end using Material UI with a Go back end.

## Features

- Desktop client built with [Wails](https://wails.io/) (planned)
- React + Material UI interface located under `internal/ui`
- Go back end located under `cmd`

## Installation

The project requires **Go**, **Node.js** and **Wails**. The following commands
show a minimal setup on Debian/Ubuntu. Other platforms can use the official
download pages.

### Go

```bash
sudo apt-get install golang-go
```

You can also install manually from <https://go.dev/dl/>.

### Node.js

Install Node.js 18 or later via NodeSource:

```bash
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs
```

### Wails

Install the Wails CLI using `go install`:

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

Ensure `$(go env GOPATH)/bin` is on your `PATH` so the `wails` command is
available.

## Development

Run the front end and back end separately when working on the project.

1. Install UI dependencies:
   ```bash
   cd internal/ui && npm install
   ```
2. Start the React development server:
   ```bash
   npm start
   ```
   This serves the UI on <http://localhost:3000>.
3. In another terminal start the Go back end:
   ```bash
   cd cmd && go run .
   ```

With Wails installed you can later run `wails dev` to launch the combined
desktop app in development mode.

## Usage Example

Running the Go program prints a small greeting:

```console
$ go run ./cmd
```

The output looks like:

```
 ____             _     _                          ____ _     ___
| __ )  __ _ _ __(_)___| |_ ___ _   _  ___ _ __   / ___| |   |_ _|
|  _ \ / _` | '__| / __| __/ _ \ | | |/ _ \ '__| | |   | |    | |
| |_) | (_| | |  | \__ \ ||  __/ |_| |  __/ |    | |___| |___ | |
|____/ \__,_|_|  |_|___/\__\___|\__,_|\___|_|     \____|_____|___|
```

## Future Plans

- Implement the tax declaration forms and local data storage
- Package the application via Wails for Windows, macOS and Linux
- Add PDF export and optional cloud sync

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on the development
workflow. In short:

- create feature branches from `main` for your work,
- format Go code with `go fmt` and JavaScript code with Prettier,
- ensure `go build ./...` and `npm run build` succeed before opening a pull
  request.

Contributions and feedback are welcome!
