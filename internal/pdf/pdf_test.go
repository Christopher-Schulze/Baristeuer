package pdf

import (
	"context"
	"os"
	"strings"
	"testing"

	"baristeuer/internal/data"
)

func TestNewGeneratorEnvVar(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("BARISTEUER_PDFDIR", dir)
	g := NewGenerator("", nil)
	if g.BasePath != dir {
		t.Fatalf("expected %s, got %s", dir, g.BasePath)
	}
}

func TestGenerateReport(t *testing.T) {
	dir := t.TempDir()
	store, err := data.NewStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	ctx := context.Background()

	proj := &data.Project{Name: "Test"}
	if err := store.CreateProject(ctx, proj); err != nil {
		t.Fatal(err)
	}
	if err := store.CreateIncome(ctx, &data.Income{ProjectID: proj.ID, Source: "donation", Amount: 100}); err != nil {
		t.Fatal(err)
	}
	if err := store.CreateExpense(ctx, &data.Expense{ProjectID: proj.ID, Category: "supplies", Amount: 20}); err != nil {
		t.Fatal(err)
	}

	g := NewGenerator(dir, store)
	path, err := g.GenerateReport(proj.ID)
	if err != nil {
		t.Fatalf("GenerateReport failed: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected pdf at %s", path)
	}
}

func TestFormGeneration(t *testing.T) {
	dir := t.TempDir()
	store, err := data.NewStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	ctx := context.Background()

	proj := &data.Project{Name: "Test"}
	if err := store.CreateProject(ctx, proj); err != nil {
		t.Fatal(err)
	}
	if err := store.CreateIncome(ctx, &data.Income{ProjectID: proj.ID, Source: "donation", Amount: 100}); err != nil {
		t.Fatal(err)
	}
	if err := store.CreateExpense(ctx, &data.Expense{ProjectID: proj.ID, Category: "supplies", Amount: 20}); err != nil {
		t.Fatal(err)
	}
	if err := store.CreateMember(ctx, &data.Member{Name: "Max", Email: "max@example.com", JoinDate: "2024-01-01"}); err != nil {
		t.Fatal(err)
	}

	g := NewGenerator(dir, store)
	info := FormInfo{Name: "Testverein", TaxNumber: "11/111/11111", Address: "Hauptstr. 1", FiscalYear: "2025"}
	files := []struct {
		name     string
		fn       func(int64, FormInfo) (string, error)
		expected []string
	}{
		{"kst1", g.GenerateKSt1, []string{"Einnahmen gesamt", "100.00", "Ausgaben gesamt"}},
		{"gem", g.GenerateAnlageGem, []string{"Mitglieder:", "1", "Einnahmen:", "100.00"}},
		{"gk", g.GenerateAnlageGK, []string{"Gesamte Einnahmen", "100.00"}},
		{"kst1f", g.GenerateKSt1F, []string{"Gesamteinnahmen", "100.00"}},
		{"sport", g.GenerateAnlageSport, []string{"Mitgliederzahl", "1", "Einnahmen aus Sportbetrieb"}},
	}
	for _, f := range files {
		path, err := f.fn(proj.ID, info)
		if err != nil {
			t.Fatalf("%s failed: %v", f.name, err)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s failed: %v", path, err)
		}
		for _, expect := range f.expected {
			if !strings.Contains(string(data), expect) {
				t.Fatalf("%s form missing %s in %s", f.name, expect, string(data))
			}
		}
	}

	// test GenerateAllForms
	paths, err := g.GenerateAllForms(proj.ID, info)
	if err != nil {
		t.Fatalf("GenerateAllForms failed: %v", err)
	}
	if len(paths) != len(files)+1 { // +1 for report
		t.Fatalf("expected %d files, got %d", len(files)+1, len(paths))
	}
}
