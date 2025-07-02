# Documentation

## Project Overview

This project is a desktop application for generating tax reports for non-profit organizations in Germany. It uses Go for the backend and Svelte for the frontend user interface. The application is built with Wails, which allows creating desktop applications using web technologies.

## Technology Stack

### Backend
- **Go**: Programming language for the backend.
- **Wails**: Framework for building cross-platform desktop apps with Go and web technologies.
- **Gin**: HTTP web framework for Go (used for potential future web-based features).
- **gRPC-Web**: For high-performance communication between the frontend and backend.

### Frontend
- **Svelte**: UI framework for building the user interface.
- **Bun**: Fast JavaScript runtime, used instead of Node.js.
- **Tailwind CSS**: Utility-first CSS framework for styling.
- **shadcn/ui**: A collection of re-usable UI components.
- **Dexie.js**: A wrapper for IndexedDB for client-side storage and state management.

## Directory Structure

- `cmd/`: Contains the main application entry point.
- `docs/`: Contains documentation files (`Changelog.md`, `DOCUMENTATION.md`, `TODO123.md`).
- `internal/`: Contains the core application logic.
  - `data/`: Data models and storage logic.
  - `pdf/`: Functions for PDF generation.
  - `taxlogic/`: Business logic for tax calculations, compliant with German tax law for non-profits (2025).
  - `ui/`: Svelte frontend source code.
- `scripts/`: Helper scripts for development and build processes.

## Key Decisions & Changes

- **Frontend Framework**: The entire frontend has been migrated from React to Svelte to align with the specified technology stack.
- **UI Components**: MUI has been completely replaced with Tailwind CSS and shadcn/ui for a more modern and consistent look and feel.
- **State Management**: Dexie.js has been integrated for robust client-side data management.
- **Backend Communication**: Communication between the frontend and backend now uses gRPC-Web for better performance and type safety.
- **Code Consolidation**: The tax calculation logic was centralized in `internal/taxlogic/`. The redundant and faulty `internal/taxrules` directory was removed to eliminate ambiguity and potential for error.
- **PDF Generation**: The PDF report generation was enhanced to be fully compliant with the requirements of the German tax office (Finanzamt), with improved formatting and clearer labeling.
- **Unit Tests**: Comprehensive unit tests for the tax calculation logic have been added to ensure correctness and reliability.