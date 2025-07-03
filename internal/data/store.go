package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Store wraps a sql.DB instance.
type Store struct {
	DB *sql.DB
}

// NewStore opens a SQLite database and ensures tables exist.
func NewStore(dsn string) (*Store, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	s := &Store{DB: db}
	if err := s.init(); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}

func (s *Store) init() error {
	schema := []string{
		`CREATE TABLE IF NOT EXISTS projects (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT
        );`,
		`CREATE TABLE IF NOT EXISTS incomes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            project_id INTEGER,
            source TEXT,
            amount REAL
        );`,
		`CREATE TABLE IF NOT EXISTS expenses (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            project_id INTEGER,
            category TEXT,
            amount REAL
        );`,
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
func (s *Store) CreateProject(p *Project) error {
	res, err := s.DB.Exec(`INSERT INTO projects(name) VALUES(?)`, p.Name)
	if err != nil {
		return err
	}
	p.ID, err = res.LastInsertId()
	return err
}

func (s *Store) GetProject(id int64) (*Project, error) {
	row := s.DB.QueryRow(`SELECT id, name FROM projects WHERE id=?`, id)
	var p Project
	if err := row.Scan(&p.ID, &p.Name); err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *Store) UpdateProject(p *Project) error {
	_, err := s.DB.Exec(`UPDATE projects SET name=? WHERE id=?`, p.Name, p.ID)
	return err
}

func (s *Store) DeleteProject(id int64) error {
	_, err := s.DB.Exec(`DELETE FROM projects WHERE id=?`, id)
	return err
}

// CRUD operations for Income
func (s *Store) CreateIncome(i *Income) error {
	res, err := s.DB.Exec(`INSERT INTO incomes(project_id, source, amount) VALUES(?,?,?)`, i.ProjectID, i.Source, i.Amount)
	if err != nil {
		return err
	}
	i.ID, err = res.LastInsertId()
	return err
}

func (s *Store) GetIncome(id int64) (*Income, error) {
	row := s.DB.QueryRow(`SELECT id, project_id, source, amount FROM incomes WHERE id=?`, id)
	var i Income
	if err := row.Scan(&i.ID, &i.ProjectID, &i.Source, &i.Amount); err != nil {
		return nil, err
	}
	return &i, nil
}

func (s *Store) UpdateIncome(i *Income) error {
	_, err := s.DB.Exec(`UPDATE incomes SET project_id=?, source=?, amount=? WHERE id=?`, i.ProjectID, i.Source, i.Amount, i.ID)
	return err
}

func (s *Store) DeleteIncome(id int64) error {
	_, err := s.DB.Exec(`DELETE FROM incomes WHERE id=?`, id)
	return err
}

// CRUD operations for Expense
func (s *Store) CreateExpense(e *Expense) error {
	res, err := s.DB.Exec(`INSERT INTO expenses(project_id, category, amount) VALUES(?,?,?)`, e.ProjectID, e.Category, e.Amount)
	if err != nil {
		return err
	}
	e.ID, err = res.LastInsertId()
	return err
}

func (s *Store) GetExpense(id int64) (*Expense, error) {
	row := s.DB.QueryRow(`SELECT id, project_id, category, amount FROM expenses WHERE id=?`, id)
	var e Expense
	if err := row.Scan(&e.ID, &e.ProjectID, &e.Category, &e.Amount); err != nil {
		return nil, err
	}
	return &e, nil
}

func (s *Store) UpdateExpense(e *Expense) error {
	_, err := s.DB.Exec(`UPDATE expenses SET project_id=?, category=?, amount=? WHERE id=?`, e.ProjectID, e.Category, e.Amount, e.ID)
	return err
}

func (s *Store) DeleteExpense(id int64) error {
	_, err := s.DB.Exec(`DELETE FROM expenses WHERE id=?`, id)
	return err
}

// SumIncomeByProject returns the total income amount for a project.
func (s *Store) SumIncomeByProject(projectID int64) (float64, error) {
	row := s.DB.QueryRow(`SELECT COALESCE(SUM(amount),0) FROM incomes WHERE project_id=?`, projectID)
	var total float64
	if err := row.Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

// SumExpenseByProject returns the total expense amount for a project.
func (s *Store) SumExpenseByProject(projectID int64) (float64, error) {
	row := s.DB.QueryRow(`SELECT COALESCE(SUM(amount),0) FROM expenses WHERE project_id=?`, projectID)
	var total float64
	if err := row.Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}
