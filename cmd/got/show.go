package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/grokify/gameofthrones"
	"github.com/spf13/cobra"
)

var (
	showFormat string
	showQuery  string
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show characters or organizations",
	Long:  `Display Game of Thrones characters or organizations with optional filtering.`,
}

var showCharactersCmd = &cobra.Command{
	Use:   "characters",
	Short: "Show characters",
	Long:  `Display Game of Thrones characters with optional search filtering.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		chars := gameofthrones.Characters()

		// Filter by query if provided
		if showQuery != "" {
			query := strings.ToLower(showQuery)
			var filtered []gameofthrones.Character
			for _, char := range chars {
				if strings.Contains(strings.ToLower(char.Character.DisplayName), query) ||
					strings.Contains(strings.ToLower(char.Actor.DisplayName), query) ||
					strings.Contains(strings.ToLower(char.Organization.Name), query) {
					filtered = append(filtered, char)
				}
			}
			chars = filtered
		}

		return outputCharacters(chars, showFormat)
	},
}

var showOrganizationsCmd = &cobra.Command{
	Use:   "organizations",
	Short: "Show organizations",
	Long:  `Display Game of Thrones organizations with optional search filtering.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		orgs := gameofthrones.Organizations

		// Filter by query if provided
		if showQuery != "" {
			query := strings.ToLower(showQuery)
			var filtered []string
			for _, org := range orgs {
				if strings.Contains(strings.ToLower(org), query) {
					filtered = append(filtered, org)
				}
			}
			orgs = filtered
		}

		return outputOrganizations(orgs, showFormat)
	},
}

func init() {
	showCmd.AddCommand(showCharactersCmd)
	showCmd.AddCommand(showOrganizationsCmd)

	showCharactersCmd.Flags().StringVarP(&showFormat, "format", "f", "text", "Output format (json|text)")
	showCharactersCmd.Flags().StringVarP(&showQuery, "query", "q", "", "Search query to filter results")

	showOrganizationsCmd.Flags().StringVarP(&showFormat, "format", "f", "text", "Output format (json|text)")
	showOrganizationsCmd.Flags().StringVarP(&showQuery, "query", "q", "", "Search query to filter results")
}

func outputCharacters(chars []gameofthrones.Character, format string) error {
	switch strings.ToLower(format) {
	case "json":
		data, err := json.MarshalIndent(chars, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	case "text":
		fmt.Printf("Found %d character(s)\n\n", len(chars))
		for _, char := range chars {
			fmt.Printf("Character: %s\n", char.Character.DisplayName)
			fmt.Printf("  Actor:        %s\n", char.Actor.DisplayName)
			if char.Organization.Name != "" {
				fmt.Printf("  Organization: %s\n", char.Organization.Name)
			}
			fmt.Println()
		}
	default:
		return fmt.Errorf("unknown format: %s (use json or text)", format)
	}
	return nil
}

func outputOrganizations(orgs []string, format string) error {
	switch strings.ToLower(format) {
	case "json":
		data, err := json.MarshalIndent(orgs, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	case "text":
		fmt.Printf("Found %d organization(s)\n\n", len(orgs))
		for _, org := range orgs {
			fmt.Println(org)
		}
	default:
		return fmt.Errorf("unknown format: %s (use json or text)", format)
	}
	return nil
}
