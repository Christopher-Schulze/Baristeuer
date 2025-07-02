package data

import "testing"

func TestProjectCRUD(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	p := &Project{Name: "Test"}
	if err := s.CreateProject(p); err != nil {
		t.Fatal(err)
	}

	got, err := s.GetProject(p.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != p.Name {
		t.Fatalf("expected %s, got %s", p.Name, got.Name)
	}

	p.Name = "Updated"
	if err := s.UpdateProject(p); err != nil {
		t.Fatal(err)
	}

	got, err = s.GetProject(p.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != "Updated" {
		t.Fatalf("update failed: got %s", got.Name)
	}

	if err := s.DeleteProject(p.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.GetProject(p.ID); err == nil {
		t.Fatalf("expected error after delete")
	}
}

func TestIncomeCRUD(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	proj := &Project{Name: "Income Project"}
	if err := s.CreateProject(proj); err != nil {
		t.Fatal(err)
	}

	i := &Income{ProjectID: proj.ID, Source: "donation", Amount: 10}
	if err := s.CreateIncome(i); err != nil {
		t.Fatal(err)
	}

	got, err := s.GetIncome(i.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Amount != i.Amount {
		t.Fatalf("expected %f, got %f", i.Amount, got.Amount)
	}

	i.Amount = 20
	if err := s.UpdateIncome(i); err != nil {
		t.Fatal(err)
	}
	got, _ = s.GetIncome(i.ID)
	if got.Amount != 20 {
		t.Fatalf("update failed")
	}

	if err := s.DeleteIncome(i.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.GetIncome(i.ID); err == nil {
		t.Fatalf("expected error after delete")
	}
}

func TestExpenseCRUD(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	proj := &Project{Name: "Expense Project"}
	if err := s.CreateProject(proj); err != nil {
		t.Fatal(err)
	}

	e := &Expense{ProjectID: proj.ID, Category: "supplies", Amount: 5}
	if err := s.CreateExpense(e); err != nil {
		t.Fatal(err)
	}

	got, err := s.GetExpense(e.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Amount != e.Amount {
		t.Fatalf("expected %f, got %f", e.Amount, got.Amount)
	}

	e.Amount = 8
	if err := s.UpdateExpense(e); err != nil {
		t.Fatal(err)
	}
	got, _ = s.GetExpense(e.ID)
	if got.Amount != 8 {
		t.Fatalf("update failed")
	}

	if err := s.DeleteExpense(e.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.GetExpense(e.ID); err == nil {
		t.Fatalf("expected error after delete")
	}
}
