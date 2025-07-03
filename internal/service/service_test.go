package service

import (
	"context"
	"errors"
	"testing"
)

func TestDataService_AddIncome(t *testing.T) {
	ds, err := NewDataService(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()
	proj, err := ds.CreateProject(ctx, "Income Project")
	if err != nil {
		t.Fatal(err)
	}

	inc, err := ds.AddIncome(ctx, proj.ID, "donation", 15)
	if err != nil {
		t.Fatalf("AddIncome returned error: %v", err)
	}
	if inc.ID == 0 {
		t.Fatalf("expected income ID to be set")
	}

	list, err := ds.ListIncomes(ctx, proj.ID)
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

	ctx := context.Background()
	proj, err := ds.CreateProject(ctx, "Income UD")
	if err != nil {
		t.Fatal(err)
	}

	inc, err := ds.AddIncome(ctx, proj.ID, "donation", 10)
	if err != nil {
		t.Fatal(err)
	}

	if err := ds.UpdateIncome(ctx, inc.ID, proj.ID, "donation", 20); err != nil {
		t.Fatalf("UpdateIncome failed: %v", err)
	}

	list, _ := ds.ListIncomes(ctx, proj.ID)
	if len(list) != 1 || list[0].Amount != 20 {
		t.Fatalf("update failed: %+v", list)
	}

	if err := ds.DeleteIncome(ctx, inc.ID); err != nil {
		t.Fatalf("DeleteIncome failed: %v", err)
	}
	list, _ = ds.ListIncomes(ctx, proj.ID)
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

	ctx := context.Background()
	proj, err := ds.CreateProject(ctx, "Expense Project")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := ds.AddExpense(ctx, proj.ID, "supplies", 5); err != nil {
		t.Fatal(err)
	}
	if _, err := ds.AddExpense(ctx, proj.ID, "food", 7); err != nil {
		t.Fatal(err)
	}

	expenses, err := ds.ListExpenses(ctx, proj.ID)
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

	ctx := context.Background()
	proj, err := ds.CreateProject(ctx, "Expense UD")
	if err != nil {
		t.Fatal(err)
	}

	exp, err := ds.AddExpense(ctx, proj.ID, "supplies", 5)
	if err != nil {
		t.Fatal(err)
	}

	if err := ds.UpdateExpense(ctx, exp.ID, proj.ID, "supplies", 8); err != nil {
		t.Fatalf("UpdateExpense failed: %v", err)
	}

	list, _ := ds.ListExpenses(ctx, proj.ID)
	if len(list) != 1 || list[0].Amount != 8 {
		t.Fatalf("update failed: %+v", list)
	}

	if err := ds.DeleteExpense(ctx, exp.ID); err != nil {
		t.Fatalf("DeleteExpense failed: %v", err)
	}
	list, _ = ds.ListExpenses(ctx, proj.ID)
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

	ctx := context.Background()
	proj, err := ds.CreateProject(ctx, "Tax Project")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := ds.AddIncome(ctx, proj.ID, "donation", 50000); err != nil {
		t.Fatal(err)
	}
	if _, err := ds.AddExpense(ctx, proj.ID, "rent", 2000); err != nil {
		t.Fatal(err)
	}

	result, err := ds.CalculateProjectTaxes(ctx, proj.ID)
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

	ctx := context.Background()
	m, err := ds.AddMember(ctx, "Bob", "bob@example.com", "2024-01-10")
	if err != nil {
		t.Fatalf("AddMember returned error: %v", err)
	}
	if m.ID == 0 {
		t.Fatalf("expected member ID to be set")
	}

	members, err := ds.ListMembers(ctx)
	if err != nil {
		t.Fatalf("ListMembers returned error: %v", err)
	}
	if len(members) != 1 || members[0].Email != "bob@example.com" {
		t.Fatalf("unexpected members: %+v", members)
	}
}

func TestDataService_ContextCancellation(t *testing.T) {
	ds, err := NewDataService(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()
	proj, err := ds.CreateProject(ctx, "ctx")
	if err != nil {
		t.Fatal(err)
	}

	cctx, cancel := context.WithCancel(ctx)
	cancel()
	_, err = ds.AddIncome(cctx, proj.ID, "donation", 1)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context canceled, got %v", err)
	}
}
