package taxlogic

import "time"

// TaxConfig holds the constants for a given tax year.
type TaxConfig struct {
	RevenueExemptionLimit   float64
	ProfitAllowance         float64
	CorporateTaxRate        float64
	SolidaritySurchargeRate float64
}

// DefaultConfig2025 returns the tax configuration for the year 2025.
func DefaultConfig2025() TaxConfig {
	return TaxConfig{
		RevenueExemptionLimit:   45000.00,
		ProfitAllowance:         5000.00,
		CorporateTaxRate:        0.15,
		SolidaritySurchargeRate: 0.055,
	}
}

// TaxResult holds the detailed results of a tax calculation.
type TaxResult struct {
	Revenue               float64 // Total revenue
	Expenses              float64 // Total expenses
	Profit                float64 // Calculated profit (Revenue - Expenses)
	TaxableIncome         float64 // Income subject to tax after allowances
	IsTaxable             bool    // True if revenue exceeds the exemption limit
	CorporateTax          float64 // Calculated corporate tax
	SolidaritySurcharge   float64 // Calculated solidarity surcharge
	TotalTax              float64 // Total tax liability
	RevenueExemptionLimit float64 // The revenue exemption limit used for calculation
	ProfitAllowance       float64 // The profit allowance used for calculation
	Timestamp             int64   // Unix timestamp of the calculation
}

// CalculateTaxesWithConfig calculates taxes using the provided configuration.
func CalculateTaxesWithConfig(revenue, expenses float64, cfg TaxConfig) TaxResult {
	profit := revenue - expenses

	result := TaxResult{
		Revenue:               revenue,
		Expenses:              expenses,
		Profit:                profit,
		IsTaxable:             revenue > cfg.RevenueExemptionLimit,
		RevenueExemptionLimit: cfg.RevenueExemptionLimit,
		ProfitAllowance:       cfg.ProfitAllowance,
		Timestamp:             time.Now().Unix(),
	}

	if !result.IsTaxable {
		// If total revenue from economic business operations does not exceed the limit, the profit is tax-free.
		return result
	}

	// If the revenue limit is exceeded, the full profit (minus the allowance) is taxable.
	taxableIncome := profit - cfg.ProfitAllowance
	if taxableIncome < 0 {
		taxableIncome = 0
	}
	result.TaxableIncome = taxableIncome

	// Calculate corporate tax (Körperschaftsteuer).
	corporateTax := taxableIncome * cfg.CorporateTaxRate
	result.CorporateTax = corporateTax

	// Calculate solidarity surcharge (Solidaritätszuschlag).
	solidaritySurcharge := corporateTax * cfg.SolidaritySurchargeRate
	result.SolidaritySurcharge = solidaritySurcharge

	// Calculate total tax liability.
	result.TotalTax = corporateTax + solidaritySurcharge

	return result
}

// CalculateTaxes calculates taxes using the default configuration for 2025.
func CalculateTaxes(revenue, expenses float64) TaxResult {
	return CalculateTaxesWithConfig(revenue, expenses, DefaultConfig2025())
}
