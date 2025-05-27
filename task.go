package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const STATUS_TODO = "todo"
const STATUS_PROGRESS = "in-progress"
const STATUS_DONE = "done"

func addTask(taskName string) error {

	data, err := getData()
	if err != nil {
		return fmt.Errorf("error fetching json data : %w", err)
	}

	data.MetaData.TaskCount++

	newTask := Task{
		ID:          data.MetaData.TaskCount,
		Description: taskName,
		Status:      STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if data.Tasks == nil {
		data.Tasks = make(map[int]Task)
	}

	data.Tasks[newTask.ID] = newTask

	err = saveJSON(data)
	if err != nil {
		return fmt.Errorf("error saving json : %w", err)
	}

	return nil
}

func updateTask(taskID int, taskName string) error {

	data, err := getData()
	if err != nil {
		return fmt.Errorf("error fetching json data : %w", err)
	}

	task, ok := data.Tasks[taskID]

	if !ok {
		return errors.New("task for given ID not present")
	}

	task.Description = taskName
	task.UpdatedAt = time.Now()

	data.Tasks[taskID] = task

	err = saveJSON(data)
	if err != nil {
		return fmt.Errorf("error saving json : %w", err)
	}

	return nil
}

func listAllTasks() ([]Task, error) {
	tasks := make([]Task, 0)
	data, err := getData()
	if err != nil {
		return []Task{}, fmt.Errorf("error fetching json data : %w", err)
	}

	for _, task := range data.Tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func deleteTask(taskID int) error {
	data, err := getData()
	if err != nil {
		return fmt.Errorf("error fetching json data : %w", err)
	}

	_, ok := data.Tasks[taskID]

	if !ok {
		return errors.New("given task ID not present")
	}

	delete(data.Tasks, taskID)

	err = saveJSON(data)
	if err != nil {
		return fmt.Errorf("error saving json : %w", err)
	}

	return nil
}

func listStatus(statusName string) ([]Task, error) {
	tasks := make([]Task, 0)
	data, err := getData()
	if err != nil {
		return []Task{}, fmt.Errorf("error fetching json data : %w", err)
	}

	for _, task := range data.Tasks {
		if task.Status == statusName {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func markTask(status string, taskID int) error {
	data, err := getData()
	if err != nil {
		return fmt.Errorf("error fetching json data : %w", err)
	}

	task, ok := data.Tasks[taskID]
	if !ok {
		return errors.New("given task ID not found")
	}

	newStatus, _ := strings.CutPrefix(status, "mark-")

	if newStatus != "done" && newStatus != "in-progress" {
		return errors.New("status can only be done or in-progress")
	}

	task.Status = newStatus

	data.Tasks[taskID] = task
	task.UpdatedAt = time.Now()

	err = saveJSON(data)
	if err != nil {
		return fmt.Errorf("error saving json : %w", err)
	}

	return nil
}

func reset() error {
	data := Data{}
	err := saveJSON(data)
	if err != nil {
		return fmt.Errorf("error saving json : %w", err)
	}

	return nil
}
