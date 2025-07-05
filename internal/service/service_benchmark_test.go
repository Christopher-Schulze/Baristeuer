package service

import (
	"context"
	"io"
	"log/slog"
	"testing"
)

func benchmarkAddIncome(b *testing.B, n int) {
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		b.Fatal(err)
	}
	defer ds.Close()
	ctx := context.Background()
	proj, err := ds.CreateProject(ctx, "bench")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			if _, err := ds.AddIncome(ctx, proj.ID, "donation", 1); err != nil {
				b.Fatalf("add income: %v", err)
			}
		}
	}
}

func BenchmarkAddIncomeMass1e2(b *testing.B) { benchmarkAddIncome(b, 100) }
func BenchmarkAddIncomeMass1e3(b *testing.B) { benchmarkAddIncome(b, 1000) }

func benchmarkAddExpense(b *testing.B, n int) {
	ds, err := NewDataService(":memory:", slog.New(slog.NewTextHandler(io.Discard, nil)), nil, nil)
	if err != nil {
		b.Fatal(err)
	}
	defer ds.Close()
	ctx := context.Background()
	proj, err := ds.CreateProject(ctx, "bench")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			if _, err := ds.AddExpense(ctx, proj.ID, "rent", 1); err != nil {
				b.Fatalf("add expense: %v", err)
			}
		}
	}
}

func BenchmarkAddExpenseMass1e2(b *testing.B) { benchmarkAddExpense(b, 100) }
func BenchmarkAddExpenseMass1e3(b *testing.B) { benchmarkAddExpense(b, 1000) }
