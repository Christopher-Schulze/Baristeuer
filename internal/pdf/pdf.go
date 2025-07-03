package pdf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"

	"baristeuer/internal/data"
	"baristeuer/internal/taxlogic"
)

// Generator handles PDF creation.
type Generator struct {
	BasePath string
	store    *data.Store
}

// NewGenerator returns a new Generator for storing reports.
func NewGenerator(basePath string, store *data.Store) *Generator {
	if basePath == "" {
		basePath = filepath.Join(".", "internal", "data", "reports")
	}
	return &Generator{BasePath: basePath, store: store}
}

// GenerateReport creates a tax report PDF for the given project.
func (g *Generator) GenerateReport(projectID int64) (string, error) {
	revenue, err := g.store.SumIncomeByProject(projectID)
	if err != nil {
		return "", fmt.Errorf("fetch revenue: %w", err)
	}
	expenses, err := g.store.SumExpenseByProject(projectID)
	if err != nil {
		return "", fmt.Errorf("fetch expenses: %w", err)
	}
	taxResult := taxlogic.CalculateTaxes(revenue, expenses)

	// Ensure the directory exists
	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	fileName := fmt.Sprintf("tax_report_%d.pdf", taxResult.Timestamp)
	filePath := filepath.Join(g.BasePath, fileName)

	// PDF generation
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Steuerbericht 2025 (Gemeinnützige Organisation)")
	pdf.Ln(20)

	// Summary Section
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "1. Zusammenfassung der Finanzen")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(60, 10, "Einnahmen:")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f EUR", taxResult.Revenue))
	pdf.Ln(8)
	pdf.Cell(60, 10, "Ausgaben:")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f EUR", taxResult.Expenses))
	pdf.Ln(8)
	pdf.Cell(60, 10, "Gewinn:")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f EUR", taxResult.Profit))
	pdf.Ln(15)

	// Tax Calculation Details
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "2. Steuerliche Prüfung und Berechnung")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(60, 10, "Umsatzfreigrenze:")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f EUR", taxResult.RevenueExemptionLimit))
	pdf.Ln(8)
	pdf.Cell(60, 10, "Steuerpflicht aktiv:")
	pdf.Cell(0, 10, fmt.Sprintf("%t", taxResult.IsTaxable))
	pdf.Ln(8)
	pdf.Cell(60, 10, "Gewinnfreibetrag:")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f EUR", taxResult.ProfitAllowance))
	pdf.Ln(8)
	pdf.Cell(60, 10, "Steuerpflichtiges Einkommen:")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f EUR", taxResult.TaxableIncome))
	pdf.Ln(15)

	// Final Tax Result
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "3. Finale Steuerlast")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(60, 10, "Körperschaftsteuer (15%):")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f EUR", taxResult.CorporateTax))
	pdf.Ln(8)
	pdf.Cell(60, 10, "Solidaritätszuschlag (5.5%):")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f EUR", taxResult.SolidaritySurcharge))
	pdf.Ln(8)
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(60, 10, "Gesamtsteuer:")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f EUR", taxResult.TotalTax))
	pdf.Ln(10)

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return "", fmt.Errorf("failed to write PDF: %w", err)
	}
	return filePath, nil
}

// GenerateKSt1 creates a placeholder "KSt 1" form for the given project.
func (g *Generator) GenerateKSt1(projectID int64) (string, error) {
	return g.createSimpleForm(projectID, "KSt 1")
}

// GenerateAnlageGem creates a placeholder "Anlage Gem" form for the given project.
func (g *Generator) GenerateAnlageGem(projectID int64) (string, error) {
	return g.createSimpleForm(projectID, "Anlage Gem")
}

// GenerateAnlageGK creates a placeholder "Anlage GK" form for the given project.
func (g *Generator) GenerateAnlageGK(projectID int64) (string, error) {
	return g.createSimpleForm(projectID, "Anlage GK")
}

// GenerateKSt1F creates a placeholder "KSt 1F" form for the given project.
func (g *Generator) GenerateKSt1F(projectID int64) (string, error) {
	return g.createSimpleForm(projectID, "KSt 1F")
}

// GenerateAnlageSport creates a placeholder "Anlage Sport" form for the given project.
func (g *Generator) GenerateAnlageSport(projectID int64) (string, error) {
	return g.createSimpleForm(projectID, "Anlage Sport")
}

// GenerateAllForms creates all available forms for the given project and returns their paths.
func (g *Generator) GenerateAllForms(projectID int64) ([]string, error) {
	forms := []func(int64) (string, error){
		g.GenerateReport,
		g.GenerateKSt1,
		g.GenerateAnlageGem,
		g.GenerateAnlageGK,
		g.GenerateKSt1F,
		g.GenerateAnlageSport,
	}
	var paths []string
	for _, f := range forms {
		p, err := f(projectID)
		if err != nil {
			return nil, err
		}
		paths = append(paths, p)
	}
	return paths, nil
}

// createSimpleForm writes a minimal PDF with the given title.
func (g *Generator) createSimpleForm(projectID int64, title string) (string, error) {
	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	fileName := fmt.Sprintf("%s_%d.pdf", title, projectID)
	filePath := filepath.Join(g.BasePath, fileName)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, title)
	pdf.Ln(20)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "(Platzhalter)")

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return "", fmt.Errorf("failed to write PDF: %w", err)
	}
	return filePath, nil
}
