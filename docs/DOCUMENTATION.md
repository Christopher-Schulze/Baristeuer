# Documentation

## Project Overview

Baristeuer is a desktop application for generating tax reports for non-profit organizations in Germany. It uses Go for the backend and a React frontend styled with Material UI. The application is bundled with Wails so it can run as a native program.

## Technology Stack

### Backend

- **Go**: Language for all backend logic.
- **Wails**: Framework for building the desktop application.
- **SQLite**: Lightweight database used through `database/sql`.

### Frontend

- **React**: UI library for building the interface.
- **Material UI**: Component library for styling.
- **Vite**: Development server and build tool.

## Directory Structure

- `cmd/`: Application entry point and Wails setup.
- `docs/`: Documentation files (`Changelog.md`, `DOCUMENTATION.md`).
- `internal/`: Core application logic.
  - `data/`: Data models and SQLite storage.
  - `pdf/`: PDF generation utilities.
  - `taxlogic/`: Business logic for tax calculations for German non-profits.
  - `ui/`: React frontend source code.
- `scripts/`: Helper scripts for development.

## Installing Dependencies

Install all Go modules and Node packages after cloning the repository:

```bash
go work sync
npm ci --prefix internal/ui
```

These commands set up the workspace modules and fetch the UI dependencies.
Run `make vet` and `make test` afterward to verify everything works locally.

## Workspace Setup and Running Tests

The repository uses a Go **workspace** (`go.work`) to manage multiple modules.
These modules are referenced by their paths:

- `baristeuer` (located in `cmd/`)
- `baristeuer/internal`
- `baristeuer/internal/pdf`

Before running any tests, make sure the workspace modules are synchronized so
dependencies are resolved correctly:

```bash
go work sync
```

Once the workspace is synced, run the test suite from the repository root:

```bash
go test ./cmd/... ./internal/... ./internal/pdf/...
```

Run this command from the repository root after executing `go work sync`. It
ensures all workspace modules are included and prevents missing dependency
errors. Future Go releases may support running `go test ./...` across workspace
modules automatically.

### Frontend Tests

To run the React unit tests you must install dependencies first:

```bash
npm ci --prefix internal/ui
npm test --prefix internal/ui
```

## Configuration

Runtime options can be provided via a JSON configuration file. By default the
application looks for `config.json` in the working directory. Copy
`config.example.json` to `config.json` when running the application for the first
time. Example:

```json
{
  "dbPath": "baristeuer.db",
  "pdfDir": "./reports",
  "logFile": "baristeuer.log",
  "logLevel": "info",
  "cloudUploadURL": "https://example.com/upload",
  "cloudDownloadURL": "https://example.com/download",
  "cloudToken": "my-secret-token"
}
```

Set `cloudUploadURL` and `cloudDownloadURL` to enable HTTP-based sync. The
`cloudToken` value is sent as a bearer token for authentication.

If `pdfDir` is omitted, generated PDFs are stored in `./internal/data/reports`.
Alternatively the environment variable `BARISTEUER_PDFDIR` can override the output directory at runtime.

Command line flags override values from the file. The log level can also be adjusted at runtime via the `SetLogLevel` method exposed by the `DataService`.

### Language Selection

The application starts in German. Use the dropdown in the top toolbar to switch between German and English at any time.

## Key Features

- **React + Material UI Interface**: UI built with React components styled using Material UI.
- **Multi-language Support**: Interface available in German and English. Change the language via the toolbar dropdown.
- **PDF Generation**: Creates detailed tax reports in PDF format for submission to the German tax office.
- **Formularerstellung**: Erzeugt vereinfachte Ausgaben der Formblätter KSt 1, Anlage Gem/GK, KSt 1F und Anlage Sport als PDF.
- **Konfigurierbares PDF-Verzeichnis**: Das Ausgabeziel lässt sich über `pdfDir` in der Konfiguration oder die Umgebungsvariable `BARISTEUER_PDFDIR` steuern.
- **SQLite Storage**: Simple persistence layer to store project data.
- **Member Tracking**: Manage club members with names, emails and join dates.
- **Unit Tests**: Tests covering the tax calculation logic.
- **Backup & Restore**: The `DataService` can export and restore the SQLite database for easy backups. Use `-exportdb <file>` to dump the database or `-restoredb <file>` to load it before the UI starts.
- **CSV Export**: Use the `ExportProjectCSV` method to save all incomes and expenses of a project to a CSV file.
- **Docker-Unterstützung**: Ein `Dockerfile` ermöglicht den containerisierten Build des Projekts.

## Optional Plugins

At startup the application scans the `plugins/` directory in the working
directory. All files with the `.so` extension are loaded as Go plugins. Each
plugin must export a `New` function returning a value that implements the
following interface:

```go
type Plugin interface {
    Init(*service.DataService) error
}
```

`Init` is called after the `DataService` is initialized, allowing the plugin to
register additional functionality. If the `plugins/` directory does not exist,
the application starts normally without loading any extensions.

## Cross-Platform Compatibility

Die Anwendung wurde erfolgreich unter **macOS** und **Windows** getestet. Alle Funktionen stehen auf beiden Plattformen ohne Einschränkungen zur Verfügung.

## Release Workflow

The creation of a new release is automated through GitHub Actions. When a tag is
pushed, `.github/workflows/release.yml` builds the application, signs the
checksums and publishes a GitHub release. The workflow performs the following
steps:

1. Checkout the repository with submodules and set up Go and Node.js.
2. Install the Wails CLI and all dependencies.
3. Execute `scripts/package.sh <tag>` to build binaries for macOS, Windows and
   Linux.
4. Archive the packages and generate a `SHA256SUMS` file.
5. Import the GPG private key from the `GPG_PRIVATE_KEY` secret and sign the
   checksums.
6. Create the GitHub release and upload all artifacts.

To run the packaging step locally or inspect its behaviour, execute
`./scripts/package.sh` from the repository root. The script mirrors the build
logic used in the release workflow.

## Continuous Integration

Every push and pull request triggers a CI workflow defined in
`.github/workflows/ci.yml`. This workflow installs all dependencies, runs the
frontend linter and test suite, and executes `go vet` as well as all Go unit
tests. The CI ensures that code quality checks pass before changes are merged.

## Final Project Assessment

### Current State

The Baristeuer application is stable and covers the core features. The backend, built with Go, provides robust tax calculation and data management logic. The React-based frontend, styled with Material UI, offers an intuitive and responsive user interface. Für Steuerformulare stehen inzwischen einfache PDF-Layouts für KSt 1, Anlage Gem/GK, KSt 1F und Anlage Sport bereit. Die Anwendung wurde auf macOS und Windows für Basisszenarien getestet.

### Production Readiness

Offizielle Layouts und umfangreichere PDF-Berichte sind noch in Arbeit. Die CI/CD-Pipeline führt Tests aus, und der aktuelle Funktionsumfang gilt als produktionsreif.

### Potential Next Steps

While the application is complete, future enhancements could include:

- **Advanced Reporting:** Adding more detailed analytics and customizable report templates.
- **Cloud Sync:** Implementing an optional cloud synchronization feature for backing up and sharing data across devices.
- **Plugin System:** Developing a plugin architecture to allow for third-party extensions and integrations.
