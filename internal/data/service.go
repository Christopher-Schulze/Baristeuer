package data

import "context"

// DataService handles data operations for expenses.
type DataService struct {
	ctx      context.Context
	expenses []Expense
}

// NewDataService creates a new DataService.
func NewDataService() *DataService {
	return &DataService{
		expenses: []Expense{},
	}
}

// Startup is called when the app starts.
func (s *DataService) Startup(ctx context.Context) {
	s.ctx = ctx
}

// AddIncome adds a new expense to the in-memory store.
// Note: The method is named AddIncome as requested, but it adds an Expense record.
func (s *DataService) AddIncome(description string, amount float64) {
	expense := Expense{
		Category: description,
		Amount:   amount,
	}
	s.expenses = append(s.expenses, expense)
}

// ListExpenses returns all the current expenses.
func (s *DataService) ListExpenses() []Expense {
	return s.expenses
}