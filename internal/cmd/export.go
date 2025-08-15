package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export [session-id]",
	Short: "Export a session",
	Long:  `Export a session's conversation history to a file.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sessionID := args[0]
		
		// Get format flag
		format, _ := cmd.Flags().GetString("format")
		
		app, err := setupApp(cmd)
		if err != nil {
			return err
		}
		defer app.Shutdown()

		// Get the session
		session, err := app.Sessions.Get(cmd.Context(), sessionID)
		if err != nil {
			return fmt.Errorf("failed to get session: %w", err)
		}

		// Get the messages
		messages, err := app.Messages.List(cmd.Context(), sessionID)
		if err != nil {
			return fmt.Errorf("failed to get messages: %w", err)
		}

		// Create output filename
		filename := fmt.Sprintf("session_%s_%s.%s", 
			sessionID, 
			time.Now().Format("20060102_150405"),
			format)

		// Export based on format
		switch format {
		case "json":
			return exportAsJSON(filename, session, messages)
		case "markdown":
			return exportAsMarkdown(filename, session, messages)
		default:
			return fmt.Errorf("unsupported format: %s", format)
		}
	},
}

func exportAsJSON(filename string, session interface{}, messages interface{}) error {
	data := map[string]interface{}{
		"session":  session,
		"messages": messages,
	}
	
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	fmt.Printf("Session exported to %s\n", filename)
	return nil
}

func exportAsMarkdown(filename string, session interface{}, messages interface{}) error {
	// For now, we'll just create a simple markdown file
	// In a full implementation, we'd format the messages properly
	content := fmt.Sprintf("# Session Export\n\nSession data would be here in a full implementation.\n")
	
	return os.WriteFile(filename, []byte(content), 0644)
}

func init() {
	exportCmd.Flags().StringP("format", "f", "json", "Export format (json, markdown)")
	rootCmd.AddCommand(exportCmd)
}