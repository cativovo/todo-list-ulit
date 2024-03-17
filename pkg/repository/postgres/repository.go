package repository

import (
	"context"

	tododb "github.com/cativovo/todo-list-ulit/pkg/repository/postgres/generated"
	"github.com/cativovo/todo-list-ulit/pkg/todo"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostgresRepository struct {
	ctx     context.Context
	queries *tododb.Queries
}

func NewPostgresRepository(connStr string) (*PostgresRepository, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{
		ctx:     ctx,
		queries: tododb.New(conn),
	}, nil
}

func (p *PostgresRepository) AddTodo(t todo.Todo) (todo.Todo, error) {
	var description pgtype.Text
	if err := description.Scan(t.Description); err != nil {
		return todo.Todo{}, err
	}

	insertTodoParams := tododb.InsertTodoParams{
		Title:       t.Title,
		Description: description,
		Completed:   t.Completed,
	}

	newTodo, err := p.queries.InsertTodo(p.ctx, insertTodoParams)
	if err != nil {
		return todo.Todo{}, err
	}

	id, err := newTodo.ID.Value()
	if err != nil {
		return todo.Todo{}, err
	}

	return todo.Todo{
		ID:          id.(string),
		Title:       newTodo.Title,
		Description: newTodo.Description.String,
		CreatedAt:   newTodo.CreatedAt.Time,
		UpdatedAt:   newTodo.UpdatedAt.Time,
	}, nil
}

func (p *PostgresRepository) GetTodo(id string) (todo.Todo, error) {
	var uuid pgtype.UUID
	// invalid uuid
	if uuid.Scan(id) != nil {
		return todo.Todo{}, todo.ErrNotFound
	}

	foundTodo, err := p.queries.GetTodo(p.ctx, uuid)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = todo.ErrNotFound
		}

		return todo.Todo{}, err
	}

	return todo.Todo{
		ID:          id,
		Title:       foundTodo.Title,
		Description: foundTodo.Description.String,
		CreatedAt:   foundTodo.CreatedAt.Time,
		UpdatedAt:   foundTodo.UpdatedAt.Time,
	}, nil
}

func (p *PostgresRepository) GetTodos() ([]todo.Todo, error) {
	todos, err := p.queries.GetTodos(p.ctx)
	if err != nil {
		return nil, err
	}

	result := make([]todo.Todo, len(todos))

	for i, v := range todos {
		id, err := v.ID.Value()
		if err != nil {
			return nil, err
		}

		result[i] = todo.Todo{
			ID:          id.(string),
			Title:       v.Title,
			Description: v.Description.String,
			CreatedAt:   v.CreatedAt.Time,
			UpdatedAt:   v.UpdatedAt.Time,
			Completed:   v.Completed,
		}
	}

	return result, nil
}

func (p *PostgresRepository) UpdateTodo(t todo.Todo) (todo.Todo, error) {
	var uuid pgtype.UUID
	// invalid uuid
	if uuid.Scan(t.ID) != nil {
		return todo.Todo{}, todo.ErrNotFound
	}

	var description pgtype.Text
	if err := description.Scan(t.Description); err != nil {
		if err == pgx.ErrNoRows {
			err = todo.ErrNotFound
		}

		return todo.Todo{}, err
	}

	updatedAt, err := p.queries.UpdateTodo(p.ctx, tododb.UpdateTodoParams{
		ID:          uuid,
		Title:       t.Title,
		Description: description,
		Completed:   t.Completed,
	})
	if err != nil {
		return todo.Todo{}, err
	}

	t.UpdatedAt = updatedAt.Time

	return t, nil
}

func (p *PostgresRepository) DeleteTodo(id string) error {
	var uuid pgtype.UUID
	// invalid uuid
	if uuid.Scan(id) != nil {
		return todo.ErrNotFound
	}

	return p.queries.DeleteTodo(p.ctx, uuid)
}
