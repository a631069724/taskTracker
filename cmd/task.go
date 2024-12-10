package cmd

import (
	"fmt"
	"strconv"
	"taskTracker/server"
	"taskTracker/task"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		server.Manager.AddTask(task.NewTask(args[0]))
		server.SaveTasks()
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			server.Manager.ListAll()
		} else if args[0] == "done" {
			server.Manager.ListDone()
		} else if args[0] == "todo" {
			server.Manager.ListTodo()
		} else if args[0] == "in-progress" {
			server.Manager.ListDoing()
		}
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		des := args[1]
		server.Manager.UpdateTask(id, des)
		server.SaveTasks()
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
		}
		server.Manager.DeleteTask(id)
		server.SaveTasks()
	},
}

var doneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
		}
		err = server.Manager.Done(id)
		if err != nil {
			fmt.Println(err)
		}
		server.SaveTasks()
	},
}

var doingCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark a task as in-progress",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		err := server.Manager.Doing(id)
		if err != nil {
			fmt.Println(err)
		}
		server.SaveTasks()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(doingCmd)
}
