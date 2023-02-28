package main

import (
	"encoding/json"
	"fmt"
	"worker/tasks/emailsend"
)

type Task struct {
	TaskName      string `json:"taskName"`
	CorrelationId string `json:"correlationId"`
	TaskDetails   string `json:"taskDetails"`
}

func parseTask(taskString string) (Task, error) {
	task := Task{}
	err := json.Unmarshal([]byte(taskString), &task)
	return task, err
}

func processTask(task Task) error {
	var err error = nil
	switch task.TaskName {
	case "email-send":
		err = emailsend.HandleTask(task.TaskDetails)
		break
	default:
		err = fmt.Errorf("unsupported task name")
	}
	if err != nil {
		return fmt.Errorf("error while processing task (%s) => %v", task.TaskName, err.Error())
	}
	return err
}
