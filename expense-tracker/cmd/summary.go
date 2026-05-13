package cmd

import (
	"expense-tracker/handler"

	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show a summary of expenses",
	Run: func(cmd *cobra.Command, args []string) {
		handler.SummaryExpenses()
	},
}
