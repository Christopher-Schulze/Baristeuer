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

Running the command after `go work sync` ensures all workspace modules are
included and prevents missing dependency errors. With Go 1.23 or newer you can
also use `go test ./...`, which traverses all modules listed in `go.work`.

### Frontend Tests
To run the React unit tests install dependencies first:

```bash
npm ci --prefix internal/ui
npm test --prefix internal/ui
```

## Key Features

- **React + Material UI Interface**: UI built with React components styled using Material UI.
- **PDF Generation**: Creates detailed tax reports in PDF format for submission to the German tax office.
- **SQLite Storage**: Simple persistence layer to store project data.
- **Member Tracking**: Manage club members with names, emails and join dates.
- **Unit Tests**: Tests covering the tax calculation logic.

## Cross-Platform Compatibility
The application is required to be fully functional on both **macOS** and **Windows** operating systems. All features must be tested and verified on both platforms to ensure a consistent user experience.

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
