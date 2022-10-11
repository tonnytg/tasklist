package entities

type TaskService struct {
	Persistence TaskPersistenceInterface
}

func (s *TaskService) Get(ID uint16) (TaskInterface, error) {
	task, err := s.Persistence.Get(ID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) Create(name string, description string) (TaskInterface, error) {
	task := NewTask()
	task.SetName(name)
	task.SetDescription(description)

	result, err := s.Persistence.Save(task)
	if err != nil {
		return nil, err
	}
	return result, nil
}
