package data

import "time"

// Project represents a club project or activity.
type Project struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// Income represents money received for a project.
type Income struct {
	ID          int64     `json:"id"`
	ProjectID   int64     `json:"project_id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

// Expense represents money spent for a project.
type Expense struct {
	ID          int64     `json:"id"`
	ProjectID   int64     `json:"project_id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
