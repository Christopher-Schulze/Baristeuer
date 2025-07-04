package taxlogic

import (
	"fmt"
	"math"
	"testing"
)

// Helper function to compare floats with a tolerance.
func floatEquals(a, b float64) bool {
	const tolerance = 1e-9
	return math.Abs(a-b) < tolerance
}

func TestCalculateTaxes(t *testing.T) {
	years := []int{2025, 2026}
	testCases := []struct {
		name     string
		revenue  float64
		expenses float64
		expected TaxResult
	}{
		{
			name:     "Revenue below exemption limit",
			revenue:  40000.00,
			expenses: 10000.00,
			expected: TaxResult{
				Revenue:               40000.00,
				Expenses:              10000.00,
				Profit:                30000.00,
				IsTaxable:             false,
				TaxableIncome:         0,
				CorporateTax:          0,
				SolidaritySurcharge:   0,
				TotalTax:              0,
				RevenueExemptionLimit: 45000.00,
				ProfitAllowance:       5000.00,
			},
		},
		{
			name:     "Revenue above exemption limit, profit above allowance",
			revenue:  50000.00,
			expenses: 20000.00,
			expected: TaxResult{
				Revenue:               50000.00,
				Expenses:              20000.00,
				Profit:                30000.00,
				IsTaxable:             true,
				TaxableIncome:         25000.00, // 30000 - 5000
				CorporateTax:          3750.00,  // 25000 * 0.15
				SolidaritySurcharge:   206.25,   // 3750 * 0.055
				TotalTax:              3956.25,  // 3750 + 206.25
				RevenueExemptionLimit: 45000.00,
				ProfitAllowance:       5000.00,
			},
		},
		{
			name:     "Revenue above exemption limit, profit below allowance",
			revenue:  46000.00,
			expenses: 42000.00,
			expected: TaxResult{
				Revenue:               46000.00,
				Expenses:              42000.00,
				Profit:                4000.00,
				IsTaxable:             true,
				TaxableIncome:         0, // 4000 - 5000 < 0
				CorporateTax:          0,
				SolidaritySurcharge:   0,
				TotalTax:              0,
				RevenueExemptionLimit: 45000.00,
				ProfitAllowance:       5000.00,
			},
		},
		{
			name:     "Revenue above exemption limit, profit equals allowance",
			revenue:  48000.00,
			expenses: 43000.00,
			expected: TaxResult{
				Revenue:               48000.00,
				Expenses:              43000.00,
				Profit:                5000.00,
				IsTaxable:             true,
				TaxableIncome:         0, // 5000 - 5000
				CorporateTax:          0,
				SolidaritySurcharge:   0,
				TotalTax:              0,
				RevenueExemptionLimit: 45000.00,
				ProfitAllowance:       5000.00,
			},
		},
		{
			name:     "Loss-making scenario",
			revenue:  50000.00,
			expenses: 60000.00,
			expected: TaxResult{
				Revenue:               50000.00,
				Expenses:              60000.00,
				Profit:                -10000.00,
				IsTaxable:             true,
				TaxableIncome:         0, // -10000 - 5000 < 0
				CorporateTax:          0,
				SolidaritySurcharge:   0,
				TotalTax:              0,
				RevenueExemptionLimit: 45000.00,
				ProfitAllowance:       5000.00,
			},
		},
		{
			name:     "Zero profit",
			revenue:  50000.00,
			expenses: 50000.00,
			expected: TaxResult{
				Revenue:               50000.00,
				Expenses:              50000.00,
				Profit:                0.00,
				IsTaxable:             true,
				TaxableIncome:         0, // 0 - 5000 < 0
				CorporateTax:          0,
				SolidaritySurcharge:   0,
				TotalTax:              0,
				RevenueExemptionLimit: 45000.00,
				ProfitAllowance:       5000.00,
			},
		},
	}

	for _, year := range years {
		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%d/%s", year, tc.name), func(t *testing.T) {
				result := CalculateTaxes(tc.revenue, tc.expenses, year)

				if result.IsTaxable != tc.expected.IsTaxable {
					t.Errorf("IsTaxable: got %v, want %v", result.IsTaxable, tc.expected.IsTaxable)
				}
				if !floatEquals(result.Profit, tc.expected.Profit) {
					t.Errorf("Profit: got %f, want %f", result.Profit, tc.expected.Profit)
				}
				if !floatEquals(result.TaxableIncome, tc.expected.TaxableIncome) {
					t.Errorf("TaxableIncome: got %f, want %f", result.TaxableIncome, tc.expected.TaxableIncome)
				}
				if !floatEquals(result.CorporateTax, tc.expected.CorporateTax) {
					t.Errorf("CorporateTax: got %f, want %f", result.CorporateTax, tc.expected.CorporateTax)
				}
				if !floatEquals(result.SolidaritySurcharge, tc.expected.SolidaritySurcharge) {
					t.Errorf("SolidaritySurcharge: got %f, want %f", result.SolidaritySurcharge, tc.expected.SolidaritySurcharge)
				}
				if !floatEquals(result.TotalTax, tc.expected.TotalTax) {
					t.Errorf("TotalTax: got %f, want %f", result.TotalTax, tc.expected.TotalTax)
				}
				if result.Year != year {
					t.Errorf("Year: got %d, want %d", result.Year, year)
				}
			})
		}
	}
}
