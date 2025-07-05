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
  "logLevel": "info"
}
```

Command line flags override values from the file. The log level can also be adjusted at runtime via the `SetLogLevel` method exposed by the `DataService`.

### Language Selection

The application starts in German. Use the dropdown in the top toolbar to switch between German and English at any time.

## Key Features

- **React + Material UI Interface**: UI built with React components styled using Material UI.
- **Multi-language Support**: Interface available in German and English. Change the language via the toolbar dropdown.
- **PDF Generation**: Creates detailed tax reports in PDF format for submission to the German tax office.
- **SQLite Storage**: Simple persistence layer to store project data.
- **Member Tracking**: Manage club members with names, emails and join dates.
- **Unit Tests**: Tests covering the tax calculation logic.
- **Backup & Restore**: The `DataService` can export and restore the SQLite database for easy backups. Use `-exportdb <file>` to dump the database or `-restoredb <file>` to load it before the UI starts.
- **CSV Export**: Use the `ExportProjectCSV` method to save all incomes and expenses of a project to a CSV file.

## Cross-Platform Compatibility

The application is required to be fully functional on both **macOS** and **Windows** operating systems. All features must be tested and verified on both platforms to ensure a consistent user experience.

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

The Baristeuer application is stable and covers the core features. The backend, built with Go, provides robust tax calculation and data management logic. The React-based frontend, styled with Material UI, offers an intuitive and responsive user interface. Erste Vorlagen für Steuerformulare (KSt 1, Anlage Gem/GK usw.) werden als PDF erzeugt. Die Anwendung wurde auf macOS und Windows für Basisszenarien getestet.

### Production Readiness

Weitere Funktionen wie vollständige Steuerformulare und detaillierte PDF-Berichte sind noch in Arbeit. Die CI/CD-Pipeline führt Tests aus, dennoch ist das Projekt derzeit als Vorabversion zu betrachten.

### Potential Next Steps

While the application is complete, future enhancements could include:

- **Advanced Reporting:** Adding more detailed analytics and customizable report templates.
- **Cloud Sync:** Implementing an optional cloud synchronization feature for backing up and sharing data across devices.
- **Plugin System:** Developing a plugin architecture to allow for third-party extensions and integrations.
