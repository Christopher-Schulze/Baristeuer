# Documentation

## Project Overview

Baristeuer is a desktop application for generating tax reports for non-profit organizations in Germany. It uses Go for the backend and a React frontend styled with Tailwind CSS. The application is bundled with Wails so it can run as a native program.

## Technology Stack

### Backend
- **Go**: Language for all backend logic.
- **Wails**: Framework for building the desktop application.
- **SQLite**: Lightweight database used through `database/sql`.

### Frontend
- **React**: UI library for building the interface.
- **Tailwind CSS**: Utility-first CSS framework for styling.
- **Vite**: Development server and build tool.

## Directory Structure

- `cmd/`: Application entry point and Wails setup.
- `docs/`: Documentation files (`Changelog.md`, `DOCUMENTATION.md`, `TODO123.md`).
- `internal/`: Core application logic.
  - `data/`: Data models and SQLite storage.
  - `pdf/`: PDF generation utilities.
  - `taxlogic/`: Business logic for tax calculations for German non-profits.
  - `ui/`: React frontend source code.
- `scripts/`: Helper scripts for development.

## Key Features

- **React + Tailwind Interface**: UI built with React components styled using Tailwind.
- **PDF Generation**: Creates detailed tax reports in PDF format for submission to the German tax office.
- **SQLite Storage**: Simple persistence layer to store project data.
- **Unit Tests**: Tests covering the tax calculation logic.
