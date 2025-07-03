# Changelog

## [3.7.2025] - CI Workflow

### Added
- **CI Workflow**: Set up a GitHub Actions workflow (`.github/workflows/build.yml`) to automatically build and test the application on every push to the `main` branch. The workflow builds for macOS, Windows, and Linux, runs Go and Bun tests, and uploads the compiled binaries as build artifacts.

## [3.7.2025] - Release 1.0

### Added
- **Unit Tests**: Added comprehensive unit tests for the tax calculation logic in `internal/taxlogic/` to ensure correctness and reliability.
- **Tailwind Styling**: Replaced the old Material UI components with Tailwind CSS in the React frontend.
- **Documentation Updates**: Updated `DOCUMENTATION.md` and `TODO123.md` to describe the React/Tailwind stack.

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
  - **Frontend**: Updated the `App.jsx` component to accept user input for revenue and expenses and generate reports via the backend.

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
- Completely revised `README.md` with a new structure, embedded logo, and detailed feature descriptions.

### Removed
- Removed social media and contributor sections from `README.md`.