# Tax Rules Module

This package provides tax calculations for club projects. Rules for each year
are stored in separate JSON files so that new years can be added without code
changes.

## Configuration

For 2025 the file `rules2025.json` defines the following fields:

- `allowance`  – basic tax-free amount
- `rate`       – tax rate applied above the threshold
- `threshold`  – income level after which tax applies

## Usage Example

```go
import "baristeuer/internal/taxrules"

p := taxrules.Project{Revenue: 10000, Expenses: 2000}
tax := taxrules.CalculateTax(p)
```

`tax` will contain the amount due according to the 2025 rules.
