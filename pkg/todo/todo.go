package todo

// Repository represents a storage medium for storing and retrieving todos.
type Repository interface {
	GetAll() ([]Todo, error)
	GetByID(id int64) (*Todo, error)
	Create(todo *Todo) (*Todo, error)
	Update(todo *Todo) error
	Delete(id int64) error
}

// TodoService represents the contract of providing todo related data.
type TodoService interface {
	GetAll() ([]Todo, error)
	GetByID(id int64) (*Todo, error)
	Create(todo *Todo) (*Todo, error)
	Update(todo *Todo) error
	Delete(id int64) error
}

// A Todo represents an item on the todo list.
type Todo struct {
	// ID is the unique identifier of the todo.
	ID int64 `json:"id"`
	// Title is the actual todo content.
	Title string `json:"title"`
	// Completed represents the completion status.
	Completed bool `json:"completed"`
}

// Service represents a concrete implementation of a TodoService.
type Service struct {
	repo Repository
}

// NewService returns a new implementation of a todo service.
func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) GetAll() ([]Todo, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id int64) (*Todo, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Create(todo *Todo) (*Todo, error) {
	return s.repo.Create(todo)
}

func (s *Service) Update(todo *Todo) error {
	return s.repo.Update(todo)
}

func (s *Service) Delete(id int64) error {
	return s.repo.Delete(id)
}
