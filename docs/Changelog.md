# Changelog

## [7.7.2025] - Frontend Migration

### Changed

- Completed the migration to Svelte with DaisyUI, fully replacing the old React and Material UI interface.
- Updated README and documentation to reflect the new stack.

## [6.7.2025] - Documentation Update

### Changed

- Clarified production readiness and cross-platform support in `DOCUMENTATION.md`.


## [5.7.2025] - Build and Configuration Improvements

### Added

- `Dockerfile` für den containerisierten Build.
- `BARISTEUER_PDFDIR` Umgebungsvariable, um das PDF-Ausgabeverzeichnis zu überschreiben.
- Zusätzliche Komponenten- und Servicetests.
- Go-Format-Prüfung innerhalb der CI.
- Fehlerweitergabe beim automatischen Erstellen der Konfigurationsdatei.

### Changed

- Hilfsfunktionen zur Erstellung von PDF-Formularzeilen wurden eingeführt.

## [5.7.2025] - CSV Export

### Added

- `ExportProjectCSV` service method to save project data as CSV
- Settings panel option to trigger CSV export from the UI

## [3.7.2025] - UI Validation and Logging

### Added

- **Cleanup Wails Bindings**: Removed outdated bindings from `cmd/internal/ui/wailsjs`.
- **Input Validation**: Added positive amount checks in the React `App.jsx` component (later migrated to Svelte).
- **PDF Form Fields**: Extended placeholder forms with fill-in lines.
- **Service Logging**: Introduced a basic logger in `internal/service`.

## [3.7.2025] - README Logo Path Fix

### Fixed

- **README.md**: Corrected the image path for the logo to use the relative path (`./logo.png`).

## [3.7.2025] - Path Correction

### Fixed

- **Logo Path**: Corrected the relative path to the logo in `README.md` to ensure it displays correctly.

## [3.7.2025] - README Refinement

### Changed

- **README Revision**: Revised `README.md` to adopt a more neutral and factual tone. Updated the title, embedded a smaller logo, and streamlined the feature descriptions to be more concise.

## [3.7.2025] - Final Review and Documentation

### Added

- **Final Project Assessment**: Added a final assessment to `DOCUMENTATION.md`, confirming the project's production readiness and outlining potential future enhancements.

### Changed

- **TODOs Finalized**: Updated `TODO123.md` to mark all development tasks as complete and revised descriptions to reflect the final project state.
- **README Correction**: Temporarily corrected the frontend technology in `README.md` from "Svelte" to "React" to match the stack at the time. (This was later reverted when migrating to Svelte and DaisyUI.)

## [3.7.2025] - Wails Bindings

### Changed

- **Wails Bindings**: Regenerated the Wails module to expose the new `DataService` methods to the frontend. This involved fixing several compilation and project setup issues.

## [3.7.2025] - Data Handling

### Added

- **In-Memory Data Service**: Implemented a new Go service (`internal/data/service.go`) with `AddExpense` and `ListExpenses` methods using an in-memory slice for data storage.
- **UI for Data Entry**: Added a React `App.jsx` form for income and expenses, later ported to Svelte.

### Changed

- **Service Binding**: Updated `cmd/main.go` to bind the new `DataService`, making it available to the frontend.
- **Frontend Logic**: Refactored the React `App.jsx` to use the new data service methods. This logic was subsequently rewritten in Svelte.

## [3.7.2025] - Documentation

### Changed

- **README Revision**: Completely revised `README.md` to be more comprehensive and engaging. Embedded the project logo, added detailed feature descriptions, and included information about the technology stack.

## [3.7.2025] - CI Workflow

### Added

- **CI Workflow**: Set up a GitHub Actions workflow (`.github/workflows/build.yml`) to automatically build and test the application on every push to the `main` branch. The workflow builds for macOS, Windows, and Linux, runs Go and Bun tests, and uploads the compiled binaries as build artifacts.

## [3.7.2025] - Release 1.0

### Added

- **Unit Tests**: Added comprehensive unit tests for the tax calculation logic in `internal/taxlogic/` to ensure correctness and reliability.
- **Material UI Reinstated (Temporary)**: After experimenting with Tailwind CSS, the React frontend briefly returned to Material UI components. This setup was replaced entirely by Svelte and DaisyUI in the final migration.
- **Documentation Updates**: Updated `DOCUMENTATION.md` and `TODO123.md` to describe the then-current React/Tailwind stack. Documentation was revised again after the move to Svelte.

### Changed

- **Code Refactoring**: Consolidated all tax logic into `tax_logic.go`, ensuring a single source of truth.
- **PDF Enhancement**: Optimized PDF layout, fonts, and labels for improved readability and official compliance.

### Fixed

- **Compilation Error**: Resolved a critical bug in the PDF generator by adding the missing `Timestamp` field to the `TaxResult` struct.
- **File Naming**: Corrected a typo in the logo filename, renaming `loco.png` to `logo.png`.

### Removed

- **Redundant Logic**: Deleted the entire `internal/taxrules` directory, which contained faulty and unused tax calculation logic.
- **Ghost File**: Addressed the non-existent `tax_calculator.go` by correcting the application's control flow.

## [2.7.2025] - Core Feature Implementation

### Added

- **Core Application Logic**: Implemented the core tax calculation and reporting features.
  - **Tax Logic**: Developed the tax calculation engine for German non-profits for the year 2025 in `internal/taxlogic/tax_logic.go`.
  - **PDF Generation**: Enhanced the PDF report in `internal/pdf/pdf.go` to use the new tax logic and formatted it for submission to the German tax office (Finanzamt).
  - **Frontend**: Initially implemented a React `App.jsx` component for collecting revenue and expenses. This was later rewritten in Svelte.

### Changed

- **Refactoring**: Consolidated all tax calculation logic into `internal/taxlogic/tax_logic.go`.
- **API Change**: Modified the `GenerateReport` function in `internal/pdf/pdf.go` to accept `revenue` and `expenses` directly.

### Removed

- Redundant placeholder logic in the tax calculation and PDF generation modules.

## [2.7.2025] - Initial Setup

### Added

- Initial project setup with basic structure and documentation files.
- **Hot Reload Configuration**: Configured Vite for hot reload functionality on macOS.
- **Frontend Build Fix**: Resolved a PostCSS configuration error that prevented the development server from starting.

## [v1.1.0] - 2025-07-03

### Added

- Expanded `DOCUMENTATION.md` to include cross-platform compatibility requirements (macOS and Windows).

### Changed

- Updated the application name to "Bari$teuer" in `wails.json`.
- Set `logo.png` as the application icon in `wails.json`.

### Removed

- Removed social media and contributor sections from `README.md`.
