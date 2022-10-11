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
