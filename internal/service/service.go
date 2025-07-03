package service

import (
	"baristeuer/internal/data"
	"baristeuer/internal/taxlogic"
	"context"
	"log"
	"os"
)

// DataService provides application methods used by the UI.
type DataService struct {
	store  *data.Store
	logger *log.Logger
}

// NewDataService creates a new service with the given datastore location.
func NewDataService(dsn string) (*DataService, error) {
	s, err := data.NewStore(dsn)
	if err != nil {
		return nil, err
	}
	l := log.New(os.Stdout, "DataService: ", log.LstdFlags)
	return &DataService{store: s, logger: l}, nil
}

// NewDataServiceFromStore wraps an existing store.
func NewDataServiceFromStore(store *data.Store) *DataService {
	l := log.New(os.Stdout, "DataService: ", log.LstdFlags)
	return &DataService{store: store, logger: l}
}

// CreateProject creates a project by name.
func (ds *DataService) CreateProject(ctx context.Context, name string) (*data.Project, error) {
	p := &data.Project{Name: name}
	if err := ds.store.CreateProject(ctx, p); err != nil {
		return nil, err
	}
	ds.logger.Printf("Created project %d", p.ID)
	return p, nil
}

// ListIncomes returns all incomes for the given project.
func (ds *DataService) ListIncomes(ctx context.Context, projectID int64) ([]data.Income, error) {
	rows, err := ds.store.DB.QueryContext(ctx, `SELECT id, project_id, source, amount FROM incomes WHERE project_id=?`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incomes []data.Income
	for rows.Next() {
		var i data.Income
		if err := rows.Scan(&i.ID, &i.ProjectID, &i.Source, &i.Amount); err != nil {
			return nil, err
		}
		incomes = append(incomes, i)
	}
	ds.logger.Printf("Listed %d incomes for project %d", len(incomes), projectID)
	return incomes, nil
}

// AddIncome adds a new income to the given project.
func (ds *DataService) AddIncome(ctx context.Context, projectID int64, source string, amount float64) (*data.Income, error) {
	i := &data.Income{ProjectID: projectID, Source: source, Amount: amount}
	if err := ds.store.CreateIncome(ctx, i); err != nil {
		return nil, err
	}
	ds.logger.Printf("Added income %.2f to project %d", amount, projectID)
	return i, nil
}

// UpdateIncome updates an existing income entry.
func (ds *DataService) UpdateIncome(ctx context.Context, id int64, projectID int64, source string, amount float64) error {
	i := &data.Income{ID: id, ProjectID: projectID, Source: source, Amount: amount}
	if err := ds.store.UpdateIncome(ctx, i); err != nil {
		return err
	}
	ds.logger.Printf("Updated income %d", id)
	return nil
}

// DeleteIncome removes an income entry by ID.
func (ds *DataService) DeleteIncome(ctx context.Context, id int64) error {
	if err := ds.store.DeleteIncome(ctx, id); err != nil {
		return err
	}
	ds.logger.Printf("Deleted income %d", id)
	return nil
}

// AddExpense adds a new expense to the given project.
func (ds *DataService) AddExpense(ctx context.Context, projectID int64, category string, amount float64) (*data.Expense, error) {
	e := &data.Expense{ProjectID: projectID, Category: category, Amount: amount}
	if err := ds.store.CreateExpense(ctx, e); err != nil {
		return nil, err
	}
	ds.logger.Printf("Added expense %.2f to project %d", amount, projectID)
	return e, nil
}

// UpdateExpense updates an existing expense entry.
func (ds *DataService) UpdateExpense(ctx context.Context, id int64, projectID int64, category string, amount float64) error {
	e := &data.Expense{ID: id, ProjectID: projectID, Category: category, Amount: amount}
	if err := ds.store.UpdateExpense(ctx, e); err != nil {
		return err
	}
	ds.logger.Printf("Updated expense %d", id)
	return nil
}

// DeleteExpense removes an expense entry by ID.
func (ds *DataService) DeleteExpense(ctx context.Context, id int64) error {
	if err := ds.store.DeleteExpense(ctx, id); err != nil {
		return err
	}
	ds.logger.Printf("Deleted expense %d", id)
	return nil
}

// ListExpenses returns all expenses for the given project.
func (ds *DataService) ListExpenses(ctx context.Context, projectID int64) ([]data.Expense, error) {
	rows, err := ds.store.DB.QueryContext(ctx, `SELECT id, project_id, category, amount FROM expenses WHERE project_id=?`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []data.Expense
	for rows.Next() {
		var e data.Expense
		if err := rows.Scan(&e.ID, &e.ProjectID, &e.Category, &e.Amount); err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}
	ds.logger.Printf("Listed %d expenses for project %d", len(expenses), projectID)
	return expenses, nil
}

// AddMember creates a new member.
func (ds *DataService) AddMember(ctx context.Context, name, email, joinDate string) (*data.Member, error) {
	m := &data.Member{Name: name, Email: email, JoinDate: joinDate}
	if err := ds.store.CreateMember(ctx, m); err != nil {
		return nil, err
	}
	ds.logger.Printf("Added member %s", name)
	return m, nil
}

// ListMembers returns all members sorted by name.
func (ds *DataService) ListMembers(ctx context.Context) ([]data.Member, error) {
	members, err := ds.store.ListMembers(ctx)
	if err != nil {
		return nil, err
	}
	ds.logger.Printf("Listed %d members", len(members))
	return members, nil
}

// CalculateProjectTaxes returns a detailed tax calculation for the given project.
func (ds *DataService) CalculateProjectTaxes(ctx context.Context, projectID int64) (taxlogic.TaxResult, error) {
	revenue, err := ds.store.SumIncomeByProject(ctx, projectID)
	if err != nil {
		return taxlogic.TaxResult{}, err
	}
	expenses, err := ds.store.SumExpenseByProject(ctx, projectID)
	if err != nil {
		return taxlogic.TaxResult{}, err
	}
	result := taxlogic.CalculateTaxes(revenue, expenses)
	ds.logger.Printf("Calculated taxes for project %d: %.2f EUR", projectID, result.TotalTax)
	return result, nil
}

// Close closes the underlying datastore.
func (ds *DataService) Close() error {
	return ds.store.Close()
}
