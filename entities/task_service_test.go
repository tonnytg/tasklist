package entities_test

import (
	"github.com/golang/mock/gomock"
	"github.com/tonnytg/tasklist/entities"
	"github.com/tonnytg/tasklist/entities/mocks"
	"testing"
)

func TestTaskService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	task := mock_entities.NewMockTaskInterface(ctrl)
	persistence := mock_entities.NewMockTaskPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(task, nil).AnyTimes()

	service := entities.TaskService{
		Persistence: persistence,
	}
	// create hash uuid
	uuidHash := "b6e0c7d0-8c9a-4b9f-9c0e-5e4166513311"
	result, err := service.Get(uuidHash)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result != task {
		t.Errorf("Expected %v, got %v", task, result)
	}
}

func TestTaskCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	task := mock_entities.NewMockTaskInterface(ctrl)
	persistence := mock_entities.NewMockTaskPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(task, nil).AnyTimes()

	service := entities.TaskService{Persistence: persistence}
	result, err := service.Create("test", "test", entities.Body{Content: "testBody"}, "backlog")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result != task {
		t.Errorf("Expected %v, got %v", task, result)
	}
}
