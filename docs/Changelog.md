# Changelog

## [3.7.2025] - Release 1.0

### Added
- **Unit Tests**: Added comprehensive unit tests for the tax calculation logic in `internal/taxlogic/` to ensure correctness and reliability.
- **Final Documentation**: Updated `DOCUMENTATION.md` and `TODO123.md` to reflect the final state of the project.

### Changed
- **Code Refactoring**: Consolidated all tax logic into `tax_logic.go`, ensuring a single source of truth.
- **PDF Enhancement**: Optimized PDF layout, fonts, and labels for improved readability and official compliance.
- **Project Finalization**: Marked all tasks in `TODO123.md` as complete.

### Fixed
- **Compilation Error**: Resolved a critical bug in the PDF generator by adding the missing `Timestamp` field to the `TaxResult` struct.

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