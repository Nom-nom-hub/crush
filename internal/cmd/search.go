package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search sessions",
	Long:  `Search through session titles or content.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]

		app, err := setupApp(cmd)
		if err != nil {
			return err
		}
		defer app.Shutdown()

		// Get all sessions
		sessions, err := app.Sessions.List(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to list sessions: %w", err)
		}

		// For now, we'll just print a message
		// In a full implementation, we'd search through session titles and content
		fmt.Printf("Searching for: %s\n\n", query)
		fmt.Printf("Found %d sessions:\n\n", len(sessions))
		
		// Print sessions that match the query (case-insensitive)
		queryLower := strings.ToLower(query)
		foundCount := 0
		
		for _, session := range sessions {
			if strings.Contains(strings.ToLower(session.Title), queryLower) {
				fmt.Printf("- %s (%s)\n", session.Title, session.ID)
				foundCount++
			}
		}
		
		if foundCount == 0 {
			fmt.Println("No sessions found matching the query.")
		}
		
		fmt.Println("\n[In a full implementation, this would search through session content as well.]")
		
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}