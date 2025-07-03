package service

import (
	"baristeuer/internal/data"
	"baristeuer/internal/taxlogic"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// DataService provides application methods used by the UI.
type DataService struct {
	store  *data.Store
	logger *slog.Logger
}

// NewDataService creates a new service with the given datastore location.
func NewDataService(dsn string) (*DataService, error) {
	s, err := data.NewStore(dsn)
	if err != nil {
		return nil, fmt.Errorf("create store: %w", err)
	}
	l := newLogger()
	return &DataService{store: s, logger: l}, nil
}

// NewDataServiceFromStore wraps an existing store.
func NewDataServiceFromStore(store *data.Store) *DataService {
	l := newLogger()
	return &DataService{store: store, logger: l}
}

func newLogger() *slog.Logger {
	level := slog.LevelInfo
	if lv := strings.ToLower(os.Getenv("LOG_LEVEL")); lv != "" {
		switch lv {
		case "debug":
			level = slog.LevelDebug
		case "warn":
			level = slog.LevelWarn
		case "error":
			level = slog.LevelError
		}
	}
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	return slog.New(handler)
}

// CreateProject creates a project by name.
func (ds *DataService) CreateProject(name string) (*data.Project, error) {
	p := &data.Project{Name: name}
	if err := ds.store.CreateProject(p); err != nil {
		return nil, fmt.Errorf("create project: %w", err)
	}
	ds.logger.Info("created project", "id", p.ID)
	return p, nil
}

// ListIncomes returns all incomes for the given project.
func (ds *DataService) ListIncomes(projectID int64) ([]data.Income, error) {
	incomes, err := ds.store.ListIncomes(projectID)
	if err != nil {
		return nil, fmt.Errorf("list incomes: %w", err)
	}
	ds.logger.Info("listed incomes", "project", projectID, "count", len(incomes))
	return incomes, nil
}

// AddIncome adds a new income to the given project.
func (ds *DataService) AddIncome(projectID int64, source string, amount float64) (*data.Income, error) {
	i := &data.Income{ProjectID: projectID, Source: source, Amount: amount}
	if err := ds.store.CreateIncome(i); err != nil {
		return nil, fmt.Errorf("create income: %w", err)
	}
	ds.logger.Info("added income", "project", projectID, "amount", amount)
	return i, nil
}

// UpdateIncome updates an existing income entry.
func (ds *DataService) UpdateIncome(id int64, projectID int64, source string, amount float64) error {
	i := &data.Income{ID: id, ProjectID: projectID, Source: source, Amount: amount}
	if err := ds.store.UpdateIncome(i); err != nil {
		return fmt.Errorf("update income: %w", err)
	}
	ds.logger.Info("updated income", "id", id)
	return nil
}

// DeleteIncome removes an income entry by ID.
func (ds *DataService) DeleteIncome(id int64) error {
	if err := ds.store.DeleteIncome(id); err != nil {
		return fmt.Errorf("delete income: %w", err)
	}
	ds.logger.Info("deleted income", "id", id)
	return nil
}

// AddExpense adds a new expense to the given project.
func (ds *DataService) AddExpense(projectID int64, category string, amount float64) (*data.Expense, error) {
	e := &data.Expense{ProjectID: projectID, Category: category, Amount: amount}
	if err := ds.store.CreateExpense(e); err != nil {
		return nil, fmt.Errorf("create expense: %w", err)
	}
	ds.logger.Info("added expense", "project", projectID, "amount", amount)
	return e, nil
}

// UpdateExpense updates an existing expense entry.
func (ds *DataService) UpdateExpense(id int64, projectID int64, category string, amount float64) error {
	e := &data.Expense{ID: id, ProjectID: projectID, Category: category, Amount: amount}
	if err := ds.store.UpdateExpense(e); err != nil {
		return fmt.Errorf("update expense: %w", err)
	}
	ds.logger.Info("updated expense", "id", id)
	return nil
}

// DeleteExpense removes an expense entry by ID.
func (ds *DataService) DeleteExpense(id int64) error {
	if err := ds.store.DeleteExpense(id); err != nil {
		return fmt.Errorf("delete expense: %w", err)
	}
	ds.logger.Info("deleted expense", "id", id)
	return nil
}

// ListExpenses returns all expenses for the given project.
func (ds *DataService) ListExpenses(projectID int64) ([]data.Expense, error) {
	expenses, err := ds.store.ListExpenses(projectID)
	if err != nil {
		return nil, fmt.Errorf("list expenses: %w", err)
	}
	ds.logger.Info("listed expenses", "project", projectID, "count", len(expenses))
	return expenses, nil
}

// AddMember creates a new member.
func (ds *DataService) AddMember(name, email, joinDate string) (*data.Member, error) {
	m := &data.Member{Name: name, Email: email, JoinDate: joinDate}
	if err := ds.store.CreateMember(m); err != nil {
		return nil, fmt.Errorf("create member: %w", err)
	}
	ds.logger.Info("added member", "name", name)
	return m, nil
}

// ListMembers returns all members sorted by name.
func (ds *DataService) ListMembers() ([]data.Member, error) {
	members, err := ds.store.ListMembers()
	if err != nil {
		return nil, fmt.Errorf("list members: %w", err)
	}
	ds.logger.Info("listed members", "count", len(members))
	return members, nil
}

// CalculateProjectTaxes returns a detailed tax calculation for the given project.
func (ds *DataService) CalculateProjectTaxes(projectID int64) (taxlogic.TaxResult, error) {
	revenue, err := ds.store.SumIncomeByProject(projectID)
	if err != nil {
		return taxlogic.TaxResult{}, fmt.Errorf("sum income: %w", err)
	}
	expenses, err := ds.store.SumExpenseByProject(projectID)
	if err != nil {
		return taxlogic.TaxResult{}, fmt.Errorf("sum expense: %w", err)
	}
	result := taxlogic.CalculateTaxes(revenue, expenses)
	ds.logger.Info("calculated taxes", "project", projectID, "total", result.TotalTax)
	return result, nil
}

// Close closes the underlying datastore.
func (ds *DataService) Close() error {
	return ds.store.Close()
}
