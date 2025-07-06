package report

import (
	"math"
	"sort"
)

// Average returns the arithmetic mean of values or 0 for an empty slice.
func Average(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// Trend returns the difference between the last and first value.
func Trend(values []float64) float64 {
	if len(values) < 2 {
		return 0
	}
	return values[len(values)-1] - values[0]
}

// Median returns the middle value or the mean of the two middle values.
func Median(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)
	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

// StandardDeviation returns the population standard deviation.
func StandardDeviation(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	mean := Average(values)
	var sum float64
	for _, v := range values {
		diff := v - mean
		sum += diff * diff
	}
	variance := sum / float64(len(values))
	return math.Sqrt(variance)
}

type Statistics struct {
	AverageIncome  float64
	AverageExpense float64
	MedianIncome   float64
	MedianExpense  float64
	StdDevIncome   float64
	StdDevExpense  float64
	Trend          float64
	Year           int
}

// Calculate computes basic statistics from income and expense values.
func Calculate(incomes, expenses []float64, year int) Statistics {
	avgInc := Average(incomes)
	avgExp := Average(expenses)
	return Statistics{
		AverageIncome:  avgInc,
		AverageExpense: avgExp,
		MedianIncome:   Median(incomes),
		MedianExpense:  Median(expenses),
		StdDevIncome:   StandardDeviation(incomes),
		StdDevExpense:  StandardDeviation(expenses),
		Trend:          avgInc - avgExp,
		Year:           year,
	}
}
