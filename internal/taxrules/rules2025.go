package taxrules

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Rules struct {
	Allowance float64 `json:"allowance"`
	Rate      float64 `json:"rate"`
	Threshold float64 `json:"threshold"`
}

var rules2025 Rules

func init() {
	if err := loadRules2025(); err != nil {
		panic(err)
	}
}

func loadRules2025() error {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	path := filepath.Join(dir, "rules2025.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config: %w", err)
	}
	if err := json.Unmarshal(data, &rules2025); err != nil {
		return fmt.Errorf("parse config: %w", err)
	}
	return nil
}

// Project represents a simple club project with revenue and expenses.
type Project struct {
	Revenue  float64
	Expenses float64
}

// CalculateTax computes the tax due for a project based on the 2025 rules.
func CalculateTax(p Project) float64 {
	taxable := p.Revenue - p.Expenses - rules2025.Allowance
	if taxable <= rules2025.Threshold {
		return 0
	}
	return (taxable - rules2025.Threshold) * rules2025.Rate
}
