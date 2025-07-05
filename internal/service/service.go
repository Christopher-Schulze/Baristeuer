package service

import (
	"baristeuer/internal/config"
	"baristeuer/internal/data"
	syncsvc "baristeuer/internal/sync"
	"baristeuer/internal/taxlogic"
	"context"
	"encoding/csv"
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

func selectSyncClient(cfg *config.Config) syncsvc.Client {
	if cfg != nil && cfg.CloudUploadURL != "" && cfg.CloudDownloadURL != "" {
		return syncsvc.NewRemoteClient(cfg.CloudUploadURL, cfg.CloudDownloadURL, cfg.CloudToken)
	}
	return syncsvc.NewLocalClient("")
}

// DataService provides application methods used by the UI.
type DataService struct {
	store      *data.Store
	logger     *slog.Logger
	logCloser  io.Closer
	syncClient syncsvc.Client
}

// NewDataService creates a new service with the given datastore location.
func NewDataService(dsn string, logger *slog.Logger, closer io.Closer, cfg *config.Config) (*DataService, error) {
	s, err := data.NewStore(dsn)
	if err != nil {
		return nil, fmt.Errorf("create store: %w", err)
	}
	if logger == nil {
		logger, closer = NewLogger("", "info", "text")
	}
	client := selectSyncClient(cfg)
	return &DataService{store: s, logger: logger, logCloser: closer, syncClient: client}, nil
}

// NewDataServiceFromStore wraps an existing store.
func NewDataServiceFromStore(store *data.Store, logger *slog.Logger, closer io.Closer, cfg *config.Config) *DataService {
	if logger == nil {
		logger, closer = NewLogger("", "info", "text")
	}
	client := selectSyncClient(cfg)
	return &DataService{store: store, logger: logger, logCloser: closer, syncClient: client}
}

// CreateProject creates a project by name.
func (ds *DataService) CreateProject(ctx context.Context, name string) (*data.Project, error) {
	p := &data.Project{Name: name}
	if err := ds.store.CreateProject(ctx, p); err != nil {
		return nil, fmt.Errorf("create project: %w", err)
	}
	ds.logger.Info("created project", "id", p.ID)
	return p, nil
}

// ListProjects returns all projects.
func (ds *DataService) ListProjects() ([]data.Project, error) {
	projects, err := ds.store.ListProjects()
	if err != nil {
		return nil, fmt.Errorf("list projects: %w", err)
	}
	ds.logger.Info("listed projects", "count", len(projects))
	return projects, nil
}

// GetProject fetches a project by ID.
func (ds *DataService) GetProject(ctx context.Context, id int64) (*data.Project, error) {
	p, err := ds.store.GetProject(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get project: %w", err)
	}
	return p, nil
}

// UpdateProject updates a project name.
func (ds *DataService) UpdateProject(ctx context.Context, id int64, name string) error {
	p := &data.Project{ID: id, Name: name}
	if err := ds.store.UpdateProject(ctx, p); err != nil {
		return fmt.Errorf("update project: %w", err)
	}
	ds.logger.Info("updated project", "id", id)
	return nil
}

// DeleteProject removes a project by ID.
func (ds *DataService) DeleteProject(ctx context.Context, id int64) error {
	if err := ds.store.DeleteProject(ctx, id); err != nil {
		return fmt.Errorf("delete project: %w", err)
	}
	ds.logger.Info("deleted project", "id", id)
	return nil
}

// ListIncomes returns all incomes for the given project.
func (ds *DataService) ListIncomes(ctx context.Context, projectID int64) ([]data.Income, error) {
	incomes, err := ds.store.ListIncomes(ctx, projectID)
	if err != nil {
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
		return fmt.Errorf("update income: %w", err)
	}
	ds.logger.Info("updated income", "id", id)
	return nil
}

// DeleteIncome removes an income entry by ID.
func (ds *DataService) DeleteIncome(ctx context.Context, id int64) error {
	if err := ds.store.DeleteIncome(ctx, id); err != nil {
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
		return fmt.Errorf("update expense: %w", err)
	}
	ds.logger.Info("updated expense", "id", id)
	return nil
}

// DeleteExpense removes an expense entry by ID.
func (ds *DataService) DeleteExpense(ctx context.Context, id int64) error {
	if err := ds.store.DeleteExpense(ctx, id); err != nil {
		return fmt.Errorf("delete expense: %w", err)
	}
	ds.logger.Info("deleted expense", "id", id)
	return nil
}

// ListExpenses returns all expenses for the given project.
func (ds *DataService) ListExpenses(ctx context.Context, projectID int64) ([]data.Expense, error) {
	expenses, err := ds.store.ListExpenses(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("list expenses: %w", err)
	}
	ds.logger.Info("listed expenses", "project", projectID, "count", len(expenses))
	return expenses, nil
}

// AddMember creates a new member.
func (ds *DataService) AddMember(ctx context.Context, name, email, joinDate string) (*data.Member, error) {
	m := &data.Member{Name: name, Email: email, JoinDate: joinDate}
	if err := ds.store.CreateMember(ctx, m); err != nil {
		return nil, fmt.Errorf("create member: %w", err)
	}
	ds.logger.Info("added member", "name", name)
	return m, nil
}

// UpdateMember updates an existing member.
func (ds *DataService) UpdateMember(ctx context.Context, id int64, name, email, joinDate string) error {
	m := &data.Member{ID: id, Name: name, Email: email, JoinDate: joinDate}
	if err := ds.store.UpdateMember(ctx, m); err != nil {
		return fmt.Errorf("update member: %w", err)
	}
	ds.logger.Info("updated member", "id", id)
	return nil
}

// ListMembers returns all members sorted by name.
func (ds *DataService) ListMembers(ctx context.Context) ([]data.Member, error) {
	members, err := ds.store.ListMembers(ctx)
	if err != nil {
		return nil, fmt.Errorf("list members: %w", err)
	}
	ds.logger.Info("listed members", "count", len(members))
	return members, nil
}

// DeleteMember removes a member by ID.
func (ds *DataService) DeleteMember(ctx context.Context, id int64) error {
	if err := ds.store.DeleteMember(ctx, id); err != nil {
		return fmt.Errorf("delete member: %w", err)
	}
	ds.logger.Info("deleted member", "id", id)
	return nil
}

// CalculateProjectTaxes returns a detailed tax calculation for the given project.
func (ds *DataService) CalculateProjectTaxes(ctx context.Context, projectID int64, year int) (taxlogic.TaxResult, error) {
	revenue, err := ds.store.SumIncomeByProject(ctx, projectID)
	if err != nil {
		return taxlogic.TaxResult{}, fmt.Errorf("sum income: %w", err)
	}
	expenses, err := ds.store.SumExpenseByProject(ctx, projectID)
	if err != nil {
		return taxlogic.TaxResult{}, fmt.Errorf("sum expense: %w", err)
	}
	result := taxlogic.CalculateTaxes(revenue, expenses, year)
	ds.logger.Info("calculated taxes", "project", projectID, "total", result.TotalTax)
	return result, nil
}

// ExportDatabase copies the underlying SQLite file to the given path.
func (ds *DataService) ExportDatabase(dest string) error {
	srcPath := ds.store.Path()
	in, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("open source: %w", err)
	}
	defer in.Close()
	out, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("create dest: %w", err)
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy db: %w", err)
	}
	ds.logger.Info("exported database", "dest", dest)
	return nil
}

// SetLogLevel changes the active log level (debug, info, warn, error).
func (ds *DataService) SetLogLevel(level string) {
	SetLogLevel(level)
	ds.logger.Info("log level changed", "level", level)
}

// SetLogFormat updates the log output format (text or json).
func (ds *DataService) SetLogFormat(format string) {
	SetLogFormat(format)
	ds.logger = Logger()
	ds.logger.Info("log format changed", "format", format)
}

// RestoreDatabase replaces the current SQLite file with the one at src.
// The service closes the datastore, copies the file and reopens the connection.
func (ds *DataService) RestoreDatabase(src string) error {
	destPath := ds.store.Path()
	if err := ds.store.Close(); err != nil {
		return fmt.Errorf("close store: %w", err)
	}
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open src: %w", err)
	}
	defer in.Close()
	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("create dest: %w", err)
	}
	if _, err := io.Copy(out, in); err != nil {
		out.Close()
		return fmt.Errorf("copy db: %w", err)
	}
	out.Close()
	store, err := data.NewStore(destPath)
	if err != nil {
		return fmt.Errorf("reopen store: %w", err)
	}
	ds.store = store
	ds.logger.Info("restored database", "src", src)
	return nil
}

// SyncUpload uploads the current database using the configured sync client.
func (ds *DataService) SyncUpload(ctx context.Context) error {
	if ds.syncClient == nil {
		return fmt.Errorf("no sync client configured")
	}
	if err := ds.syncClient.Upload(ctx, ds.store.Path()); err != nil {
		return fmt.Errorf("sync upload: %w", err)
	}
	ds.logger.Info("sync upload complete")
	return nil
}

// SyncDownload downloads the database and replaces the local file.
func (ds *DataService) SyncDownload(ctx context.Context) error {
	if ds.syncClient == nil {
		return fmt.Errorf("no sync client configured")
	}
	destPath := ds.store.Path()
	if err := ds.store.Close(); err != nil {
		return fmt.Errorf("close store: %w", err)
	}
	if err := ds.syncClient.Download(ctx, destPath); err != nil {
		return fmt.Errorf("sync download: %w", err)
	}
	store, err := data.NewStore(destPath)
	if err != nil {
		return fmt.Errorf("reopen store: %w", err)
	}
	ds.store = store
	ds.logger.Info("sync download complete")
	return nil
}

// ExportProjectCSV writes all incomes and expenses of a project into a CSV file.
func (ds *DataService) ExportProjectCSV(ctx context.Context, projectID int64, dest string) error {
	incomes, err := ds.store.ListIncomes(ctx, projectID)
	if err != nil {
		return fmt.Errorf("list incomes: %w", err)
	}
	expenses, err := ds.store.ListExpenses(ctx, projectID)
	if err != nil {
		return fmt.Errorf("list expenses: %w", err)
	}
	f, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	if err := w.Write([]string{"type", "name", "amount"}); err != nil {
		return err
	}
	for _, inc := range incomes {
		w.Write([]string{"income", inc.Source, fmt.Sprintf("%.2f", inc.Amount)})
	}
	for _, exp := range expenses {
		w.Write([]string{"expense", exp.Category, fmt.Sprintf("%.2f", exp.Amount)})
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return fmt.Errorf("write csv: %w", err)
	}
	ds.logger.Info("exported csv", "project", projectID, "dest", dest)
	return nil
}

// Close closes the underlying datastore.
func (ds *DataService) Close() error {
	if ds.logCloser != nil {
		ds.logCloser.Close()
	}
	return ds.store.Close()
}
