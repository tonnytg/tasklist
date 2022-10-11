package entities

import (
	"github.com/golang/mock/gomock"
	mock_entities "github.com/tonnytg/tasklist/entities/mocks"
	"testing"
)

func TestTaskService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	task := mock_entities.NewMockTaskInterface(ctrl)
	persistence := mock_entities.NewMockTaskPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(task, nil).AnyTimes()

	service := TaskService{Persistence: persistence}
	result, err := service.Get(1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result != task {
		t.Errorf("Expected %v, got %v", task, result)
	}
}