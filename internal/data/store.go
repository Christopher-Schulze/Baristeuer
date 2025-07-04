package data

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Store wraps a sql.DB instance.
type Store struct {
	DB   *sql.DB
	path string
}

// NewStore opens a SQLite database and ensures tables exist.
func NewStore(dsn string) (*Store, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	s := &Store{DB: db, path: dsn}
	if err := s.init(); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}

// Path returns the database path used to open the store.
func (s *Store) Path() string { return s.path }

func (s *Store) init() error {
	if _, err := s.DB.Exec(`PRAGMA foreign_keys = ON`); err != nil {
		return err
	}
	schema := []string{
		`CREATE TABLE IF NOT EXISTS projects (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT
        );`,
		`CREATE TABLE IF NOT EXISTS incomes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
            source TEXT,
            amount REAL
        );`,
		`CREATE TABLE IF NOT EXISTS expenses (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
            category TEXT,
            amount REAL
        );`,
		`CREATE TABLE IF NOT EXISTS members (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            email TEXT,
            join_date TEXT
        );`,
		`CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE,
            password_hash TEXT
        );`,
		`CREATE INDEX IF NOT EXISTS idx_incomes_project_id ON incomes(project_id);`,
		`CREATE INDEX IF NOT EXISTS idx_expenses_project_id ON expenses(project_id);`,
		`CREATE INDEX IF NOT EXISTS idx_members_name ON members(name);`,
		`CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);`,
	}
	for _, stmt := range schema {
		if _, err := s.DB.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}

// Close closes the underlying database.
func (s *Store) Close() error { return s.DB.Close() }

// CRUD operations for Project
func (s *Store) CreateProject(ctx context.Context, p *Project) error {
	res, err := s.DB.ExecContext(ctx, `INSERT INTO projects(name) VALUES(?)`, p.Name)
	if err != nil {
		return err
	}
	p.ID, err = res.LastInsertId()
	return err
}

func (s *Store) GetProject(ctx context.Context, id int64) (*Project, error) {
	row := s.DB.QueryRowContext(ctx, `SELECT id, name FROM projects WHERE id=?`, id)
	var p Project
	if err := row.Scan(&p.ID, &p.Name); err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *Store) UpdateProject(ctx context.Context, p *Project) error {
	_, err := s.DB.ExecContext(ctx, `UPDATE projects SET name=? WHERE id=?`, p.Name, p.ID)
	return err
}

func (s *Store) DeleteProject(ctx context.Context, id int64) error {
	_, err := s.DB.ExecContext(ctx, `DELETE FROM projects WHERE id=?`, id)
	return err
}

// ListProjects returns all projects ordered by id.
func (s *Store) ListProjects() ([]Project, error) {
	rows, err := s.DB.Query(`SELECT id, name FROM projects ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var p Project
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

// CRUD operations for Income
func (s *Store) CreateIncome(ctx context.Context, i *Income) error {
	res, err := s.DB.ExecContext(ctx, `INSERT INTO incomes(project_id, source, amount) VALUES(?,?,?)`, i.ProjectID, i.Source, i.Amount)
	if err != nil {
		return err
	}
	i.ID, err = res.LastInsertId()
	return err
}

func (s *Store) GetIncome(ctx context.Context, id int64) (*Income, error) {
	row := s.DB.QueryRowContext(ctx, `SELECT id, project_id, source, amount FROM incomes WHERE id=?`, id)
	var i Income
	if err := row.Scan(&i.ID, &i.ProjectID, &i.Source, &i.Amount); err != nil {
		return nil, err
	}
	return &i, nil
}

func (s *Store) UpdateIncome(ctx context.Context, i *Income) error {
	_, err := s.DB.ExecContext(ctx, `UPDATE incomes SET project_id=?, source=?, amount=? WHERE id=?`, i.ProjectID, i.Source, i.Amount, i.ID)
	return err
}

func (s *Store) DeleteIncome(ctx context.Context, id int64) error {
	_, err := s.DB.ExecContext(ctx, `DELETE FROM incomes WHERE id=?`, id)
	return err
}

// ListIncomes returns all incomes for a project ordered by id.
func (s *Store) ListIncomes(ctx context.Context, projectID int64) ([]Income, error) {
	rows, err := s.DB.QueryContext(ctx, `SELECT id, project_id, source, amount FROM incomes WHERE project_id=? ORDER BY id`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Income
	for rows.Next() {
		var i Income
		if err := rows.Scan(&i.ID, &i.ProjectID, &i.Source, &i.Amount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

// CRUD operations for Expense
func (s *Store) CreateExpense(ctx context.Context, e *Expense) error {
	res, err := s.DB.ExecContext(ctx, `INSERT INTO expenses(project_id, category, amount) VALUES(?,?,?)`, e.ProjectID, e.Category, e.Amount)
	if err != nil {
		return err
	}
	e.ID, err = res.LastInsertId()
	return err
}

func (s *Store) GetExpense(ctx context.Context, id int64) (*Expense, error) {
	row := s.DB.QueryRowContext(ctx, `SELECT id, project_id, category, amount FROM expenses WHERE id=?`, id)
	var e Expense
	if err := row.Scan(&e.ID, &e.ProjectID, &e.Category, &e.Amount); err != nil {
		return nil, err
	}
	return &e, nil
}

func (s *Store) UpdateExpense(ctx context.Context, e *Expense) error {
	_, err := s.DB.ExecContext(ctx, `UPDATE expenses SET project_id=?, category=?, amount=? WHERE id=?`, e.ProjectID, e.Category, e.Amount, e.ID)
	return err
}

func (s *Store) DeleteExpense(ctx context.Context, id int64) error {
	_, err := s.DB.ExecContext(ctx, `DELETE FROM expenses WHERE id=?`, id)
	return err
}

// ListExpenses returns all expenses for a project ordered by id.
func (s *Store) ListExpenses(ctx context.Context, projectID int64) ([]Expense, error) {
	rows, err := s.DB.QueryContext(ctx, `SELECT id, project_id, category, amount FROM expenses WHERE project_id=? ORDER BY id`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Expense
	for rows.Next() {
		var e Expense
		if err := rows.Scan(&e.ID, &e.ProjectID, &e.Category, &e.Amount); err != nil {
			return nil, err
		}
		items = append(items, e)
	}
	return items, nil
}

// SumIncomeByProject returns the total income amount for a project.
func (s *Store) SumIncomeByProject(ctx context.Context, projectID int64) (float64, error) {
	row := s.DB.QueryRowContext(ctx, `SELECT COALESCE(SUM(amount),0) FROM incomes WHERE project_id=?`, projectID)
	var total float64
	if err := row.Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

// SumExpenseByProject returns the total expense amount for a project.
func (s *Store) SumExpenseByProject(ctx context.Context, projectID int64) (float64, error) {
	row := s.DB.QueryRowContext(ctx, `SELECT COALESCE(SUM(amount),0) FROM expenses WHERE project_id=?`, projectID)
	var total float64
	if err := row.Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

// CRUD operations for Member
func (s *Store) CreateMember(ctx context.Context, m *Member) error {
	res, err := s.DB.ExecContext(ctx, `INSERT INTO members(name, email, join_date) VALUES(?,?,?)`, m.Name, m.Email, m.JoinDate)
	if err != nil {
		return err
	}
	m.ID, err = res.LastInsertId()
	return err
}

func (s *Store) GetMember(ctx context.Context, id int64) (*Member, error) {
	row := s.DB.QueryRowContext(ctx, `SELECT id, name, email, join_date FROM members WHERE id=?`, id)
	var m Member
	if err := row.Scan(&m.ID, &m.Name, &m.Email, &m.JoinDate); err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *Store) UpdateMember(ctx context.Context, m *Member) error {
	_, err := s.DB.ExecContext(ctx, `UPDATE members SET name=?, email=?, join_date=? WHERE id=?`, m.Name, m.Email, m.JoinDate, m.ID)
	return err
}

func (s *Store) DeleteMember(ctx context.Context, id int64) error {
	_, err := s.DB.ExecContext(ctx, `DELETE FROM members WHERE id=?`, id)
	return err
}

func (s *Store) ListMembers(ctx context.Context) ([]Member, error) {
	rows, err := s.DB.QueryContext(ctx, `SELECT id, name, email, join_date FROM members ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var m Member
		if err := rows.Scan(&m.ID, &m.Name, &m.Email, &m.JoinDate); err != nil {
			return nil, err
		}
		members = append(members, m)
	}
	return members, nil
}

// CRUD operations for User
func (s *Store) CreateUser(ctx context.Context, u *User) error {
	res, err := s.DB.ExecContext(ctx, `INSERT INTO users(username, password_hash) VALUES(?,?)`, u.Username, u.PasswordHash)
	if err != nil {
		return err
	}
	u.ID, err = res.LastInsertId()
	return err
}

func (s *Store) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	row := s.DB.QueryRowContext(ctx, `SELECT id, username, password_hash FROM users WHERE username=?`, username)
	var u User
	if err := row.Scan(&u.ID, &u.Username, &u.PasswordHash); err != nil {
		return nil, err
	}
	return &u, nil
}
