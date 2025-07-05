package report

import "testing"

func TestAverage(t *testing.T) {
	if avg := Average([]float64{1, 2, 3}); avg != 2 {
		t.Fatalf("expected 2 got %f", avg)
	}
	if Average(nil) != 0 {
		t.Fatalf("expected 0 for empty slice")
	}
}

func TestCalculate(t *testing.T) {
	stats := Calculate([]float64{10, 20}, []float64{5, 15}, 2025)
	if stats.AverageIncome != 15 {
		t.Fatalf("avg income expected 15 got %f", stats.AverageIncome)
	}
	if stats.AverageExpense != 10 {
		t.Fatalf("avg expense expected 10 got %f", stats.AverageExpense)
	}
	if stats.Trend != 5 {
		t.Fatalf("trend expected 5 got %f", stats.Trend)
	}
	if stats.Year != 2025 {
		t.Fatalf("year expected 2025 got %d", stats.Year)
	}
}
