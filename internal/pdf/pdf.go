package pdf

import (
	"context"
	"fmt"
	"log/slog"
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
		if env := os.Getenv("BARISTEUER_PDFDIR"); env != "" {
			basePath = env
		} else {
			basePath = filepath.Join(".", "internal", "data", "reports")
		}
	}
	return &Generator{BasePath: basePath, store: store}
}

// GenerateReport creates a tax report PDF for the given project.
func (g *Generator) GenerateReport(projectID int64) (string, error) {
	ctx := context.Background()
	revenue, err := g.store.SumIncomeByProject(ctx, projectID)
	if err != nil {
		slog.Error("fetch revenue", "err", err, "project", projectID)
		return "", fmt.Errorf("fetch revenue: %w", err)
	}
	expenses, err := g.store.SumExpenseByProject(ctx, projectID)
	if err != nil {
		slog.Error("fetch expenses", "err", err, "project", projectID)
		return "", fmt.Errorf("fetch expenses: %w", err)
	}
	taxResult := taxlogic.CalculateTaxes(revenue, expenses)

	// Ensure the directory exists
	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		slog.Error("create report dir", "err", err, "path", g.BasePath)
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
		slog.Error("write report", "err", err, "path", filePath)
		return "", fmt.Errorf("failed to write PDF: %w", err)
	}
	return filePath, nil
}

// GenerateKSt1 creates a simplified "KSt 1" form for the given project with
// layout similar to the official template. The content here is intentionally
// generic but demonstrates how fields would be positioned in a real form.
func (g *Generator) GenerateKSt1(projectID int64) (string, error) {
	ctx := context.Background()
	p, _ := g.store.GetProject(ctx, projectID)

	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		slog.Error("create kst1 dir", "err", err, "path", g.BasePath)
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	fileName := fmt.Sprintf("KSt_1_%d.pdf", projectID)
	filePath := filepath.Join(g.BasePath, fileName)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "KSt 1 - K\xC3\xB6rperschaftsteuererkl\xC3\xA4rung")
	pdf.Ln(12)

	name := "____________________"
	if p != nil {
		name = p.Name
	}

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "1. Angaben zur K\xC3\xB6rperschaft")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 12)

	pdf.Cell(60, 8, "Name des Vereins:")
	pdf.Cell(0, 8, name)
	pdf.Ln(8)
	pdf.Cell(60, 8, "Rechtsform:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)
	pdf.Cell(60, 8, "Steuernummer:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)
	pdf.Cell(60, 8, "Finanzamt:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)
	pdf.Cell(60, 8, "Straße, Hausnummer:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)
	pdf.Cell(60, 8, "PLZ, Ort:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)
	pdf.Cell(60, 8, "Vertreten durch:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)
	pdf.Cell(60, 8, "Veranlagungszeitraum:")
	pdf.Cell(0, 8, "2025")
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "2. Weitere Angaben")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(60, 8, "Beg\xC3\xBCnstigte Zwecke:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)
	pdf.Cell(60, 8, "T\xC3\xA4tigkeitsbereich:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)
	pdf.Cell(60, 8, "Satzungsdatum:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(10)

	pdf.MultiCell(0, 6, "Alle Angaben sind gem\xC3\xA4\xC3\x9F den Vorgaben der Finanzverwaltung zu machen.", "", "L", false)

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		slog.Error("write kst1", "err", err, "path", filePath)
		return "", fmt.Errorf("failed to write PDF: %w", err)
	}
	return filePath, nil
}

// GenerateAnlageGem creates a simplified "Anlage Gem" form. It mirrors the
// structure of the official form but uses generic placeholder fields.
func (g *Generator) GenerateAnlageGem(projectID int64) (string, error) {
	ctx := context.Background()
	p, _ := g.store.GetProject(ctx, projectID)

	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		slog.Error("create gem dir", "err", err, "path", g.BasePath)
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	fileName := fmt.Sprintf("Anlage_Gem_%d.pdf", projectID)
	filePath := filepath.Join(g.BasePath, fileName)

	title := "Anlage Gem - Angaben zur Gemeinn\xC3\xBCtzigkeit"
	if p != nil {
		title = fmt.Sprintf("Anlage Gem - %s", p.Name)
	}

	name := "____________________"
	if p != nil {
		name = p.Name
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, title)
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(60, 8, "Name des Vereins:")
	pdf.Cell(0, 8, name)
	pdf.Ln(8)
	pdf.Cell(60, 8, "Anschrift des Vereins:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)

	pdf.Cell(60, 8, "T\xC3\xA4tigkeit des Vereins:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)

	pdf.Cell(60, 8, "Steuerbeg\xC3\xBCnstigte Zwecke:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)
	pdf.Cell(60, 8, "Vertreten durch:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)

	pdf.Cell(60, 8, "Verwendung der Mittel:")
	pdf.Cell(0, 8, "____________________")
	pdf.Ln(8)
	pdf.Cell(60, 8, "Bankverbindung:")
	pdf.Cell(0, 8, "____________________")

	pdf.Ln(10)
	pdf.MultiCell(0, 6, "Bitte Formular vollst\xC3\xA4ndig ausf\xC3\xBCllen und dem KSt 1 beif\xC3\xBCgen.", "", "L", false)

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		slog.Error("write gem", "err", err, "path", filePath)
		return "", fmt.Errorf("failed to write PDF: %w", err)
	}
	return filePath, nil
}

// GenerateAnlageGK creates a placeholder "Anlage GK" form for the given project.
func (g *Generator) GenerateAnlageGK(projectID int64) (string, error) {
	ctx := context.Background()
	p, _ := g.store.GetProject(ctx, projectID)
	title := "Anlage GK - Angaben zu Gesch\xC3\xA4ftsbetrieben"
	if p != nil {
		title = fmt.Sprintf("Anlage GK - %s", p.Name)
	}
	lines := []string{
		title,
		"Bezeichnung des wirtschaftlichen Gesch\xC3\xA4ftsbetriebs: ____________________",
		"Gewinne/Verluste: ____________________",
		"Umsatz des Vorjahres: ____________________",
		"(Bitte Formular vollst\xC3\xA4ndig ausf\xC3\xBCllen)",
	}
	return g.createForm(projectID, "Anlage GK", lines)
}

// GenerateKSt1F creates a placeholder "KSt 1F" form for the given project.
func (g *Generator) GenerateKSt1F(projectID int64) (string, error) {
	ctx := context.Background()
	p, _ := g.store.GetProject(ctx, projectID)
	title := "KSt 1F - Erweiterte K\xC3\xB6rperschaftsteuererkl\xC3\xA4rung"
	if p != nil {
		title = fmt.Sprintf("KSt 1F - %s", p.Name)
	}
	lines := []string{
		title, // "KSt 1F - Erweiterte Körperschaftsteuererklärung"
		"Angaben zu Beteiligungen: ____________________",
		"Beteiligungen an Kapitalgesellschaften: ____________________",
		"Erhaltene F\xC3\xB6rdermittel: ____________________",
		"(Bitte Formular vollst\xC3\xA4ndig ausf\xC3\xBCllen)",
	}
	return g.createForm(projectID, "KSt 1F", lines)
}

// GenerateAnlageSport creates a placeholder "Anlage Sport" form for the given project.
func (g *Generator) GenerateAnlageSport(projectID int64) (string, error) {
	ctx := context.Background()
	p, _ := g.store.GetProject(ctx, projectID)
	title := "Anlage Sport - Sportvereine"
	if p != nil {
		title = fmt.Sprintf("Anlage Sport - %s", p.Name)
	}
	lines := []string{
		title, // heading
		"Mitgliederzahl: ____________________",
		"Einnahmen aus Sportbetrieb: ____________________",
		"Anzahl der Übungsleiter: ____________________",
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
			slog.Error("generate form", "err", err)
			return nil, err
		}
		paths = append(paths, p)
	}
	return paths, nil
}

// createForm writes a PDF with the given title and content lines.
func (g *Generator) createForm(projectID int64, title string, lines []string) (string, error) {
	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		slog.Error("create form dir", "err", err, "path", g.BasePath)
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
		slog.Error("write form", "err", err, "path", filePath)
		return "", fmt.Errorf("failed to write PDF: %w", err)
	}
	return filePath, nil
}
