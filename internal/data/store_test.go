package data

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func newTestStore(t *testing.T) *Store {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	s := NewStore(db)
	if err := s.Init(context.Background()); err != nil {
		t.Fatal(err)
	}
	return s
}

func TestProjectCRUD(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()

	p := &Project{Name: "Proj", Description: "desc"}
	if err := s.CreateProject(ctx, p); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := s.GetProject(ctx, p.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != p.Name {
		t.Fatalf("want name %q got %q", p.Name, got.Name)
	}

	p.Name = "NewName"
	if err := s.UpdateProject(ctx, p); err != nil {
		t.Fatalf("update: %v", err)
	}
	got, _ = s.GetProject(ctx, p.ID)
	if got.Name != "NewName" {
		t.Fatalf("update failed")
	}

	if err := s.DeleteProject(ctx, p.ID); err != nil {
		t.Fatalf("delete: %v", err)
	}
	if _, err := s.GetProject(ctx, p.ID); err == nil {
		t.Fatalf("expected error after delete")
	}
}

func TestIncomeExpenseCRUD(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()

	p := &Project{Name: "Proj"}
	if err := s.CreateProject(ctx, p); err != nil {
		t.Fatal(err)
	}

	in := &Income{ProjectID: p.ID, Amount: 10, Description: "donation", Date: time.Now()}
	if err := s.CreateIncome(ctx, in); err != nil {
		t.Fatalf("create income: %v", err)
	}

	ex := &Expense{ProjectID: p.ID, Amount: 5, Description: "cost", Date: time.Now()}
	if err := s.CreateExpense(ctx, ex); err != nil {
		t.Fatalf("create expense: %v", err)
	}

	ins, err := s.ListIncomesByProject(ctx, p.ID)
	if err != nil || len(ins) != 1 {
		t.Fatalf("list incomes: %v len=%d", err, len(ins))
	}

	exs, err := s.ListExpensesByProject(ctx, p.ID)
	if err != nil || len(exs) != 1 {
		t.Fatalf("list expenses: %v len=%d", err, len(exs))
	}
}
