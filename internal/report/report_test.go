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

func TestMedian(t *testing.T) {
	if m := Median([]float64{1, 3, 2}); m != 2 {
		t.Fatalf("expected 2 got %f", m)
	}
	if m := Median([]float64{1, 2, 3, 4}); m != 2.5 {
		t.Fatalf("expected 2.5 got %f", m)
	}
	if Median(nil) != 0 {
		t.Fatalf("expected 0 for empty slice")
	}
}

func TestStandardDeviation(t *testing.T) {
	if sd := StandardDeviation([]float64{2, 4}); sd != 1 {
		t.Fatalf("expected 1 got %f", sd)
	}
	if StandardDeviation(nil) != 0 {
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
	if stats.MedianIncome != 15 {
		t.Fatalf("median income expected 15 got %f", stats.MedianIncome)
	}
	if stats.MedianExpense != 10 {
		t.Fatalf("median expense expected 10 got %f", stats.MedianExpense)
	}
	if stats.StdDevIncome != 5 {
		t.Fatalf("stddev income expected 5 got %f", stats.StdDevIncome)
	}
	if stats.StdDevExpense != 5 {
		t.Fatalf("stddev expense expected 5 got %f", stats.StdDevExpense)
	}
	if stats.Trend != 5 {
		t.Fatalf("trend expected 5 got %f", stats.Trend)
	}
	if stats.Year != 2025 {
		t.Fatalf("year expected 2025 got %d", stats.Year)
	}
}
