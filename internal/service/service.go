package service

import (
	"baristeuer/internal/data"
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
func (ds *DataService) CreateProject(name string) (*data.Project, error) {
	p := &data.Project{Name: name}
	if err := ds.store.CreateProject(p); err != nil {
		return nil, err
	}
	ds.logger.Printf("Created project %d", p.ID)
	return p, nil
}

// ListIncomes returns all incomes for the given project.
func (ds *DataService) ListIncomes(projectID int64) ([]data.Income, error) {
	rows, err := ds.store.DB.Query(`SELECT id, project_id, source, amount FROM incomes WHERE project_id=?`, projectID)
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
func (ds *DataService) AddIncome(projectID int64, source string, amount float64) (*data.Income, error) {
	i := &data.Income{ProjectID: projectID, Source: source, Amount: amount}
	if err := ds.store.CreateIncome(i); err != nil {
		return nil, err
	}
	ds.logger.Printf("Added income %.2f to project %d", amount, projectID)
	return i, nil
}

// AddExpense adds a new expense to the given project.
func (ds *DataService) AddExpense(projectID int64, category string, amount float64) (*data.Expense, error) {
	e := &data.Expense{ProjectID: projectID, Category: category, Amount: amount}
	if err := ds.store.CreateExpense(e); err != nil {
		return nil, err
	}
	ds.logger.Printf("Added expense %.2f to project %d", amount, projectID)
	return e, nil
}

// ListExpenses returns all expenses for the given project.
func (ds *DataService) ListExpenses(projectID int64) ([]data.Expense, error) {
	rows, err := ds.store.DB.Query(`SELECT id, project_id, category, amount FROM expenses WHERE project_id=?`, projectID)
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

// Close closes the underlying datastore.
func (ds *DataService) Close() error {
	return ds.store.Close()
}
