package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	updateID          int
	updateDescription string
	updateAmount      float64
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing expense",
	Run: func(cmd *cobra.Command, args []string) {
		expenseHandler.UpdateExpense(updateID, updateDescription, updateAmount)
	},
}

func init() {
	updateCmd.Flags().IntVarP(&updateID, "id", "i", 0, "ID of the expense to update")
	updateCmd.Flags().StringVarP(&updateDescription, "description", "d", "", "New description of the expense")
	updateCmd.Flags().Float64VarP(&updateAmount, "amount", "a", 0.0, "New amount of the expense")

	err := updateCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Println("Error marking flag as required:", err)
	}

	err = updateCmd.MarkFlagRequired("description")
	if err != nil {
		fmt.Println("Error marking flag as required:", err)
	}

	err = updateCmd.MarkFlagRequired("amount")
	if err != nil {
		fmt.Println("Error marking flag as required:", err)
	}
}
