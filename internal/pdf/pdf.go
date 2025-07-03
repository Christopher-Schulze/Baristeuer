package pdf

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	ctx := context.Background()
	revenue, err := g.store.SumIncomeByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch revenue: %w", err)
	}
	expenses, err := g.store.SumExpenseByProject(ctx, projectID)
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

// GenerateKSt1 creates a placeholder "KSt 1" form for the given project.
func (g *Generator) GenerateKSt1(projectID int64) (string, error) {
	p, _ := g.store.GetProject(context.Background(), projectID)
	nameLine := "Name des Vereins: ____________________"
	if p != nil {
		nameLine = fmt.Sprintf("Name des Vereins: %s", p.Name)
	}
	lines := []string{
		"K\xC3\xB6rperschaftsteuererkl\xC3\xA4rung f\xC3\xBCr Vereine", // "Körperschaftsteuererklärung für Vereine"
		nameLine,
		"Steuernummer: ____________________",
		"Veranlagungszeitraum: 2025",
		"(Bitte Formular vollständig ausf\xC3\xBCllen)",
	}
	return g.createForm(projectID, "KSt 1", lines)
}

// GenerateAnlageGem creates a placeholder "Anlage Gem" form for the given project.
func (g *Generator) GenerateAnlageGem(projectID int64) (string, error) {
	p, _ := g.store.GetProject(context.Background(), projectID)
	prefix := "Anlage Gem - Angaben zur Gemeinn\xC3\xBCtzigkeit"
	if p != nil {
		prefix = fmt.Sprintf("Anlage Gem - %s", p.Name)
	}
	lines := []string{
		prefix,
		"T\xC3\xA4tigkeit des Vereins: ____________________",
		"Steuerbeg\xC3\xBCnstigte Zwecke: ____________________",
		"(Bitte Formular vollst\xC3\xA4ndig ausf\xC3\xBCllen)",
	}
	return g.createForm(projectID, "Anlage Gem", lines)
}

// GenerateAnlageGK creates a placeholder "Anlage GK" form for the given project.
func (g *Generator) GenerateAnlageGK(projectID int64) (string, error) {
	p, _ := g.store.GetProject(context.Background(), projectID)
	title := "Anlage GK - Angaben zu Gesch\xC3\xA4ftsbetrieben"
	if p != nil {
		title = fmt.Sprintf("Anlage GK - %s", p.Name)
	}
	lines := []string{
		title,
		"Bezeichnung des wirtschaftlichen Gesch\xC3\xA4ftsbetriebs: ____________________",
		"Gewinne/Verluste: ____________________",
		"(Bitte Formular vollst\xC3\xA4ndig ausf\xC3\xBCllen)",
	}
	return g.createForm(projectID, "Anlage GK", lines)
}

// GenerateKSt1F creates a placeholder "KSt 1F" form for the given project.
func (g *Generator) GenerateKSt1F(projectID int64) (string, error) {
	p, _ := g.store.GetProject(context.Background(), projectID)
	title := "KSt 1F - Erweiterte K\xC3\xB6rperschaftsteuererkl\xC3\xA4rung"
	if p != nil {
		title = fmt.Sprintf("KSt 1F - %s", p.Name)
	}
	lines := []string{
		title, // "KSt 1F - Erweiterte Körperschaftsteuererklärung"
		"Angaben zu Beteiligungen: ____________________",
		"Erhaltene F\xC3\xB6rdermittel: ____________________",
		"(Bitte Formular vollst\xC3\xA4ndig ausf\xC3\xBCllen)",
	}
	return g.createForm(projectID, "KSt 1F", lines)
}

// GenerateAnlageSport creates a placeholder "Anlage Sport" form for the given project.
func (g *Generator) GenerateAnlageSport(projectID int64) (string, error) {
	p, _ := g.store.GetProject(context.Background(), projectID)
	title := "Anlage Sport - Sportvereine"
	if p != nil {
		title = fmt.Sprintf("Anlage Sport - %s", p.Name)
	}
	lines := []string{
		title, // heading
		"Mitgliederzahl: ____________________",
		"Einnahmen aus Sportbetrieb: ____________________",
		"(Bitte Formular vollst\xC3\xA4ndig ausf\xC3\xBCllen)",
	}
	return g.createForm(projectID, "Anlage Sport", lines)
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
	return g.createForm(projectID, title, []string{"(Platzhalter)"})
}

// createForm writes a PDF with the given title and content lines.
func (g *Generator) createForm(projectID int64, title string, lines []string) (string, error) {
	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	sanitized := strings.ReplaceAll(title, " ", "_")
	fileName := fmt.Sprintf("%s_%d.pdf", sanitized, projectID)
	filePath := filepath.Join(g.BasePath, fileName)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, title)
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	for _, l := range lines {
		pdf.Cell(0, 8, l)
		pdf.Ln(8)
	}

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return "", fmt.Errorf("failed to write PDF: %w", err)
	}
	return filePath, nil
}
