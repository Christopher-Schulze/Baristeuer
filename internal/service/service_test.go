package service

import "testing"

func TestDataService_ProjectCRUD(t *testing.T) {
	ds, err := NewDataService(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	p, err := ds.CreateProject("Proj1")
	if err != nil {
		t.Fatal(err)
	}
	if p.ID == 0 {
		t.Fatalf("expected ID set")
	}

	list, err := ds.ListProjects()
	if err != nil || len(list) != 1 {
		t.Fatalf("list projects: %v", list)
	}

	if err := ds.UpdateProject(p.ID, "Updated"); err != nil {
		t.Fatal(err)
	}
	proj, _ := ds.store.GetProject(p.ID)
	if proj.Name != "Updated" {
		t.Fatalf("update failed")
	}

	if err := ds.DeleteProject(p.ID); err != nil {
		t.Fatal(err)
	}
	list, _ = ds.ListProjects()
	if len(list) != 0 {
		t.Fatalf("expected empty list, got %v", list)
	}
}
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

func TestDataService_UpdateDeleteIncome(t *testing.T) {
	ds, err := NewDataService(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	proj, err := ds.CreateProject("Income UD")
	if err != nil {
		t.Fatal(err)
	}

	inc, err := ds.AddIncome(proj.ID, "donation", 10)
	if err != nil {
		t.Fatal(err)
	}

	if err := ds.UpdateIncome(inc.ID, proj.ID, "donation", 20); err != nil {
		t.Fatalf("UpdateIncome failed: %v", err)
	}

	list, _ := ds.ListIncomes(proj.ID)
	if len(list) != 1 || list[0].Amount != 20 {
		t.Fatalf("update failed: %+v", list)
	}

	if err := ds.DeleteIncome(inc.ID); err != nil {
		t.Fatalf("DeleteIncome failed: %v", err)
	}
	list, _ = ds.ListIncomes(proj.ID)
	if len(list) != 0 {
		t.Fatalf("expected empty list, got %+v", list)
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

func TestDataService_UpdateDeleteExpense(t *testing.T) {
	ds, err := NewDataService(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	proj, err := ds.CreateProject("Expense UD")
	if err != nil {
		t.Fatal(err)
	}

	exp, err := ds.AddExpense(proj.ID, "supplies", 5)
	if err != nil {
		t.Fatal(err)
	}

	if err := ds.UpdateExpense(exp.ID, proj.ID, "supplies", 8); err != nil {
		t.Fatalf("UpdateExpense failed: %v", err)
	}

	list, _ := ds.ListExpenses(proj.ID)
	if len(list) != 1 || list[0].Amount != 8 {
		t.Fatalf("update failed: %+v", list)
	}

	if err := ds.DeleteExpense(exp.ID); err != nil {
		t.Fatalf("DeleteExpense failed: %v", err)
	}
	list, _ = ds.ListExpenses(proj.ID)
	if len(list) != 0 {
		t.Fatalf("expected empty list, got %+v", list)
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

func TestDataService_MemberOperations(t *testing.T) {
	ds, err := NewDataService(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	m, err := ds.AddMember("Bob", "bob@example.com", "2024-01-10")
	if err != nil {
		t.Fatalf("AddMember returned error: %v", err)
	}
	if m.ID == 0 {
		t.Fatalf("expected member ID to be set")
	}

	members, err := ds.ListMembers()
	if err != nil {
		t.Fatalf("ListMembers returned error: %v", err)
	}
	if len(members) != 1 || members[0].Email != "bob@example.com" {
		t.Fatalf("unexpected members: %+v", members)
	}
}
