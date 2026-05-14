package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense",
	Run: func(cmd *cobra.Command, args []string) {
		expenseHandler.DeleteExpense(deleteID)
	},
}

func init() {
	// flags for the delete command
	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "ID of the expense to delete")

	// validation for required flags
	err := deleteCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Println("Error marking flag as required:", err)
	}
}
