package cmd

import (
	"github.com/spf13/cobra"
)

var summaryMonth int
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show a summary of expenses",
	Run: func(cmd *cobra.Command, args []string) {
		expenseHandler.SummaryExpenses(summaryMonth)
	},
}

func init() {
	summaryCmd.Flags().IntVarP(&summaryMonth, "month", "m", 0, "Month for the summary (1-12)")
}
