package report

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

type Statistics struct {
	AverageIncome  float64
	AverageExpense float64
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
		Trend:          avgInc - avgExp,
		Year:           year,
	}
}
