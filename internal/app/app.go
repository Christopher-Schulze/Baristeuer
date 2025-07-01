package app

import "baristeuer/internal/pdf"

// App is a Wails binding placeholder.
type App struct{}

// GenerateReport generates a PDF report for the given project.
func (a *App) GenerateReport(projectID string) (string, error) {
	return pdf.GenerateReport(projectID)
}
