package pdf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"

	taxrules "baristeuer/internal/taxrules"
)

// Generator handles PDF creation.
type Generator struct {
	BasePath string
}

// NewGenerator returns a new Generator storing reports under basePath.
func NewGenerator(basePath string) *Generator {
	if basePath == "" {
		basePath = filepath.Join(".", "internal", "data")
	}
	return &Generator{BasePath: basePath}
}

// GenerateReport fetches project data and creates a PDF report.
func (g *Generator) GenerateReport(projectID string) (string, error) {
	// placeholder project data
	p := taxrules.Project{Revenue: 1000, Expenses: 500}

	dir := filepath.Join(g.BasePath, projectID)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("create project dir: %w", err)
	}
	file := filepath.Join(dir, "report.pdf")

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, fmt.Sprintf("Report for project %s", projectID))
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Revenue: %.2f", p.Revenue))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Expenses: %.2f", p.Expenses))
	pdf.Ln(8)
	tax := taxrules.CalculateTax(p)
	pdf.Cell(40, 10, fmt.Sprintf("Tax: %.2f", tax))

	if err := pdf.OutputFileAndClose(file); err != nil {
		return "", fmt.Errorf("write pdf: %w", err)
	}
	return file, nil
}
