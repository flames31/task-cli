package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Basic task todo CLI app written in go"
	app.Usage = "Use the 'help' command to find more!"
	godotenv.Load(".env")

	app.Commands = []*cli.Command{
		{
			Name:        "add",
			HelpName:    "add",
			Action:      CommandAdd,
			ArgsUsage:   ` `,
			Args:        true,
			Usage:       `Add a task using "add <task_name>"`,
			Description: `Add a task.`,
		},
		{
			Name:        "update",
			HelpName:    "update",
			Action:      CommandUpdate,
			ArgsUsage:   ` `,
			Args:        true,
			Usage:       `Update a task using "update <task_ID> <task_name>"`,
			Description: `Update a task.`,
		},
		{
			Name:        "list",
			HelpName:    "list",
			Action:      CommandList,
			ArgsUsage:   ` `,
			Args:        true,
			Usage:       `List either all tasks using "list" or based on status "list <status>`,
			Description: `List tasks.`,
		},
		{
			Name:        "delete",
			HelpName:    "delete",
			Action:      CommandDelete,
			ArgsUsage:   ` `,
			Args:        true,
			Usage:       `Delete a task using "delete <task_ID>"`,
			Description: `Delete a task.`,
		},
		{
			Name:        "mark",
			HelpName:    "mark tasks to in-progress or done",
			Action:      CommandMark,
			ArgsUsage:   ` `,
			Args:        true,
			Usage:       `Mark a task as in-progress or done using "mark <status> <task_ID>"`,
			Description: `Mark a task as in-progress or done.`,
		},
		{
			Name:        "reset",
			HelpName:    "reset",
			Action:      CommandReset,
			ArgsUsage:   ` `,
			Args:        false,
			Usage:       `Reset using "reset"`,
			Description: `Reset all data.`,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
