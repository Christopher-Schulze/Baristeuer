package taxlogic

import "time"

// Tax constants for German non-profit organizations for the year 2025.
const (
	// Revenue threshold for economic business operations. If revenue is below this, profits are tax-exempt.
	RevenueExemptionLimit = 45000.00
	// Tax allowance for profits from economic business operations.
	ProfitAllowance = 5000.00
	// Corporate tax rate.
	CorporateTaxRate = 0.15
	// Solidarity surcharge rate (on top of corporate tax).
	SolidaritySurchargeRate = 0.055
)

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

// CalculateTaxes calculates the corporate tax and solidarity surcharge for a non-profit organization.
// It considers the specific tax rules for Germany in 2025 for non-profits.
func CalculateTaxes(revenue, expenses float64) TaxResult {
	profit := revenue - expenses

	result := TaxResult{
		Revenue:               revenue,
		Expenses:              expenses,
		Profit:                profit,
		IsTaxable:             revenue > RevenueExemptionLimit,
		RevenueExemptionLimit: RevenueExemptionLimit,
		ProfitAllowance:       ProfitAllowance,
		Timestamp:             time.Now().Unix(),
	}

	if !result.IsTaxable {
		// If total revenue from economic business operations does not exceed the limit, the profit is tax-free.
		return result
	}

	// If the revenue limit is exceeded, the full profit (minus the allowance) is taxable.
	taxableIncome := profit - ProfitAllowance
	if taxableIncome < 0 {
		taxableIncome = 0
	}
	result.TaxableIncome = taxableIncome

	// Calculate corporate tax (Körperschaftsteuer).
	corporateTax := taxableIncome * CorporateTaxRate
	result.CorporateTax = corporateTax

	// Calculate solidarity surcharge (Solidaritätszuschlag).
	solidaritySurcharge := corporateTax * SolidaritySurchargeRate
	result.SolidaritySurcharge = solidaritySurcharge

	// Calculate total tax liability.
	result.TotalTax = corporateTax + solidaritySurcharge

	return result
}