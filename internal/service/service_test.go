package service

import "testing"

func TestDataService_AddIncome(t *testing.T) {
	ds, err := NewDataService(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	proj, err := ds.CreateProject("Income Project")
	if err != nil {
		t.Fatal(err)
	}

	inc, err := ds.AddIncome(proj.ID, "donation", 15)
	if err != nil {
		t.Fatalf("AddIncome returned error: %v", err)
	}
	if inc.ID == 0 {
		t.Fatalf("expected income ID to be set")
	}

	list, err := ds.ListIncomes(proj.ID)
	if err != nil {
		t.Fatalf("ListIncomes returned error: %v", err)
	}
	if len(list) != 1 || list[0].Amount != 15 {
		t.Fatalf("unexpected incomes: %+v", list)
	}
}

func TestDataService_ListExpenses(t *testing.T) {
	ds, err := NewDataService(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	proj, err := ds.CreateProject("Expense Project")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := ds.AddExpense(proj.ID, "supplies", 5); err != nil {
		t.Fatal(err)
	}
	if _, err := ds.AddExpense(proj.ID, "food", 7); err != nil {
		t.Fatal(err)
	}

	expenses, err := ds.ListExpenses(proj.ID)
	if err != nil {
		t.Fatalf("ListExpenses returned error: %v", err)
	}
	if len(expenses) != 2 {
		t.Fatalf("expected 2 expenses, got %d", len(expenses))
	}
}

func TestDataService_CalculateProjectTaxes(t *testing.T) {
	ds, err := NewDataService(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	proj, err := ds.CreateProject("Tax Project")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := ds.AddIncome(proj.ID, "donation", 50000); err != nil {
		t.Fatal(err)
	}
	if _, err := ds.AddExpense(proj.ID, "rent", 2000); err != nil {
		t.Fatal(err)
	}

	result, err := ds.CalculateProjectTaxes(proj.ID)
	if err != nil {
		t.Fatalf("CalculateProjectTaxes returned error: %v", err)
	}
	if !result.IsTaxable {
		t.Fatalf("expected project to be taxable")
	}
	if result.TotalTax <= 0 {
		t.Fatalf("expected positive tax, got %f", result.TotalTax)
	}
}
