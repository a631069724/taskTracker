package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "A simple task tracker",
}

func Execute() {
	rootCmd.Execute()

}
