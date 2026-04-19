package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "got",
	Short: "Game of Thrones demo data CLI",
	Long: `A CLI tool for working with Game of Thrones character and organization data.

Generate demo data for CRM systems like Salesforce and Pipedrive,
or export data in various formats for testing and development.`,
}

func init() {
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(salesforceCmd)
	rootCmd.AddCommand(pipedriveCmd)
	rootCmd.AddCommand(ringcentralCmd)
}
