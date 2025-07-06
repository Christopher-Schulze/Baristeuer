package pdf

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jung-kurt/gofpdf"
	gofpdi "github.com/jung-kurt/gofpdf/contrib/gofpdi"

	"baristeuer/internal/config"
	"baristeuer/internal/data"
	"baristeuer/internal/taxlogic"
)

// Error returned when a PDF file cannot be created or written.
var ErrWritePDF = errors.New("write PDF")

const templatesDir = "internal/pdf/templates"

func addTemplate(pdf *gofpdf.Fpdf, name string) {
	pdfPath := filepath.Join(templatesDir, name+".pdf")
	if _, err := os.Stat(pdfPath); err == nil {
		imp := gofpdi.NewImporter()
		tpl := imp.ImportPage(pdf, pdfPath, 1, "/MediaBox")
		imp.UseImportedTemplate(pdf, tpl, 0, 0, 210, 297)
		return
	}
	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(200, 0, 0)
	pdf.MultiCell(0, 4,
		fmt.Sprintf("Vorlage %s.pdf nicht gefunden. Bitte Original in %s ablegen.", name, templatesDir),
		"", "L", false)
	pdf.SetTextColor(0, 0, 0)
	metaPath := filepath.Join(templatesDir, name+".txt")
	if b, err := os.ReadFile(metaPath); err == nil {
		pdf.Ln(2)
		pdf.MultiCell(0, 4, string(b), "", "L", false)
		pdf.Ln(4)
	}
}

// Generator handles PDF creation.
type Generator struct {
	BasePath string
	store    *data.Store
	cfg      *config.Config
}

// FormInfo contains data to fill the various tax forms.
type FormInfo struct {
	Name        string
	TaxNumber   string
	Address     string
	City        string
	BankAccount string
	Activity    string
	FiscalYear  string
}

// Validate checks if the required fields are set and values are plausible.
func (f FormInfo) Validate() error {
	if f.Name == "" {
		return fmt.Errorf("name is required")
	}
	if f.TaxNumber == "" {
		return fmt.Errorf("tax number is required")
	}
	if f.Address == "" {
		return fmt.Errorf("address is required")
	}
	if f.FiscalYear != "" {
		year, err := strconv.Atoi(f.FiscalYear)
		if err != nil || year < 1900 || year > 2100 {
			return fmt.Errorf("invalid fiscal year")
		}
	}
	return nil
}

// NewGenerator returns a new Generator for storing reports.
func NewGenerator(basePath string, store *data.Store, cfg *config.Config) *Generator {
	if cfg == nil {
		cfg = &config.Config{}
	}
	if basePath == "" {
		if env := os.Getenv("BARISTEUER_PDFDIR"); env != "" {
			basePath = env
		} else if cfg.PDFDir != "" {
			basePath = cfg.PDFDir
		} else {
			basePath = config.DefaultConfig.PDFDir
		}
	}
	return &Generator{BasePath: basePath, store: store, cfg: cfg}
}

// SetTaxYear updates the active tax year used for calculations and forms.
func (g *Generator) SetTaxYear(year int) {
	g.cfg.TaxYear = year
}

func (g *Generator) formInfo() FormInfo {
	return FormInfo{
		Name:       g.cfg.FormName,
		TaxNumber:  g.cfg.FormTaxNumber,
		Address:    g.cfg.FormAddress,
		FiscalYear: fmt.Sprintf("%d", g.cfg.TaxYear),
	}
}

// addField writes a single label/value row at the current Y position and
// advances the cursor for the next line.
func addField(pdf *gofpdf.Fpdf, label, value string) {
	const lineH = 8.0
	pdf.CellFormat(70, lineH, label, "1", 0, "L", false, 0, "")
	pdf.CellFormat(110, lineH, value, "1", 0, "L", false, 0, "")
	pdf.Ln(lineH)
}

// tableHeader writes a single header row for a table.
func tableHeader(pdf *gofpdf.Fpdf, widths []float64, headers []string) {
	pdf.SetFont("Arial", "B", 11)
	for i, h := range headers {
		pdf.CellFormat(widths[i], 7, h, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)
	pdf.SetFont("Arial", "", 11)
}

// tableRow writes a row for a table using the provided column widths.
func tableRow(pdf *gofpdf.Fpdf, widths []float64, values []string) {
	for i, v := range values {
		pdf.CellFormat(widths[i], 7, v, "1", 0, "L", false, 0, "")
	}
	pdf.Ln(-1)
}

// GenerateReport creates a tax report PDF for the given project.
func (g *Generator) GenerateReport(ctx context.Context, projectID int64) (string, error) {
	revenue, err := g.store.SumIncomeByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch revenue: %w", err)
	}
	expenses, err := g.store.SumExpenseByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch expenses: %w", err)
	}
	year := g.cfg.TaxYear
	if year == 0 {
		year = 2025
	}
	taxResult := taxlogic.CalculateTaxes(revenue, expenses, year)

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
	pdf.Cell(0, 10, fmt.Sprintf("Steuerbericht %d (Gemeinnützige Organisation)", year))
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
		return "", fmt.Errorf("%w: %v", ErrWritePDF, err)
	}
	return filePath, nil
}

// GenerateDetailedReport creates a tax report with additional statistics like average income and expenses.
func (g *Generator) GenerateDetailedReport(ctx context.Context, projectID int64) (string, error) {
	incomes, err := g.store.ListIncomes(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch incomes: %w", err)
	}
	expensesList, err := g.store.ListExpenses(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch expenses: %w", err)
	}

	var revenue, expenses float64
	for _, inc := range incomes {
		revenue += inc.Amount
	}
	for _, exp := range expensesList {
		expenses += exp.Amount
	}
	var avgIncome, avgExpense float64
	if len(incomes) > 0 {
		avgIncome = revenue / float64(len(incomes))
	}
	if len(expensesList) > 0 {
		avgExpense = expenses / float64(len(expensesList))
	}

	year := g.cfg.TaxYear
	if year == 0 {
		year = 2025
	}
	taxResult := taxlogic.CalculateTaxes(revenue, expenses, year)

	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	fileName := fmt.Sprintf("tax_detailed_report_%d.pdf", taxResult.Timestamp)
	filePath := filepath.Join(g.BasePath, fileName)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, fmt.Sprintf("Detaillierter Steuerbericht %d", year))
	pdf.Ln(20)

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
	pdf.Ln(8)
	pdf.Cell(60, 10, "Ø Einnahmen pro Eintrag:")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f EUR", avgIncome))
	pdf.Ln(8)
	pdf.Cell(60, 10, "Ø Ausgaben pro Eintrag:")
	pdf.Cell(0, 10, fmt.Sprintf("%.2f EUR", avgExpense))
	pdf.Ln(15)

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
		return "", fmt.Errorf("%w: %v", ErrWritePDF, err)
	}
	return filePath, nil
}

// GenerateKSt1 creates a simplified "KSt 1" form for the given project with
// layout similar to the official template. The content here is intentionally
// generic but demonstrates how fields would be positioned in a real form.
func (g *Generator) GenerateKSt1(ctx context.Context, projectID int64) (string, error) {
	p, _ := g.store.GetProject(ctx, projectID)
	info := g.formInfo()
	if info.Name == "" && p != nil {
		info.Name = p.Name
	}
	if info.FiscalYear == "" {
		info.FiscalYear = "2025"
	}
	if err := info.Validate(); err != nil {
		return "", err
	}

	revenue, err := g.store.SumIncomeByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch revenue: %w", err)
	}
	expenses, err := g.store.SumExpenseByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch expenses: %w", err)
	}

	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	fileName := fmt.Sprintf("KSt_1_%d.pdf", projectID)
	filePath := filepath.Join(g.BasePath, fileName)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	pdf.AddPage()
	addTemplate(pdf, "kst1")
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 10, "KSt 1 - K\xC3\xB6rperschaftsteuererkl\xC3\xA4rung", "", 1, "C", false, 0, "")
	pdf.Ln(4)

	name := info.Name

	widths := []float64{70, 110}
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 8, "1. Angaben zur K\xC3\xB6rperschaft", "", 1, "", false, 0, "")
	tableHeader(pdf, widths, []string{"Angabe", "Wert"})
	tableRow(pdf, widths, []string{"Name des Vereins", name})
	tableRow(pdf, widths, []string{"Rechtsform", ""})
	tableRow(pdf, widths, []string{"Steuernummer", info.TaxNumber})
	tableRow(pdf, widths, []string{"Finanzamt", ""})
	tableRow(pdf, widths, []string{"Straße, Hausnummer", info.Address})
	tableRow(pdf, widths, []string{"PLZ, Ort", info.City})
	tableRow(pdf, widths, []string{"Bankverbindung", info.BankAccount})
	tableRow(pdf, widths, []string{"Vertreten durch", ""})
	tableRow(pdf, widths, []string{"Veranlagungszeitraum", info.FiscalYear})
	pdf.Ln(6)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "2. Weitere Angaben")
	pdf.Ln(8)
	tableHeader(pdf, widths, []string{"Angabe", "Wert"})
	tableRow(pdf, widths, []string{"Beg\xC3\xBCnstigte Zwecke", ""})
	tableRow(pdf, widths, []string{"T\xC3\xA4tigkeitsbereich", info.Activity})
	tableRow(pdf, widths, []string{"Satzungsdatum", ""})
	pdf.Ln(6)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "3. Finanz\xC3\xBCbersicht")
	pdf.Ln(8)
	tableHeader(pdf, widths, []string{"Posten", "Betrag"})
	tableRow(pdf, widths, []string{"Einnahmen gesamt", fmt.Sprintf("%.2f EUR", revenue)})
	tableRow(pdf, widths, []string{"Ausgaben gesamt", fmt.Sprintf("%.2f EUR", expenses)})
	tableRow(pdf, widths, []string{"Jahres\xC3\xBCberschuss", fmt.Sprintf("%.2f EUR", revenue-expenses)})
	pdf.Ln(8)

	pdf.MultiCell(0, 6, "Alle Angaben sind gem\xC3\xA4\xC3\x9F den Vorgaben der Finanzverwaltung zu machen.", "", "L", false)

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return "", fmt.Errorf("%w: %v", ErrWritePDF, err)
	}
	return filePath, nil
}

// GenerateAnlageGem creates a simplified "Anlage Gem" form. It mirrors the
// structure of the official form but uses generic placeholder fields.

func (g *Generator) GenerateAnlageGem(ctx context.Context, projectID int64) (string, error) {
	p, _ := g.store.GetProject(ctx, projectID)
	info := g.formInfo()
	if info.Name == "" && p != nil {
		info.Name = p.Name
	}
	if info.FiscalYear == "" {
		info.FiscalYear = "2025"
	}
	if err := info.Validate(); err != nil {
		return "", err
	}

	revenue, err := g.store.SumIncomeByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch revenue: %w", err)
	}
	expenses, err := g.store.SumExpenseByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch expenses: %w", err)
	}
	members, err := g.store.ListMembers(ctx)
	if err != nil {
		return "", fmt.Errorf("fetch members: %w", err)
	}
	memberCount := len(members)

	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	fileName := fmt.Sprintf("Anlage_Gem_%d.pdf", projectID)
	filePath := filepath.Join(g.BasePath, fileName)

	title := "Anlage Gem - Angaben zur Gemeinn\xC3\xBCtzigkeit"
	if p != nil {
		title = fmt.Sprintf("Anlage Gem - %s", p.Name)
	}

	name := info.Name

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	pdf.AddPage()
	addTemplate(pdf, "anlage_gem")
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 10, title, "", 1, "C", false, 0, "")
	pdf.Ln(4)

	widths := []float64{70, 110}

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "1. Vereinsdaten")
	pdf.Ln(8)
	tableHeader(pdf, widths, []string{"Angabe", "Wert"})
	tableRow(pdf, widths, []string{"Name des Vereins", name})
	tableRow(pdf, widths, []string{"Steuernummer", info.TaxNumber})
	tableRow(pdf, widths, []string{"Anschrift", info.Address})
	tableRow(pdf, widths, []string{"Ort", info.City})
	tableRow(pdf, widths, []string{"Bankverbindung", info.BankAccount})
	tableRow(pdf, widths, []string{"T\xC3\xA4tigkeitsbereich", info.Activity})
	pdf.Ln(6)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "2. Angaben zur Gemeinn\xC3\xBCtzigkeit")
	pdf.Ln(8)
	tableHeader(pdf, widths, []string{"Angabe", "Wert"})
	tableRow(pdf, widths, []string{"Steuerbeg\xC3\xBCnstigte Zwecke", ""})
	tableRow(pdf, widths, []string{"Vertreten durch", ""})
	tableRow(pdf, widths, []string{"Verwendung der Mittel", ""})
	tableRow(pdf, widths, []string{"Mitglieder", fmt.Sprintf("%d", memberCount)})
	tableRow(pdf, widths, []string{"Einnahmen", fmt.Sprintf("%.2f EUR", revenue)})
	tableRow(pdf, widths, []string{"Ausgaben", fmt.Sprintf("%.2f EUR", expenses)})
	pdf.Ln(8)

	pdf.Ln(10)
	pdf.MultiCell(0, 6, "Bitte Formular vollst\xC3\xA4ndig ausf\xC3\xBCllen und dem KSt 1 beif\xC3\xBCgen.", "", "L", false)

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return "", fmt.Errorf("%w: %v", ErrWritePDF, err)
	}
	return filePath, nil
}

// GenerateAnlageGK creates a placeholder "Anlage GK" form for the given project.
func (g *Generator) GenerateAnlageGK(ctx context.Context, projectID int64) (string, error) {
	p, _ := g.store.GetProject(ctx, projectID)
	info := g.formInfo()
	if info.Name == "" && p != nil {
		info.Name = p.Name
	}
	if err := info.Validate(); err != nil {
		return "", err
	}

	revenue, err := g.store.SumIncomeByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch revenue: %w", err)
	}
	expenses, err := g.store.SumExpenseByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch expenses: %w", err)
	}

	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	fileName := fmt.Sprintf("Anlage_GK_%d.pdf", projectID)
	filePath := filepath.Join(g.BasePath, fileName)

	title := "Anlage GK - Angaben zu Gesch\xC3\xA4ftsbetrieben"
	if p != nil {
		title = fmt.Sprintf("Anlage GK - %s", p.Name)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	pdf.AddPage()
	addTemplate(pdf, "anlage_gk")
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 10, title, "", 1, "C", false, 0, "")
	pdf.Ln(4)

	widths := []float64{70, 110}
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "1. Grunddaten")
	pdf.Ln(8)
	tableHeader(pdf, widths, []string{"Angabe", "Wert"})
	tableRow(pdf, widths, []string{"Name des Vereins", info.Name})
	tableRow(pdf, widths, []string{"Steuernummer", info.TaxNumber})
	tableRow(pdf, widths, []string{"Anschrift", info.Address})
	tableRow(pdf, widths, []string{"Ort", info.City})
	tableRow(pdf, widths, []string{"Bankverbindung", info.BankAccount})
	tableRow(pdf, widths, []string{"T\xC3\xA4tigkeitsbereich", info.Activity})
	tableRow(pdf, widths, []string{"Bezeichnung Gesch\xC3\xA4ftsbetrieb", ""})
	tableRow(pdf, widths, []string{"Gewinne/Verluste", ""})
	tableRow(pdf, widths, []string{"Umsatz des Vorjahres", ""})
	tableRow(pdf, widths, []string{"Gesamte Einnahmen", fmt.Sprintf("%.2f EUR", revenue)})
	tableRow(pdf, widths, []string{"Gesamte Ausgaben", fmt.Sprintf("%.2f EUR", expenses)})
	pdf.Ln(8)

	pdf.Ln(8)
	pdf.MultiCell(0, 6, "(Bitte Formular vollst\xC3\xA4ndig ausf\xC3\xBCllen)", "", "L", false)

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return "", fmt.Errorf("%w: %v", ErrWritePDF, err)
	}
	return filePath, nil
}

// GenerateKSt1F creates a placeholder "KSt 1F" form for the given project.
func (g *Generator) GenerateKSt1F(ctx context.Context, projectID int64) (string, error) {
	p, _ := g.store.GetProject(ctx, projectID)
	info := g.formInfo()
	if info.Name == "" && p != nil {
		info.Name = p.Name
	}

	revenue, err := g.store.SumIncomeByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch revenue: %w", err)
	}
	expenses, err := g.store.SumExpenseByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch expenses: %w", err)
	}

	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	fileName := fmt.Sprintf("KSt_1F_%d.pdf", projectID)
	filePath := filepath.Join(g.BasePath, fileName)

	title := "KSt 1F - Erweiterte K\xC3\xB6rperschaftsteuererkl\xC3\xA4rung"
	if p != nil {
		title = fmt.Sprintf("KSt 1F - %s", p.Name)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	pdf.AddPage()
	addTemplate(pdf, "kst1f")
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 10, title, "", 1, "C", false, 0, "")
	pdf.Ln(4)

	widths := []float64{70, 110}
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "1. Angaben")
	pdf.Ln(8)
	tableHeader(pdf, widths, []string{"Angabe", "Wert"})
	tableRow(pdf, widths, []string{"Name des Vereins", info.Name})
	tableRow(pdf, widths, []string{"Steuernummer", info.TaxNumber})
	tableRow(pdf, widths, []string{"Angaben zu Beteiligungen", ""})
	tableRow(pdf, widths, []string{"Beteiligungen an Kapitalgesellschaften", ""})
	tableRow(pdf, widths, []string{"Erhaltene F\xC3\xB6rdermittel", ""})
	tableRow(pdf, widths, []string{"Gesamteinnahmen", fmt.Sprintf("%.2f EUR", revenue)})
	tableRow(pdf, widths, []string{"Gesamtausgaben", fmt.Sprintf("%.2f EUR", expenses)})
	pdf.Ln(8)

	pdf.Ln(8)
	pdf.MultiCell(0, 6, "(Bitte Formular vollst\xC3\xA4ndig ausf\xC3\xBCllen)", "", "L", false)

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return "", fmt.Errorf("%w: %v", ErrWritePDF, err)
	}
	return filePath, nil
}

// GenerateAnlageSport creates a placeholder "Anlage Sport" form for the given project.
func (g *Generator) GenerateAnlageSport(ctx context.Context, projectID int64) (string, error) {
	p, _ := g.store.GetProject(ctx, projectID)
	info := g.formInfo()
	if info.Name == "" && p != nil {
		info.Name = p.Name
	}
	if err := info.Validate(); err != nil {
		return "", err
	}

	revenue, err := g.store.SumIncomeByProject(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("fetch revenue: %w", err)
	}
	members, err := g.store.ListMembers(ctx)
	if err != nil {
		return "", fmt.Errorf("fetch members: %w", err)
	}
	memberCount := len(members)

	if err := os.MkdirAll(g.BasePath, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	fileName := fmt.Sprintf("Anlage_Sport_%d.pdf", projectID)
	filePath := filepath.Join(g.BasePath, fileName)

	title := "Anlage Sport - Sportvereine"
	if p != nil {
		title = fmt.Sprintf("Anlage Sport - %s", p.Name)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	pdf.AddPage()
	addTemplate(pdf, "anlage_sport")
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 10, title, "", 1, "C", false, 0, "")
	pdf.Ln(4)

	widths := []float64{70, 110}
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "1. Angaben")
	pdf.Ln(8)
	tableHeader(pdf, widths, []string{"Angabe", "Wert"})
	tableRow(pdf, widths, []string{"Name des Vereins", info.Name})
	tableRow(pdf, widths, []string{"Steuernummer", info.TaxNumber})
	tableRow(pdf, widths, []string{"Mitgliederzahl", fmt.Sprintf("%d", memberCount)})
	tableRow(pdf, widths, []string{"Einnahmen aus Sportbetrieb", fmt.Sprintf("%.2f EUR", revenue)})
	tableRow(pdf, widths, []string{"Anzahl der Übungsleiter", ""})
	pdf.Ln(8)

	pdf.Ln(8)
	pdf.MultiCell(0, 6, "(Bitte Formular vollst\xC3\xA4ndig ausf\xC3\xBCllen)", "", "L", false)

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return "", fmt.Errorf("%w: %v", ErrWritePDF, err)
	}
	return filePath, nil
}

// GenerateAllForms creates all available forms for the given project and returns their paths.
func (g *Generator) GenerateAllForms(ctx context.Context, projectID int64) ([]string, error) {
	var paths []string

	report, err := g.GenerateReport(ctx, projectID)
	if err != nil {
		return nil, err
	}
	paths = append(paths, report)

	forms := []func(context.Context, int64) (string, error){
		g.GenerateKSt1,
		g.GenerateAnlageGem,
		g.GenerateAnlageGK,
		g.GenerateKSt1F,
		g.GenerateAnlageSport,
	}
	for _, f := range forms {
		p, err := f(ctx, projectID)
		if err != nil {
			return nil, err
		}
		paths = append(paths, p)
	}
	return paths, nil
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
		return "", fmt.Errorf("%w: %v", ErrWritePDF, err)
	}
	return filePath, nil
}
