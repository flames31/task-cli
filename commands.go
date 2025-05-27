package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

func CommandAdd(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return errors.New("task name is required for add operation")
	}

	taskName := ctx.Args().Slice()[0]

	err := addTask(taskName)
	if err != nil {
		return fmt.Errorf("error adding task %v : %w", taskName, err)
	}
	fmt.Printf("Task %v added!\n", taskName)
	return nil
}

func CommandUpdate(ctx *cli.Context) error {
	if ctx.Args().Len() < 2 {
		return errors.New("task ID and name is required for update operation")
	}

	args := ctx.Args().Slice()

	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("task ID needs to be a valid int : %w", err)
	}
	taskName := args[1]

	err = updateTask(taskID, taskName)
	if err != nil {
		return fmt.Errorf("error updating task %v : %w", taskName, err)
	}
	fmt.Printf("Task %v updated!\n", taskID)
	return nil
}

func CommandList(ctx *cli.Context) error {
	var tasks []Task
	var err error
	args := ctx.Args().Slice()

	if len(args) == 0 {
		tasks, err = listAllTasks()
	} else if len(args) == 1 {
		tasks, err = listStatus(args[0])
	} else {
		return errors.New("too many arguments passed")
	}
	if err != nil {
		return fmt.Errorf("error listing tasks : %w", err)
	}

	for _, task := range tasks {
		fmt.Printf("Task ID : %v\n", task.ID)
		fmt.Printf("Task description : %v\n", task.Description)
		fmt.Printf("Task status : %v\n", task.Status)
		fmt.Printf("Task created at : %v\n", task.CreatedAt)
		fmt.Printf("Task updated at : %v\n", task.UpdatedAt)
		fmt.Println()
	}

	return nil
}

func CommandDelete(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return errors.New("task ID is required for delete operation")
	}

	taskID, err := strconv.Atoi(ctx.Args().Slice()[0])
	if err != nil {
		return fmt.Errorf("error deleting task %v : %w", taskID, err)
	}

	err = deleteTask(taskID)
	if err != nil {
		return fmt.Errorf("error deleting task %v : %w", taskID, err)
	}
	fmt.Printf("Task %v deleted.\n", taskID)
	return nil
}

func CommandMark(ctx *cli.Context) error {
	args := ctx.Args().Slice()

	if len(args) < 2 {
		return errors.New("need a task ID and status to update")
	}

	taskID, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid task ID : %w", err)
	}
	err = markTask(args[0], taskID)
	if err != nil {
		return fmt.Errorf("error marking task : %w", err)
	}

	fmt.Printf("Task %v marked to %v!\n", taskID, args[0])

	return nil
}

func CommandReset(ctx *cli.Context) error {
	err := reset()
	if err != nil {
		return fmt.Errorf("error reseting : %w", err)
	}

	fmt.Println("Reset done.")
	return nil
}
