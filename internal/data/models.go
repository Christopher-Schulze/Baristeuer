package data

// Project represents a club project with a name.
type Project struct {
	ID     int64  `db:"id"`
	UserID int64  `db:"user_id"`
	Name   string `db:"name"`
}

// Income represents income for a project.
type Income struct {
	ID        int64   `db:"id"`
	ProjectID int64   `db:"project_id"`
	Source    string  `db:"source"`
	Amount    float64 `db:"amount"`
}

// Expense represents an expense for a project.
type Expense struct {
	ID        int64   `db:"id"`
	ProjectID int64   `db:"project_id"`
	Category  string  `db:"category"`
	Amount    float64 `db:"amount"`
}

// Member represents a club member.
type Member struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	JoinDate string `db:"join_date"`
}

// User represents an application user.
type User struct {
	ID           int64  `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}
