package todo

import "time"

type Todo struct {
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ID          string    `db:"id"`
	Completed   bool      `db:"completed"`
}
