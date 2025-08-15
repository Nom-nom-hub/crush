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

var deleteSessionCmd = &cobra.Command{
	Use:   "delete [session-id]",
	Short: "Delete a saved session",
	Long:  `Delete a saved session by providing the session ID.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sessionID := args[0]

		app, err := setupApp(cmd)
		if err != nil {
			return err
		}
		defer app.Shutdown()

		// Get the session to delete
		session, err := app.Sessions.Get(cmd.Context(), sessionID)
		if err != nil {
			return fmt.Errorf("failed to get session: %w", err)
		}

		// Confirm deletion
		fmt.Printf("Are you sure you want to delete session '%s' (%s)? (y/N): ", session.Title, sessionID)
		var confirmation string
		fmt.Scanln(&confirmation)

		if confirmation != "y" && confirmation != "Y" {
			fmt.Println("Deletion cancelled.")
			return nil
		}

		// Delete the session
		err = app.Sessions.Delete(cmd.Context(), sessionID)
		if err != nil {
			return fmt.Errorf("failed to delete session: %w", err)
		}

		fmt.Printf("Session '%s' (%s) deleted successfully.\n", session.Title, sessionID)
		return nil
	},
}

var deleteAllSessionsCmd = &cobra.Command{
	Use:   "delete-all",
	Short: "Delete all saved sessions",
	Long:  `Delete all saved sessions with confirmation.`,
	RunE: func(cmd *cobra.Command, args []string) error {
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

		if len(sessions) == 0 {
			fmt.Println("No sessions found to delete.")
			return nil
		}

		// Confirm deletion
		fmt.Printf("Are you sure you want to delete all %d sessions? This action cannot be undone. (y/N): ", len(sessions))
		var confirmation string
		fmt.Scanln(&confirmation)

		if confirmation != "y" && confirmation != "Y" {
			fmt.Println("Deletion cancelled.")
			return nil
		}

		// Delete all sessions
		deletedCount := 0
		for _, session := range sessions {
			err = app.Sessions.Delete(cmd.Context(), session.ID)
			if err != nil {
				fmt.Printf("Failed to delete session '%s' (%s): %v\n", session.Title, session.ID, err)
			} else {
				deletedCount++
			}
		}

		fmt.Printf("Deleted %d out of %d sessions.\n", deletedCount, len(sessions))
		return nil
	},
}

var exportAllSessionsCmd = &cobra.Command{
	Use:   "export-all",
	Short: "Export all sessions",
	Long:  `Export all sessions to JSON files.`,
	RunE: func(cmd *cobra.Command, args []string) error {
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

		if len(sessions) == 0 {
			fmt.Println("No sessions found to export.")
			return nil
		}

		// Export each session
		exportedCount := 0
		for _, session := range sessions {
			// Get the messages
			messages, err := app.Messages.List(cmd.Context(), session.ID)
			if err != nil {
				fmt.Printf("Failed to get messages for session '%s' (%s): %v\n", session.Title, session.ID, err)
				continue
			}

			// Create filename
			filename := fmt.Sprintf("session_%s_%s.json", session.ID, time.Now().Format("20060102_150405"))

			// Export session
			err = ExportAsJSON(filename, session, messages)
			if err != nil {
				fmt.Printf("Failed to export session '%s' (%s): %v\n", session.Title, session.ID, err)
			} else {
				exportedCount++
			}
		}

		fmt.Printf("Exported %d out of %d sessions.\n", exportedCount, len(sessions))
		return nil
	},
}

func init() {
	sessionsCmd.AddCommand(listSessionsCmd)
	sessionsCmd.AddCommand(continueSessionCmd)
	sessionsCmd.AddCommand(deleteSessionCmd)
	sessionsCmd.AddCommand(deleteAllSessionsCmd)
	sessionsCmd.AddCommand(exportAllSessionsCmd)
	rootCmd.AddCommand(sessionsCmd)
}