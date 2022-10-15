package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/tonnytg/tasklist/entities"
	mock_entities "github.com/tonnytg/tasklist/entities/mocks"
	"github.com/tonnytg/tasklist/pkg/cli"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	TaskHash := "123e4567-e89b-12d3-a456-556642440000"
	TaskName := "TaskNameTaskNameTaskNameTaskNameTaskNameTaskNameTaskNameTaskNameTaskNameTaskName"
	TaskDescription := "TaskDescription"
	TaskStatus := "backlog"

	task := entities.Task{}
	task.SetHash(TaskHash)
	task.SetName(TaskName)
	task.SetDescription(TaskDescription)
	task.SetStatus(TaskStatus)

	TaskMock := mock_entities.NewMockTaskInterface(ctrl)
	TaskMock.EXPECT().GetHash().Return(task.Hash).AnyTimes()
	TaskMock.EXPECT().GetName().Return(task.Name).AnyTimes()
	TaskMock.EXPECT().GetDescription().Return(task.Description).AnyTimes()
	TaskMock.EXPECT().GetStatus().Return(task.Status).AnyTimes()

	service := mock_entities.NewMockTaskServiceInterface(ctrl)
	service.EXPECT().Create(task.Name, task.Description, task.Status).Return(TaskMock, nil)
	service.EXPECT().Get(TaskMock.GetHash()).Return(TaskMock, nil).AnyTimes()
	service.EXPECT().Update(task.Hash, task.Name, task.Description, task.Status).Return(TaskMock, nil).AnyTimes()

	// Test create
	resultExpected := fmt.Sprintf("%s: has been created", TaskMock.GetName())
	result, err := cli.Run(service, "create", task.Hash, task.Name, task.Description, task.Status)
	if err != nil {
		t.Error(err)
	}
	if result != resultExpected {
		t.Errorf("expected %s but got %s", resultExpected, result)
	}

	// Test update
	resultExpected = fmt.Sprintf("[%s] - %s: has been updated", task.GetHash(), task.GetName())
	result, err = cli.Run(service, "update", task.Hash, task.Name, task.Description, task.Status)
	if err != nil {
		t.Error(err)
	}
	if result != resultExpected {
		t.Errorf("expected %s but got %s", resultExpected, result)
	}

	// Test get
	resultExpected = fmt.Sprintf("[%s] - %s: has been found", task.GetHash(), task.GetName())
	result, err = cli.Run(service, "get", task.Hash, "", "", "")
	if err != nil {
		t.Error(err)
	}
	if result != resultExpected {
		t.Errorf("expected %s but got %s", resultExpected, result)
	}
}
