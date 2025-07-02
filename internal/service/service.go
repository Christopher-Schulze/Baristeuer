package service

import (
	"baristeuer/internal/data"
)

// DataService provides application methods used by the UI.
type DataService struct {
	store *data.Store
}

// NewDataService creates a new service with the given datastore location.
func NewDataService(dsn string) (*DataService, error) {
	s, err := data.NewStore(dsn)
	if err != nil {
		return nil, err
	}
	return &DataService{store: s}, nil
}

// CreateProject creates a project by name.
func (ds *DataService) CreateProject(name string) (*data.Project, error) {
	p := &data.Project{Name: name}
	if err := ds.store.CreateProject(p); err != nil {
		return nil, err
	}
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
	return incomes, nil
}

// AddExpense adds a new expense to the given project.
func (ds *DataService) AddExpense(projectID int64, category string, amount float64) (*data.Expense, error) {
	e := &data.Expense{ProjectID: projectID, Category: category, Amount: amount}
	if err := ds.store.CreateExpense(e); err != nil {
		return nil, err
	}
	return e, nil
}

// Close closes the underlying datastore.
func (ds *DataService) Close() error {
	return ds.store.Close()
}
