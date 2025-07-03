package pdf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"

	"baristeuer/internal/data"
	"baristeuer/internal/taxlogic"
)

// Generator handles PDF creation.
type Generator struct {
	BasePath string
	store    *data.Store
}

// formField represents a single field in a tax form.
type formField struct {
	Label string
	Value string
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
	pdf.SetCompression(false)
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

// GenerateKSt1 creates a basic "KSt 1" corporate tax form for the project.
func (g *Generator) GenerateKSt1(projectID int64) (string, error) {
	proj, err := g.store.GetProject(projectID)
	if err != nil {
		return "", fmt.Errorf("fetch project: %w", err)
	}
	fields := []formField{
		{Label: "Finanzamt"},
		{Label: "Steuernummer"},
		{Label: "Name der Körperschaft", Value: proj.Name},
		{Label: "Jahr", Value: fmt.Sprintf("%d", time.Now().Year())},
	}
	return g.createFormWithFields(projectID, "KSt 1 (K\u00f6rperschaftsteuererkl\u00e4rung)", fields)
}

// GenerateAnlageGem creates the "Anlage Gem" form for charitable status.
func (g *Generator) GenerateAnlageGem(projectID int64) (string, error) {
	proj, err := g.store.GetProject(projectID)
	if err != nil {
		return "", fmt.Errorf("fetch project: %w", err)
	}
	fields := []formField{
		{Label: "Name der K\u00f6rperschaft", Value: proj.Name},
		{Label: "Angaben zur Gemeinn\u00fctzigkeit"},
		{Label: "T\u00e4tigkeitsbereich"},
	}
	return g.createFormWithFields(projectID, "Anlage Gem", fields)
}

// GenerateAnlageGK creates the "Anlage GK" form for business operations.
func (g *Generator) GenerateAnlageGK(projectID int64) (string, error) {
	proj, err := g.store.GetProject(projectID)
	if err != nil {
		return "", fmt.Errorf("fetch project: %w", err)
	}
	fields := []formField{
		{Label: "Name der K\u00f6rperschaft", Value: proj.Name},
		{Label: "Angaben zur Gesch\u00e4ftsf\u00fchrung"},
	}
	return g.createFormWithFields(projectID, "Anlage GK", fields)
}

// GenerateKSt1F creates the "KSt 1F" form for tax allocation.
func (g *Generator) GenerateKSt1F(projectID int64) (string, error) {
	proj, err := g.store.GetProject(projectID)
	if err != nil {
		return "", fmt.Errorf("fetch project: %w", err)
	}
	fields := []formField{
		{Label: "Name der K\u00f6rperschaft", Value: proj.Name},
		{Label: "Feststellungszeitraum"},
		{Label: "Anzahl der Beteiligten"},
	}
	return g.createFormWithFields(projectID, "KSt 1F", fields)
}

// GenerateAnlageSport creates the "Anlage Sport" form for sporting activities.
func (g *Generator) GenerateAnlageSport(projectID int64) (string, error) {
	proj, err := g.store.GetProject(projectID)
	if err != nil {
		return "", fmt.Errorf("fetch project: %w", err)
	}
	fields := []formField{
		{Label: "Name der K\u00f6rperschaft", Value: proj.Name},
		{Label: "Sportliche Bet\u00e4tigung"},
	}
	return g.createFormWithFields(projectID, "Anlage Sport", fields)
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

// createFormWithFields generates a simple form with the given fields.
func (g *Generator) createFormWithFields(projectID int64, title string, fields []formField) (string, error) {
	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	fileTitle := strings.ReplaceAll(title, " ", "_")
	fileName := fmt.Sprintf("%s_%d.pdf", fileTitle, projectID)
	filePath := filepath.Join(g.BasePath, fileName)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, title)
	pdf.Ln(15)
	pdf.SetFont("Arial", "", 12)
	for _, f := range fields {
		pdf.Cell(60, 10, f.Label+":")
		pdf.Cell(0, 10, f.Value)
		pdf.Ln(8)
	}

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return "", fmt.Errorf("failed to write PDF: %w", err)
	}
	return filePath, nil
}
