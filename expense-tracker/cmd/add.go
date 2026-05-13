package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	description string
	amount      float64
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding expense")
		fmt.Println("Description:", description)
		fmt.Println("Amount:", amount)
	},
}

func init() {
	// flags for the add command
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the expense")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0.0, "Amount of the expense")

	// validation for required flags
	err := addCmd.MarkFlagRequired("description")
	if err != nil {
		log.Fatal(err)
	}
	err = addCmd.MarkFlagRequired("amount")
	if err != nil {
		log.Fatal(err)
	}

}
