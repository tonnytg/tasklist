package cli

import (
	"fmt"
	"github.com/tonnytg/tasklist/entities"
)

func Run(
	service entities.TaskServiceInterface,
	action string,
	TaskHash string,
	TaskName string,
	TaskDescription string,
	TaskStatus string,
) (string, error) {

	var result = ""

	switch action {
	case "create":
		task, err := service.Create(TaskName, TaskDescription, TaskStatus)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("%s: has been created", task.GetName())

	case "update":
		task, err := service.Update(TaskHash, TaskName, TaskDescription, TaskStatus)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("[%d] - %s: has been updated", task.GetID(), task.GetName())

	default:
		task, err := service.Get(TaskHash)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("[%d] - %s: has been updated", task.GetID(), task.GetName())
	}
	return result, nil
}
