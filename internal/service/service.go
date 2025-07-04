package service

import (
	"baristeuer/internal/data"
	"baristeuer/internal/taxlogic"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
)

var ErrInvalidAmount = errors.New("amount must be positive")

func validateAmount(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: %w", ErrInvalidAmount)
	}
	return nil
}

// DataService provides application methods used by the UI.
type DataService struct {
	store     *data.Store
	logger    *slog.Logger
	logCloser io.Closer
}

// NewDataService creates a new service with the given datastore location.
func NewDataService(dsn string, logger *slog.Logger, closer io.Closer) (*DataService, error) {
	s, err := data.NewStore(dsn)
	if err != nil {
		return nil, fmt.Errorf("create store: %w", err)
	}
	if logger == nil {
		logger, closer = NewLogger("", "info", "text")
	}
	return &DataService{store: s, logger: logger, logCloser: closer}, nil
}

// NewDataServiceFromStore wraps an existing store.
func NewDataServiceFromStore(store *data.Store, logger *slog.Logger, closer io.Closer) *DataService {
	if logger == nil {
		logger, closer = NewLogger("", "info", "text")
	}
	return &DataService{store: store, logger: logger, logCloser: closer}
}

// CreateProject creates a project by name.
func (ds *DataService) CreateProject(ctx context.Context, name string) (*data.Project, error) {
	p := &data.Project{Name: name}
	if err := ds.store.CreateProject(ctx, p); err != nil {
		ds.logger.Error("create project failed", "err", err, "name", name)
		return nil, fmt.Errorf("create project: %w", err)
	}
	ds.logger.Info("created project", "id", p.ID)
	return p, nil
}

// ListProjects returns all projects.
func (ds *DataService) ListProjects() ([]data.Project, error) {
	projects, err := ds.store.ListProjects()
	if err != nil {
		ds.logger.Error("list projects failed", "err", err)
		return nil, fmt.Errorf("list projects: %w", err)
	}
	ds.logger.Info("listed projects", "count", len(projects))
	return projects, nil
}

// GetProject fetches a project by ID.
func (ds *DataService) GetProject(id int64) (*data.Project, error) {
	p, err := ds.store.GetProject(context.Background(), id)
	if err != nil {
		ds.logger.Error("get project failed", "err", err, "id", id)
		return nil, fmt.Errorf("get project: %w", err)
	}
	return p, nil
}

// UpdateProject updates a project name.
func (ds *DataService) UpdateProject(id int64, name string) error {
	p := &data.Project{ID: id, Name: name}
	if err := ds.store.UpdateProject(context.Background(), p); err != nil {
		ds.logger.Error("update project failed", "err", err, "id", id)
		return fmt.Errorf("update project: %w", err)
	}
	ds.logger.Info("updated project", "id", id)
	return nil
}

// DeleteProject removes a project by ID.
func (ds *DataService) DeleteProject(id int64) error {
	if err := ds.store.DeleteProject(context.Background(), id); err != nil {
		ds.logger.Error("delete project failed", "err", err, "id", id)
		return fmt.Errorf("delete project: %w", err)
	}
	ds.logger.Info("deleted project", "id", id)
	return nil
}

// ListIncomes returns all incomes for the given project.
func (ds *DataService) ListIncomes(ctx context.Context, projectID int64) ([]data.Income, error) {
	incomes, err := ds.store.ListIncomes(ctx, projectID)
	if err != nil {
		ds.logger.Error("list incomes failed", "err", err, "project", projectID)
		return nil, fmt.Errorf("list incomes: %w", err)
	}
	ds.logger.Info("listed incomes", "project", projectID, "count", len(incomes))
	return incomes, nil
}

// AddIncome adds a new income to the given project.
func (ds *DataService) AddIncome(ctx context.Context, projectID int64, source string, amount float64) (*data.Income, error) {
	if err := validateAmount(amount); err != nil {
		return nil, err
	}
	i := &data.Income{ProjectID: projectID, Source: source, Amount: amount}
	if err := ds.store.CreateIncome(ctx, i); err != nil {
		ds.logger.Error("create income failed", "err", err, "project", projectID)
		return nil, fmt.Errorf("create income: %w", err)
	}
	ds.logger.Info("added income", "project", projectID, "amount", amount)
	return i, nil
}

// UpdateIncome updates an existing income entry.
func (ds *DataService) UpdateIncome(ctx context.Context, id int64, projectID int64, source string, amount float64) error {
	if err := validateAmount(amount); err != nil {
		return err
	}
	i := &data.Income{ID: id, ProjectID: projectID, Source: source, Amount: amount}
	if err := ds.store.UpdateIncome(ctx, i); err != nil {
		ds.logger.Error("update income failed", "err", err, "id", id)
		return fmt.Errorf("update income: %w", err)
	}
	ds.logger.Info("updated income", "id", id)
	return nil
}

// DeleteIncome removes an income entry by ID.
func (ds *DataService) DeleteIncome(ctx context.Context, id int64) error {
	if err := ds.store.DeleteIncome(ctx, id); err != nil {
		ds.logger.Error("delete income failed", "err", err, "id", id)
		return fmt.Errorf("delete income: %w", err)
	}
	ds.logger.Info("deleted income", "id", id)
	return nil
}

// AddExpense adds a new expense to the given project.
func (ds *DataService) AddExpense(ctx context.Context, projectID int64, category string, amount float64) (*data.Expense, error) {
	if err := validateAmount(amount); err != nil {
		return nil, err
	}
	e := &data.Expense{ProjectID: projectID, Category: category, Amount: amount}
	if err := ds.store.CreateExpense(ctx, e); err != nil {
		ds.logger.Error("create expense failed", "err", err, "project", projectID)
		return nil, fmt.Errorf("create expense: %w", err)
	}
	ds.logger.Info("added expense", "project", projectID, "amount", amount)
	return e, nil
}

// UpdateExpense updates an existing expense entry.
func (ds *DataService) UpdateExpense(ctx context.Context, id int64, projectID int64, category string, amount float64) error {
	if err := validateAmount(amount); err != nil {
		return err
	}
	e := &data.Expense{ID: id, ProjectID: projectID, Category: category, Amount: amount}
	if err := ds.store.UpdateExpense(ctx, e); err != nil {
		ds.logger.Error("update expense failed", "err", err, "id", id)
		return fmt.Errorf("update expense: %w", err)
	}
	ds.logger.Info("updated expense", "id", id)
	return nil
}

// DeleteExpense removes an expense entry by ID.
func (ds *DataService) DeleteExpense(ctx context.Context, id int64) error {
	if err := ds.store.DeleteExpense(ctx, id); err != nil {
		ds.logger.Error("delete expense failed", "err", err, "id", id)
		return fmt.Errorf("delete expense: %w", err)
	}
	ds.logger.Info("deleted expense", "id", id)
	return nil
}

// ListExpenses returns all expenses for the given project.
func (ds *DataService) ListExpenses(ctx context.Context, projectID int64) ([]data.Expense, error) {
	expenses, err := ds.store.ListExpenses(ctx, projectID)
	if err != nil {
		ds.logger.Error("list expenses failed", "err", err, "project", projectID)
		return nil, fmt.Errorf("list expenses: %w", err)
	}
	ds.logger.Info("listed expenses", "project", projectID, "count", len(expenses))
	return expenses, nil
}

// AddMember creates a new member.
func (ds *DataService) AddMember(ctx context.Context, name, email, joinDate string) (*data.Member, error) {
	m := &data.Member{Name: name, Email: email, JoinDate: joinDate}
	if err := ds.store.CreateMember(ctx, m); err != nil {
		ds.logger.Error("create member failed", "err", err, "name", name)
		return nil, fmt.Errorf("create member: %w", err)
	}
	ds.logger.Info("added member", "name", name)
	return m, nil
}

// ListMembers returns all members sorted by name.
func (ds *DataService) ListMembers(ctx context.Context) ([]data.Member, error) {
	members, err := ds.store.ListMembers(ctx)
	if err != nil {
		ds.logger.Error("list members failed", "err", err)
		return nil, fmt.Errorf("list members: %w", err)
	}
	ds.logger.Info("listed members", "count", len(members))
	return members, nil
}

// CalculateProjectTaxes returns a detailed tax calculation for the given project.
func (ds *DataService) CalculateProjectTaxes(ctx context.Context, projectID int64) (taxlogic.TaxResult, error) {
	revenue, err := ds.store.SumIncomeByProject(ctx, projectID)
	if err != nil {
		ds.logger.Error("sum income failed", "err", err, "project", projectID)
		return taxlogic.TaxResult{}, fmt.Errorf("sum income: %w", err)
	}
	expenses, err := ds.store.SumExpenseByProject(ctx, projectID)
	if err != nil {
		ds.logger.Error("sum expense failed", "err", err, "project", projectID)
		return taxlogic.TaxResult{}, fmt.Errorf("sum expense: %w", err)
	}
	result := taxlogic.CalculateTaxes(revenue, expenses)
	ds.logger.Info("calculated taxes", "project", projectID, "total", result.TotalTax)
	return result, nil
}

// ExportDatabase copies the underlying SQLite file to the given path.
func (ds *DataService) ExportDatabase(dest string) error {
	srcPath := ds.store.Path()
	in, err := os.Open(srcPath)
	if err != nil {
		ds.logger.Error("open source failed", "err", err, "path", srcPath)
		return fmt.Errorf("open source: %w", err)
	}
	defer in.Close()
	out, err := os.Create(dest)
	if err != nil {
		ds.logger.Error("create dest failed", "err", err, "dest", dest)
		return fmt.Errorf("create dest: %w", err)
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		ds.logger.Error("copy db failed", "err", err, "dest", dest)
		return fmt.Errorf("copy db: %w", err)
	}
	ds.logger.Info("exported database", "dest", dest)
	return nil
}

// Close closes the underlying datastore.
func (ds *DataService) Close() error {
	if ds.logCloser != nil {
		ds.logCloser.Close()
	}
	return ds.store.Close()
}
