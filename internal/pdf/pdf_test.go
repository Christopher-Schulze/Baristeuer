package pdf

import (
	"os"
	"strings"
	"testing"

	"baristeuer/internal/data"
)

func assertContains(t *testing.T, path, substr string) {
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if !strings.Contains(string(b), substr) {
		t.Fatalf("%s does not contain %q", path, substr)
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
	assertContains(t, path, "Steuerbericht 2025")
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
		name string
		fn   func(int64) (string, error)
	}{
		{"kst1", g.GenerateKSt1},
		{"gem", g.GenerateAnlageGem},
		{"gk", g.GenerateAnlageGK},
		{"kst1f", g.GenerateKSt1F},
		{"sport", g.GenerateAnlageSport},
	}
	for _, f := range files {
		path, err := f.fn(proj.ID)
		if err != nil {
			t.Fatalf("%s failed: %v", f.name, err)
		}
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected file %s", path)
		}
		switch f.name {
		case "kst1":
			assertContains(t, path, "Finanzamt")
		case "gem":
			assertContains(t, path, "Gemeinn")
		case "gk":
			assertContains(t, path, "Gesch")
		case "kst1f":
			assertContains(t, path, "Feststellungszeitraum")
		case "sport":
			assertContains(t, path, "Sportliche")
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
	assertContains(t, paths[0], "Steuerbericht")
}
