# Documentation

## Project Overview

Baristeuer is a desktop application for generating tax reports for non-profit organizations in Germany. It uses Go for the backend and a Svelte frontend styled with DaisyUI. The application is bundled with Wails so it can run as a native program.

Cloud sync via HTTP endpoints is now implemented to back up and restore the SQLite database. Configuration details can be found in [docs/cloud-sync.md](cloud-sync.md).

## Technology Stack

### Backend

- **Go**: Language for all backend logic.
- **Wails**: Framework for building the desktop application.
- **SQLite**: Lightweight database used through `database/sql`.

### Frontend

- **Svelte**: UI library for building the interface.
- **DaisyUI**: Component library for styling.
- **Vite**: Development server and build tool.

## Directory Structure

- `cmd/`: Application entry point and Wails setup.
- `docs/`: Documentation files (`Changelog.md`, `DOCUMENTATION.md`).
- `internal/`: Core application logic.
  - `data/`: Data models and SQLite storage.
  - `pdf/`: PDF generation utilities.
  - `taxlogic/`: Business logic for tax calculations for German non-profits.
  - `ui/`: Svelte frontend source code.
- `scripts/`: Helper scripts for development.

## Schnellstart

1. Repository klonen und die Abh\xC3\xA4ngigkeiten installieren:

```bash
go work sync
npm ci --prefix internal/ui
```

2. Die Benutzeroberfl\xC3\xA4che bauen und das Binary erzeugen:

```bash
npm run build --prefix internal/ui
wails build -clean
```

3. Eine eigene Konfiguration anlegen:

```bash
cp config.example.json config.json
```

Pfade wie `dbPath` oder `pdfDir` k\xC3\xB6nnen in dieser Datei angepasst werden. Starten Sie die Anwendung anschlie\xC3\x9Fend mit:

```bash
./baristeuer -config config.json
```

4. Typische Abl\xC3\xA4ufe innerhalb der Anwendung:
   - Einnahmen und Ausgaben erfassen
   - Mitglieder verwalten
   - \"Bericht erstellen\" w\xC3\xA4hlen, um PDF-Formulare zu generieren
   - Optional CSV-Exporte oder Datenbank-Backups ausf\xC3\xBChren
5. Die offiziellen Formularvorlagen separat herunterladen und in
   `internal/pdf/templates/` ablegen. Ohne diese Dateien zeigt das Programm
   einen Hinweis in den erzeugten PDFs. Lege die amtlichen PDFs exakt so ab,
   wie es in `internal/pdf/templates/README.md` beschrieben ist: Benenne sie
   wie die vorhandenen Platzhalterdateien (`kst1.txt`, `kst1f.txt`,
   `anlage_gem.txt`, `anlage_gk.txt`, `anlage_sport.txt`) und ersetze die
   Endung mit `.pdf`. Die Dateien müssen also `kst1.pdf`, `kst1f.pdf`,
   `anlage_gem.pdf`, `anlage_gk.pdf` und `anlage_sport.pdf` heißen. Da die
   Originalformulare urheberrechtlich geschützt sind, dürfen sie nicht ins
   Repository eingecheckt werden.

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

To run the Svelte unit tests you must install dependencies first:

```bash
npm ci --prefix internal/ui
npm test --prefix internal/ui
```
The `npm` command and the Playwright browsers must be available in your
environment. The Makefile target `make ui-test` also expects these
dependencies to be installed.

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

Set `cloudUploadURL` and `cloudDownloadURL` to enable the built-in
HTTP-based cloud sync. The `cloudToken` value is sent as a bearer token for
authentication. See [cloud-sync.md](cloud-sync.md) for detailed setup
instructions.

If `pdfDir` is omitted, generated PDFs are stored in `./internal/data/reports`.
Alternatively the environment variable `BARISTEUER_PDFDIR` can override the output directory at runtime.

Command line flags override values from the file. The log level can also be adjusted at runtime via the `SetLogLevel` method exposed by the `DataService`.

### Language Selection

The application starts in German. Use the dropdown in the top toolbar to switch between German and English at any time.

## Key Features

- **Svelte + DaisyUI Interface**: UI built with Svelte components styled using DaisyUI.
- **Multi-language Support**: Interface available in German and English. Change the language via the toolbar dropdown.
- **PDF Generation**: Creates detailed tax reports in PDF format for submission to the German tax office.
- **Formularerstellung**: Erzeugt vereinfachte Ausgaben der Formblätter KSt 1, Anlage Gem/GK, KSt 1F und Anlage Sport als PDF.
- **Konfigurierbares PDF-Verzeichnis**: Das Ausgabeziel lässt sich über `pdfDir` in der Konfiguration oder die Umgebungsvariable `BARISTEUER_PDFDIR` steuern.
- **SQLite Storage**: Simple persistence layer to store project data.
- **Member Tracking**: Manage club members with names, emails and join dates.
- **Unit Tests**: Tests covering the tax calculation logic.
- **Backup & Restore**: The `DataService` can export and restore the SQLite database for easy backups. Use `-exportdb <file>` to dump the database or `-restoredb <file>` to load it before the UI starts.
- **CSV Export**: Use the `ExportProjectCSV` method or pass `-exportcsv <projectID>:<file>` when starting the application to save all incomes and expenses of a project. Example:
`./baristeuer -exportcsv 1:report.csv`
- **Cloud Sync**: Built-in synchronization via HTTP endpoints for backing up
  and restoring the database. See [cloud-sync.md](cloud-sync.md) for
  configuration details.
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

The Baristeuer application is stable and covers the core features. The backend, built with Go, provides robust tax calculation and data management logic. The Svelte-based frontend, styled with DaisyUI, offers an intuitive and responsive user interface. Für Steuerformulare stehen inzwischen einfache PDF-Layouts für KSt 1, Anlage Gem/GK, KSt 1F und Anlage Sport bereit. Die Anwendung wurde auf macOS und Windows für Basisszenarien getestet.

### Production Readiness

Die amtlichen Formularlayouts dürfen aus Lizenzgründen nicht im Repository liegen. Laden Sie die Vorlagen selbst herunter und speichern Sie sie im Ordner `internal/pdf/templates/` (siehe [docs/Todo-fuer-User.md](Todo-fuer-User.md)). Die CI/CD-Pipeline führt Tests aus und der aktuelle Funktionsumfang gilt als produktionsreif.

### Potential Next Steps

While the application is complete, future enhancements could include:

- **Advanced Reporting:** Adding more detailed analytics and customizable report templates.
- **Cloud Sync Enhancements:** Improving the implemented synchronization with conflict handling or scheduled backups.
- **Plugin System:** Developing a plugin architecture to allow for third-party extensions and integrations.
