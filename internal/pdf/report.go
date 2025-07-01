package pdf

import (
	"fmt"
	"os"
	"path/filepath"

	"baristeuer/internal/data"
	"github.com/jung-kurt/gofpdf"
)

// GenerateReport fetches project data and creates a PDF report.
func GenerateReport(projectID string) (string, error) {
	// Fetch data (placeholder implementation)
	project := data.GetProjectData(projectID)

	dir := filepath.Join("data", "reports", projectID)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	file := filepath.Join(dir, "report.pdf")

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, fmt.Sprintf("Report for %s", project.Name))
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 10, project.Info, "", "", false)

	if err := pdf.OutputFileAndClose(file); err != nil {
		return "", err
	}
	return file, nil
}
