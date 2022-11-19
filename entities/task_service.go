package entities

type TaskService struct {
	Persistence TaskPersistenceInterface
}

// TaskServiceInterface orquest the task actions orders
type TaskServiceInterface interface {
	Create(name string, description string, body Body, status string) (TaskInterface, error)
	Get(hash string) (TaskInterface, error)
	Update(hash string, name string, description string, body Body, status string) (TaskInterface, error)
}

func NewTaskService(persistence TaskPersistenceInterface) *TaskService {
	return &TaskService{Persistence: persistence}
}

func (s *TaskService) Create(name string, description string, status string) (TaskInterface, error) {
	task := NewTask()
	task.SetName(name)
	task.SetDescription(description)
	task.SetStatus(status)

	result, err := s.Persistence.Save(task)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *TaskService) Get(hash string) (TaskInterface, error) {
	task, err := s.Persistence.Get(hash)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) Update(ID uint16, name string, description string, status string) (TaskInterface, error) {
	task := NewTask()
	task.ID = ID
	task.SetName(name)
	task.SetDescription(description)
	task.SetStatus(status)

	result, err := s.Persistence.Save(task)
	if err != nil {
		return nil, err
	}
	return result, nil
}
