package data

import (
	"context"
	"errors"
	"testing"
)

func TestProjectCRUD(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	ctx := context.Background()
	p := &Project{Name: "Test"}
	if err := s.CreateProject(ctx, p); err != nil {
		t.Fatal(err)
	}

	got, err := s.GetProject(ctx, p.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != p.Name {
		t.Fatalf("expected %s, got %s", p.Name, got.Name)
	}

	p.Name = "Updated"
	if err := s.UpdateProject(ctx, p); err != nil {
		t.Fatal(err)
	}

	got, err = s.GetProject(ctx, p.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != "Updated" {
		t.Fatalf("update failed: got %s", got.Name)
	}

	if err := s.DeleteProject(ctx, p.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.GetProject(ctx, p.ID); err == nil {
		t.Fatalf("expected error after delete")
	}
}

func TestIncomeCRUD(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	ctx := context.Background()
	proj := &Project{Name: "Income Project"}
	if err := s.CreateProject(ctx, proj); err != nil {
		t.Fatal(err)
	}

	i := &Income{ProjectID: proj.ID, Source: "donation", Amount: 10}
	if err := s.CreateIncome(ctx, i); err != nil {
		t.Fatal(err)
	}

	got, err := s.GetIncome(ctx, i.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Amount != i.Amount {
		t.Fatalf("expected %f, got %f", i.Amount, got.Amount)
	}

	i.Amount = 20
	if err := s.UpdateIncome(ctx, i); err != nil {
		t.Fatal(err)
	}
	got, _ = s.GetIncome(ctx, i.ID)
	if got.Amount != 20 {
		t.Fatalf("update failed")
	}

	if err := s.DeleteIncome(ctx, i.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.GetIncome(ctx, i.ID); err == nil {
		t.Fatalf("expected error after delete")
	}
}

func TestExpenseCRUD(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	ctx := context.Background()
	proj := &Project{Name: "Expense Project"}
	if err := s.CreateProject(ctx, proj); err != nil {
		t.Fatal(err)
	}

	e := &Expense{ProjectID: proj.ID, Category: "supplies", Amount: 5}
	if err := s.CreateExpense(ctx, e); err != nil {
		t.Fatal(err)
	}

	got, err := s.GetExpense(ctx, e.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Amount != e.Amount {
		t.Fatalf("expected %f, got %f", e.Amount, got.Amount)
	}

	e.Amount = 8
	if err := s.UpdateExpense(ctx, e); err != nil {
		t.Fatal(err)
	}
	got, _ = s.GetExpense(ctx, e.ID)
	if got.Amount != 8 {
		t.Fatalf("update failed")
	}

	if err := s.DeleteExpense(ctx, e.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.GetExpense(ctx, e.ID); err == nil {
		t.Fatalf("expected error after delete")
	}
}

func TestMemberCRUD(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	ctx := context.Background()
	m := &Member{Name: "Alice", Email: "alice@example.com", JoinDate: "2024-01-02"}
	if err := s.CreateMember(ctx, m); err != nil {
		t.Fatal(err)
	}

	got, err := s.GetMember(ctx, m.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Email != m.Email {
		t.Fatalf("expected %s, got %s", m.Email, got.Email)
	}

	m.Name = "Alice Smith"
	if err := s.UpdateMember(ctx, m); err != nil {
		t.Fatal(err)
	}
	got, _ = s.GetMember(ctx, m.ID)
	if got.Name != "Alice Smith" {
		t.Fatalf("update failed")
	}

	members, err := s.ListMembers(ctx)
	if err != nil || len(members) != 1 {
		t.Fatalf("expected one member, got %v", members)
	}

	if err := s.DeleteMember(ctx, m.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.GetMember(ctx, m.ID); err == nil {
		t.Fatalf("expected error after delete")
	}
}

func TestStore_ContextCancellation(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = s.CreateProject(ctx, &Project{Name: "ctx"})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context canceled, got %v", err)
	}
}
