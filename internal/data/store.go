package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Store wraps a sql.DB instance.
type Store struct {
	DB *sql.DB
}

// NewStore returns a new Store.
func NewStore(db *sql.DB) *Store {
	return &Store{DB: db}
}

// Init creates database tables if they do not exist.
func (s *Store) Init(ctx context.Context) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS projects (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            description TEXT,
            created_at DATETIME
        )`,
		`CREATE TABLE IF NOT EXISTS incomes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            project_id INTEGER,
            amount REAL,
            description TEXT,
            date DATETIME,
            FOREIGN KEY(project_id) REFERENCES projects(id)
        )`,
		`CREATE TABLE IF NOT EXISTS expenses (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            project_id INTEGER,
            amount REAL,
            description TEXT,
            date DATETIME,
            FOREIGN KEY(project_id) REFERENCES projects(id)
        )`,
	}

	for _, q := range queries {
		if _, err := s.DB.ExecContext(ctx, q); err != nil {
			return fmt.Errorf("init db: %w", err)
		}
	}
	return nil
}

// --- Project CRUD ---

func (s *Store) CreateProject(ctx context.Context, p *Project) error {
	p.CreatedAt = time.Now()
	res, err := s.DB.ExecContext(ctx, `INSERT INTO projects(name, description, created_at) VALUES(?,?,?)`, p.Name, p.Description, p.CreatedAt)
	if err != nil {
		return err
	}
	p.ID, err = res.LastInsertId()
	return err
}

func (s *Store) GetProject(ctx context.Context, id int64) (Project, error) {
	var p Project
	row := s.DB.QueryRowContext(ctx, `SELECT id, name, description, created_at FROM projects WHERE id = ?`, id)
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt)
	return p, err
}

func (s *Store) UpdateProject(ctx context.Context, p *Project) error {
	_, err := s.DB.ExecContext(ctx, `UPDATE projects SET name=?, description=? WHERE id=?`, p.Name, p.Description, p.ID)
	return err
}

func (s *Store) DeleteProject(ctx context.Context, id int64) error {
	_, err := s.DB.ExecContext(ctx, `DELETE FROM projects WHERE id=?`, id)
	return err
}

func (s *Store) ListProjects(ctx context.Context) ([]Project, error) {
	rows, err := s.DB.QueryContext(ctx, `SELECT id, name, description, created_at FROM projects`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var p Project
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, rows.Err()
}

// --- Income CRUD ---

func (s *Store) CreateIncome(ctx context.Context, in *Income) error {
	res, err := s.DB.ExecContext(ctx, `INSERT INTO incomes(project_id, amount, description, date) VALUES(?,?,?,?)`, in.ProjectID, in.Amount, in.Description, in.Date)
	if err != nil {
		return err
	}
	in.ID, err = res.LastInsertId()
	return err
}

func (s *Store) GetIncome(ctx context.Context, id int64) (Income, error) {
	var in Income
	row := s.DB.QueryRowContext(ctx, `SELECT id, project_id, amount, description, date FROM incomes WHERE id=?`, id)
	err := row.Scan(&in.ID, &in.ProjectID, &in.Amount, &in.Description, &in.Date)
	return in, err
}

func (s *Store) UpdateIncome(ctx context.Context, in *Income) error {
	_, err := s.DB.ExecContext(ctx, `UPDATE incomes SET project_id=?, amount=?, description=?, date=? WHERE id=?`, in.ProjectID, in.Amount, in.Description, in.Date, in.ID)
	return err
}

func (s *Store) DeleteIncome(ctx context.Context, id int64) error {
	_, err := s.DB.ExecContext(ctx, `DELETE FROM incomes WHERE id=?`, id)
	return err
}

func (s *Store) ListIncomesByProject(ctx context.Context, projectID int64) ([]Income, error) {
	rows, err := s.DB.QueryContext(ctx, `SELECT id, project_id, amount, description, date FROM incomes WHERE project_id=?`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Income
	for rows.Next() {
		var in Income
		if err := rows.Scan(&in.ID, &in.ProjectID, &in.Amount, &in.Description, &in.Date); err != nil {
			return nil, err
		}
		items = append(items, in)
	}
	return items, rows.Err()
}

// --- Expense CRUD ---

func (s *Store) CreateExpense(ctx context.Context, ex *Expense) error {
	res, err := s.DB.ExecContext(ctx, `INSERT INTO expenses(project_id, amount, description, date) VALUES(?,?,?,?)`, ex.ProjectID, ex.Amount, ex.Description, ex.Date)
	if err != nil {
		return err
	}
	ex.ID, err = res.LastInsertId()
	return err
}

func (s *Store) GetExpense(ctx context.Context, id int64) (Expense, error) {
	var ex Expense
	row := s.DB.QueryRowContext(ctx, `SELECT id, project_id, amount, description, date FROM expenses WHERE id=?`, id)
	err := row.Scan(&ex.ID, &ex.ProjectID, &ex.Amount, &ex.Description, &ex.Date)
	return ex, err
}

func (s *Store) UpdateExpense(ctx context.Context, ex *Expense) error {
	_, err := s.DB.ExecContext(ctx, `UPDATE expenses SET project_id=?, amount=?, description=?, date=? WHERE id=?`, ex.ProjectID, ex.Amount, ex.Description, ex.Date, ex.ID)
	return err
}

func (s *Store) DeleteExpense(ctx context.Context, id int64) error {
	_, err := s.DB.ExecContext(ctx, `DELETE FROM expenses WHERE id=?`, id)
	return err
}

func (s *Store) ListExpensesByProject(ctx context.Context, projectID int64) ([]Expense, error) {
	rows, err := s.DB.QueryContext(ctx, `SELECT id, project_id, amount, description, date FROM expenses WHERE project_id=?`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Expense
	for rows.Next() {
		var ex Expense
		if err := rows.Scan(&ex.ID, &ex.ProjectID, &ex.Amount, &ex.Description, &ex.Date); err != nil {
			return nil, err
		}
		items = append(items, ex)
	}
	return items, rows.Err()
}
