<div align="center">
  <img src="logo.png" alt="Bari$teuer Logo" width="300">
</div>

# Bari$teuer - Tax Management for German Associations

Bari$teuer is a software tool designed to assist German non-profit organizations (Vereine) with their tax management. It provides features for financial data management, tax calculation, and report generation based on German tax regulations for non-profits.

---

## Features

- **Tax Calculation:** Calculates corporate tax (Körperschaftsteuer) and VAT (Umsatzsteuer) for non-profit organizations based on German tax laws.
- **Data Management:** Centralized tracking of income, expenses, and donations.
- **Member Management:** Track club members with join date and contact details.
- **Tax Overview:** Calculates taxes for a project and displays the results in the UI.
- **Reporting:** Generates tax reports for submission.
- **User Interface:** German-language interface styled with Material UI themes.
- **Cross-Platform:** Operates on both macOS and Windows.
- **PDF Export:** Exports financial reports and a set of Vereinssteuerformulare (KSt 1, Anlagen Gem/GK usw.) to PDF format.

## Technology Stack

Bari$teuer is built with the following technologies:

- **Backend:** Go
- **Frontend:** React
- **UI Framework:** Material UI
- **Desktop App Framework:** Wails

## Getting Started

To build Bari$teuer locally, clone the repository and run:

```bash
go work sync
npm install --prefix internal/ui
npm run build --prefix internal/ui
wails build -clean
```

The resulting binaries can be found in `build/bin`. Follow the official documentation for platform specific details.

## Testing

1. Run `go work sync`.
2. Execute `go test ./cmd/... ./internal/... ./internal/pdf/...`.

See [docs/DOCUMENTATION.md](docs/DOCUMENTATION.md) for more details.

---

*This project is for internal use and is not open for contributions.*
