package pdf

import (
	"os"
	"testing"

	"baristeuer/internal/data"
)

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
