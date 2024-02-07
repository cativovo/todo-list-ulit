package todo

import "time"

type Todo struct {
	CreatedAt   time.Time `db:"created_at"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ID          int       `db:"id"`
	Completed   bool      `db:"completed"`
}
