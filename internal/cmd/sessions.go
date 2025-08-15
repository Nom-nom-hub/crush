package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

var sessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "Manage chat sessions",
	Long:  `List, switch, or manage your saved chat sessions.`,
}

var listSessionsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved sessions",
	Long:  `Display a list of all saved chat sessions.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := setupApp(cmd)
		if err != nil {
			return err
		}
		defer app.Shutdown()

		sessions, err := app.Sessions.List(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to list sessions: %w", err)
		}

		if len(sessions) == 0 {
			fmt.Println("No sessions found.")
			return nil
		}

		// Create a new tabwriter
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "ID\tTITLE\tCREATED\tLAST USED\t")

		for _, session := range sessions {
			createdAt := time.Unix(session.CreatedAt, 0).Format("2006-01-02 15:04:05")
			updatedAt := time.Unix(session.UpdatedAt, 0).Format("2006-01-02 15:04:05")
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", session.ID, session.Title, createdAt, updatedAt)
		}

		// Flush the tabwriter's buffer to ensure all output is printed
		return w.Flush()
	},
}

var continueSessionCmd = &cobra.Command{
	Use:   "continue [session-id]",
	Short: "Continue a conversation from a saved session",
	Long:  `Continue a conversation from a saved session by providing the session ID.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sessionID := args[0]

		// For now, we'll just start the regular TUI
		// In the future, we could modify this to directly load the specified session
		fmt.Printf("To continue session %s, start Crush and use the session switcher (Ctrl+S)\n", sessionID)
		return nil
	},
}

func init() {
	sessionsCmd.AddCommand(listSessionsCmd)
	sessionsCmd.AddCommand(continueSessionCmd)
	rootCmd.AddCommand(sessionsCmd)
}