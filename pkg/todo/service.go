package todo

type Repository interface {
	AddTodo(t Todo) (Todo, error)
	GetTodo(id string) (Todo, error)
	GetTodos() ([]Todo, error)
	UpdateTodo(t Todo) (Todo, error)
	DeleteTodo(id string) error
}

type TodoService struct {
	repository Repository
}

func NewTodoService(r Repository) *TodoService {
	return &TodoService{
		repository: r,
	}
}

func (ts *TodoService) AddTodo(t Todo) (Todo, error) {
	return ts.repository.AddTodo(t)
}

func (ts *TodoService) GetTodo(id string) (Todo, error) {
	return ts.repository.GetTodo(id)
}

func (ts *TodoService) GetTodos() ([]Todo, error) {
	return ts.repository.GetTodos()
}

func (ts *TodoService) UpdateTodo(t Todo) (Todo, error) {
	return ts.repository.UpdateTodo(t)
}

func (ts *TodoService) DeleteTodo(id string) error {
	return ts.repository.DeleteTodo(id)
}
