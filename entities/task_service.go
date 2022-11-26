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

func (s *TaskService) Create(name string, description string, body Body, status string) (TaskInterface, error) {
	task, err := NewTask(name, description, body, status)
	if err != nil {
		return nil, err
	}

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

func (s *TaskService) Update(hash string, name string, description string, body Body, status string) (TaskInterface, error) {
	//task, err := NewTask(name, description, body, status)
	var task *Task
	err := task.SetName(name)
	if err != nil {
		return nil, err
	}
	err = task.SetDescription(description)
	if err != nil {
		return nil, err
	}

	err = task.SetBody(body)
	if err != nil {
		return nil, err
	}

	err = task.SetStatus(status)
	if err != nil {
		return nil, err
	}

	err = task.SetBody(body)
	if err != nil {
		return nil, err
	}

	result, err := s.Persistence.Save(task)
	if err != nil {
		return nil, err
	}
	return result, nil
}
