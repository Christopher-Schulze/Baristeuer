package taxlogic

import (
	"encoding/json"
	"os"
	"time"
)

// TaxConfig holds the constants for a given tax year.
type TaxConfig struct {
	RevenueExemptionLimit   float64
	ProfitAllowance         float64
	CorporateTaxRate        float64
	SolidaritySurchargeRate float64
}

// LoadConfig reads tax parameters from a JSON file.
func LoadConfig(path string) (TaxConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return TaxConfig{}, err
	}
	defer f.Close()
	var cfg TaxConfig
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return TaxConfig{}, err
	}
	return cfg, nil
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

// DefaultConfig2026 returns the tax configuration for the year 2026.
func DefaultConfig2026() TaxConfig {
	return TaxConfig{
		RevenueExemptionLimit:   45000.00,
		ProfitAllowance:         5000.00,
		CorporateTaxRate:        0.15,
		SolidaritySurchargeRate: 0.055,
	}
}

// DefaultConfig2027 returns the tax configuration for the year 2027.
func DefaultConfig2027() TaxConfig {
	return TaxConfig{
		RevenueExemptionLimit:   45000.00,
		ProfitAllowance:         5000.00,
		CorporateTaxRate:        0.15,
		SolidaritySurchargeRate: 0.055,
	}
}

// defaultConfig returns the configuration for the given year, falling back to 2025.
func defaultConfig(year int) TaxConfig {
	switch year {
	case 2027:
		return DefaultConfig2027()
	case 2026:
		return DefaultConfig2026()
	case 2025:
		fallthrough
	default:
		return DefaultConfig2025()
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
	Year                  int     // Tax year used for calculation
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
func CalculateTaxes(revenue, expenses float64, year int) TaxResult {
	result := CalculateTaxesWithConfig(revenue, expenses, defaultConfig(year))
	result.Year = year
	return result
}
