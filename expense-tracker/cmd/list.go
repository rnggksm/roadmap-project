package cmd

import (
	"expense-tracker/handler"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		handler.ListExpenses()
	},
}
