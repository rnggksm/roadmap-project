/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"expense-tracker/cmd"
	"expense-tracker/internal/service"
)

func main() {
	cmd.Execute()

	_, err := service.NewTrackerService()
	if err != nil {
		panic(err)
	}

}
