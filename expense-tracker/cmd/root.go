/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"expense-tracker/handler"
	"expense-tracker/internal/service"
	"log"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "A simple expense tracker CLI",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var expenseHandler *handler.ExpenseHandler

func init() {
	svc, err := service.NewTrackerService()
	if err != nil {
		log.Fatal(err)
	}

	expenseHandler = handler.NewExpenseHandler(svc)
	rootCmd.AddCommand(addCmd, listCmd, deleteCmd, summaryCmd, updateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
