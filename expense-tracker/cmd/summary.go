package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show a summary of expenses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Showing summary of expenses")
	},
}
