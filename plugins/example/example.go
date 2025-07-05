package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"baristeuer/internal/data"
	"baristeuer/internal/plugins"
	"baristeuer/internal/service"
)

type exporter struct{}

func New() plugins.Plugin { return &exporter{} }

func (e *exporter) Init(ds *service.DataService) error {
	ctx := context.Background()
	projects, err := ds.ListProjects()
	if err != nil {
		return fmt.Errorf("list projects: %w", err)
	}
	type item struct {
		Project  data.Project   `json:"project"`
		Incomes  []data.Income  `json:"incomes"`
		Expenses []data.Expense `json:"expenses"`
	}
	var all []item
	for _, p := range projects {
		inc, err := ds.ListIncomes(ctx, p.ID)
		if err != nil {
			return fmt.Errorf("list incomes: %w", err)
		}
		exp, err := ds.ListExpenses(ctx, p.ID)
		if err != nil {
			return fmt.Errorf("list expenses: %w", err)
		}
		all = append(all, item{Project: p, Incomes: inc, Expenses: exp})
	}
	out := os.Getenv("EXAMPLE_EXPORT_FILE")
	if out == "" {
		out = "example_export.json"
	}
	f, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("create export: %w", err)
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(all); err != nil {
		return fmt.Errorf("encode export: %w", err)
	}
	return nil
}
