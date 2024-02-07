package todo

type Repository interface {
	AddTodo(t Todo) (Todo, error)
	GetTodos() ([]Todo, error)
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

func (ts *TodoService) GetTodos() ([]Todo, error) {
	return ts.repository.GetTodos()
}
