package service

import (
	"baristeuer/internal/config"
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDataService_AddIncome(t *testing.T) {
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
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
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
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
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
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
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
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
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
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

	for _, year := range []int{2025, 2026} {
		result, err := ds.CalculateProjectTaxes(ctx, proj.ID, year)
		if err != nil {
			t.Fatalf("CalculateProjectTaxes returned error: %v", err)
		}
		if result.Year != year {
			t.Fatalf("expected year %d got %d", year, result.Year)
		}
		if !result.IsTaxable {
			t.Fatalf("expected project to be taxable")
		}
		if result.TotalTax <= 0 {
			t.Fatalf("expected positive tax, got %f", result.TotalTax)
		}
	}
}

func TestDataService_GenerateStatistics(t *testing.T) {
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()
	proj, _ := ds.CreateProject(ctx, "Stats")
	ds.AddIncome(ctx, proj.ID, "donation", 10)
	ds.AddIncome(ctx, proj.ID, "donation", 20)
	ds.AddExpense(ctx, proj.ID, "rent", 5)
	ds.AddExpense(ctx, proj.ID, "rent", 15)

	stats, err := ds.GenerateStatistics(ctx, proj.ID, 2025)
	if err != nil {
		t.Fatalf("GenerateStatistics returned error: %v", err)
	}
	if stats.AverageIncome != 15 || stats.AverageExpense != 10 ||
		stats.MedianIncome != 15 || stats.MedianExpense != 10 ||
		stats.StdDevIncome != 5 || stats.StdDevExpense != 5 ||
		stats.Trend != 5 || stats.Year != 2025 {
		t.Fatalf("unexpected stats: %+v", stats)
	}
}

func TestDataService_MemberOperations(t *testing.T) {
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
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

func TestDataService_UpdateDeleteMember(t *testing.T) {
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()

	m, err := ds.AddMember(ctx, "Bob", "bob@example.com", "2024-01-10")
	if err != nil {
		t.Fatal(err)
	}

	if err := ds.UpdateMember(ctx, m.ID, "Bobby", "bob@example.com", "2024-02-01"); err != nil {
		t.Fatalf("UpdateMember failed: %v", err)
	}

	members, _ := ds.ListMembers(ctx)
	if len(members) != 1 || members[0].Name != "Bobby" || members[0].JoinDate != "2024-02-01" {
		t.Fatalf("update failed: %+v", members)
	}

	if err := ds.DeleteMember(ctx, m.ID); err != nil {
		t.Fatalf("DeleteMember failed: %v", err)
	}

	members, _ = ds.ListMembers(ctx)
	if len(members) != 0 {
		t.Fatalf("expected empty list, got %+v", members)
	}
}

func TestDataService_AddIncome_LogOutput(t *testing.T) {
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()

	var buf bytes.Buffer
	ds.logger = slog.New(slog.NewTextHandler(&buf, nil))

	proj, _ := ds.CreateProject(ctx, "Log Project")
	if _, err := ds.AddIncome(ctx, proj.ID, "donation", 5); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "added income") {
		t.Fatalf("log output missing: %s", buf.String())
	}
}

func TestDataService_AddIncome_DatabaseClosed(t *testing.T) {
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	ds.Close()

	_, err = ds.AddIncome(context.Background(), 1, "donation", 5)
	if err == nil || !strings.Contains(err.Error(), "create income") {
		t.Fatalf("expected wrapped error, got %v", err)
	}
}

func TestDataService_ProjectOperations(t *testing.T) {
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()
	p, err := ds.CreateProject(ctx, "Proj1")
	if err != nil {
		t.Fatal(err)
	}
	list, err := ds.ListProjects()
	if err != nil || len(list) == 0 {
		t.Fatalf("ListProjects failed: %v", err)
	}
	got, err := ds.GetProject(ctx, p.ID)
	if err != nil || got.ID != p.ID {
		t.Fatalf("GetProject failed: %v", err)
	}
	if err := ds.UpdateProject(ctx, p.ID, "New"); err != nil {
		t.Fatalf("UpdateProject failed: %v", err)
	}
	updated, _ := ds.GetProject(ctx, p.ID)
	if updated.Name != "New" {
		t.Fatalf("update did not persist")
	}
	if err := ds.DeleteProject(ctx, p.ID); err != nil {
		t.Fatalf("DeleteProject failed: %v", err)
	}
	list, _ = ds.ListProjects()
	if len(list) != 0 {
		t.Fatalf("expected empty list")
	}
}

func TestDataService_InvalidAmounts(t *testing.T) {
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()

	proj, _ := ds.CreateProject(ctx, "Invalid Amount")

	if _, err := ds.AddIncome(ctx, proj.ID, "donation", -5); err == nil || !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected invalid amount error, got %v", err)
	}
	if _, err := ds.AddExpense(ctx, proj.ID, "supplies", 0); err == nil || !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected invalid amount error, got %v", err)
	}

	inc, _ := ds.AddIncome(ctx, proj.ID, "donation", 5)
	if err := ds.UpdateIncome(ctx, inc.ID, proj.ID, "donation", 0); err == nil || !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected invalid amount error, got %v", err)
	}
	exp, _ := ds.AddExpense(ctx, proj.ID, "supplies", 5)
	if err := ds.UpdateExpense(ctx, exp.ID, proj.ID, "supplies", -2); err == nil || !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected invalid amount error, got %v", err)
	}
}

func TestValidateAmount(t *testing.T) {
	if err := validateAmount(10); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	err := validateAmount(0)
	if err == nil || !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected invalid amount error, got %v", err)
	}
}

func TestDataService_LoggerClosed(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "app.log")
	logger, closer := NewLogger(logPath, "info", "text")
	if closer == nil {
		t.Fatalf("expected closer for log file")
	}
	ds, err := NewDataService(":memory:", logger, closer, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.Close(); err != nil {
		t.Fatalf("close service: %v", err)
	}
	if err := closer.Close(); err != nil {
		t.Fatalf("unexpected error closing logger: %v", err)
	}
}

func TestDataService_ExportDatabase(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test.db")
	ds, err := NewDataService(dbPath, slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	dest := filepath.Join(tmpDir, "export.db")
	if err := ds.ExportDatabase(dest); err != nil {
		t.Fatalf("ExportDatabase returned error: %v", err)
	}
	info, err := os.Stat(dest)
	if err != nil {
		t.Fatalf("exported file not found: %v", err)
	}
	if info.Size() == 0 {
		t.Fatal("expected exported file to be non-empty")
	}
}

func TestDataService_RestoreDatabase(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "orig.db")
	ds, err := NewDataService(dbPath, slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()
	proj, _ := ds.CreateProject(ctx, "One")
	_ = proj

	backup := filepath.Join(tmpDir, "backup.db")
	if err := ds.ExportDatabase(backup); err != nil {
		t.Fatalf("export: %v", err)
	}

	// add another project after backup
	if _, err := ds.CreateProject(ctx, "Two"); err != nil {
		t.Fatal(err)
	}

	if err := ds.RestoreDatabase(backup); err != nil {
		t.Fatalf("restore: %v", err)
	}

	list, _ := ds.ListProjects()
	if len(list) != 1 || list[0].Name != "One" {
		t.Fatalf("unexpected projects after restore: %+v", list)
	}
}

func TestDataService_SetLogLevel(t *testing.T) {
	logger, closer := NewLogger("", "info", "text")
	ds, err := NewDataService(":memory:", logger, closer, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	SetLogLevel("debug")
	if logLevelVar.Level() != slog.LevelDebug {
		t.Fatalf("expected debug level, got %v", logLevelVar.Level())
	}

	ds.SetLogLevel("error")
	if logLevelVar.Level() != slog.LevelError {
		t.Fatalf("expected error level, got %v", logLevelVar.Level())
	}
}

func TestDataService_SetLogFormat(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "app.log")
	logger, closer := NewLogger(logPath, "info", "text")
	ds, err := NewDataService(":memory:", logger, closer, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ds.SetLogFormat("json")
	if closer != nil {
		closer.Close()
	}

	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(data, []byte("\"format\"")) && !bytes.Contains(data, []byte("\"level\"")) {
		t.Fatalf("expected json log, got %s", data)
	}
}

func TestDataService_ExportProjectCSV(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "db.db")
	ds, err := NewDataService(dbPath, slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()
	proj, _ := ds.CreateProject(ctx, "CSV")
	if _, err := ds.AddIncome(ctx, proj.ID, "donation", 10); err != nil {
		t.Fatal(err)
	}
	if _, err := ds.AddExpense(ctx, proj.ID, "rent", 5); err != nil {
		t.Fatal(err)
	}

	dest := filepath.Join(tmpDir, "out.csv")
	if err := ds.ExportProjectCSV(ctx, proj.ID, dest); err != nil {
		t.Fatalf("ExportProjectCSV returned error: %v", err)
	}
	data, err := os.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "income") || !strings.Contains(string(data), "expense") {
		t.Fatalf("csv content unexpected: %s", data)
	}
}

func TestFormFieldGettersSetters(t *testing.T) {
	ds, err := NewDataService(":memory:", nil, nil, &config.Config{})
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()
	ds.SetFormCity("Town")
	if ds.GetFormCity() != "Town" {
		t.Fatalf("city getter/setter failed")
	}
	ds.SetFormBankAccount("DE123")
	if ds.GetFormBankAccount() != "DE123" {
		t.Fatalf("bank account getter/setter failed")
	}
	ds.SetFormRepresentative("Alice")
	if ds.GetFormRepresentative() != "Alice" {
		t.Fatalf("representative getter/setter failed")
	}
}

func TestDataService_ContextCancellation(t *testing.T) {
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()
	proj, _ := ds.CreateProject(ctx, "Cancel")

	cctx, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := ds.AddIncome(cctx, proj.ID, "donation", 1); err == nil || !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context canceled error, got %v", err)
	}

	cctx2, cancel2 := context.WithCancel(ctx)
	cancel2()
	if _, err := ds.GetProject(cctx2, proj.ID); err == nil || !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context canceled error, got %v", err)
	}
}

func TestDataService_ExportRestoreCycle(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "db.db")
	ds, err := NewDataService(dbPath, slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.Close()

	ctx := context.Background()

	proj, err := ds.CreateProject(ctx, "Cycle")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := ds.AddIncome(ctx, proj.ID, "donation", 12); err != nil {
		t.Fatal(err)
	}
	if _, err := ds.AddExpense(ctx, proj.ID, "rent", 7); err != nil {
		t.Fatal(err)
	}
	if _, err := ds.AddMember(ctx, "Alice", "alice@example.com", "2024-01-02"); err != nil {
		t.Fatal(err)
	}

	exportPath := filepath.Join(tmpDir, "backup.db")
	if err := ds.ExportDatabase(exportPath); err != nil {
		t.Fatalf("export: %v", err)
	}

	if err := os.Remove(dbPath); err != nil {
		t.Fatalf("remove db: %v", err)
	}

	if err := ds.RestoreDatabase(exportPath); err != nil {
		t.Fatalf("restore: %v", err)
	}

	projects, err := ds.ListProjects()
	if err != nil {
		t.Fatal(err)
	}
	if len(projects) != 1 || projects[0].Name != "Cycle" {
		t.Fatalf("unexpected projects: %+v", projects)
	}
	incomes, err := ds.ListIncomes(ctx, proj.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(incomes) != 1 || incomes[0].Amount != 12 {
		t.Fatalf("unexpected incomes: %+v", incomes)
	}
	expenses, err := ds.ListExpenses(ctx, proj.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(expenses) != 1 || expenses[0].Amount != 7 {
		t.Fatalf("unexpected expenses: %+v", expenses)
	}
	members, err := ds.ListMembers(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(members) != 1 || members[0].Email != "alice@example.com" {
		t.Fatalf("unexpected members: %+v", members)
	}
}
