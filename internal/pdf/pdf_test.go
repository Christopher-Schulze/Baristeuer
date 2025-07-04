package pdf

import (
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

	proj := &data.Project{Name: "Test"}
	if err := store.CreateProject(proj); err != nil {
		t.Fatal(err)
	}
	if err := store.CreateIncome(&data.Income{ProjectID: proj.ID, Source: "donation", Amount: 100}); err != nil {
		t.Fatal(err)
	}
	if err := store.CreateExpense(&data.Expense{ProjectID: proj.ID, Category: "supplies", Amount: 20}); err != nil {
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

	proj := &data.Project{Name: "Test"}
	if err := store.CreateProject(proj); err != nil {
		t.Fatal(err)
	}

	g := NewGenerator(dir, store)
	files := []struct {
		name     string
		fn       func(int64) (string, error)
		expected []string
	}{
		{"kst1", g.GenerateKSt1, []string{"KSt 1 - K\xC3\xB6rperschaftsteuererkl\xC3\xA4rung", "Rechtsform", "Satzungsdatum"}},
		{"gem", g.GenerateAnlageGem, []string{"Anlage Gem", "Anschrift des Vereins", "Bankverbindung"}},
		{"gk", g.GenerateAnlageGK, []string{"Anlage GK", "Umsatz des Vorjahres"}},
		{"kst1f", g.GenerateKSt1F, []string{"KSt 1F", "Kapitalgesellschaften"}},
		{"sport", g.GenerateAnlageSport, []string{"Anlage Sport", "\xC3\x9Cbungsleiter"}},
	}
	for _, f := range files {
		path, err := f.fn(proj.ID)
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
	paths, err := g.GenerateAllForms(proj.ID)
	if err != nil {
		t.Fatalf("GenerateAllForms failed: %v", err)
	}
	if len(paths) != len(files)+1 { // +1 for report
		t.Fatalf("expected %d files, got %d", len(files)+1, len(paths))
	}
}
