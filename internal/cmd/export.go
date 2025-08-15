package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/crush/internal/message"
	"github.com/charmbracelet/crush/internal/session"
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
			return ExportAsJSON(filename, session, messages)
		case "markdown":
			return exportAsMarkdown(filename, session, messages)
		default:
			return fmt.Errorf("unsupported format: %s", format)
		}
	},
}

func exportAsMarkdown(filename string, session session.Session, messages []message.Message) error {
	var content strings.Builder
	
	// Add session header
	content.WriteString(fmt.Sprintf("# Session: %s\n\n", session.Title))
	content.WriteString(fmt.Sprintf("**ID:** %s\n\n", session.ID))
	content.WriteString(fmt.Sprintf("**Created:** %s\n\n", time.Unix(session.CreatedAt, 0).Format("2006-01-02 15:04:05")))
	content.WriteString(fmt.Sprintf("**Last Updated:** %s\n\n", time.Unix(session.UpdatedAt, 0).Format("2006-01-02 15:04:05")))
	content.WriteString(fmt.Sprintf("**Prompt Tokens:** %d\n\n", session.PromptTokens))
	content.WriteString(fmt.Sprintf("**Completion Tokens:** %d\n\n", session.CompletionTokens))
	content.WriteString(fmt.Sprintf("**Cost:** $%.4f\n\n", session.Cost))
	
	// Add conversation
	content.WriteString("## Conversation\n\n")
	
	for _, msg := range messages {
		// Add role header
		var role string
		switch msg.Role {
		case message.User:
			role = "User"
		case message.Assistant:
			role = "Assistant"
		default:
			role = string(msg.Role)
		}
		
		content.WriteString(fmt.Sprintf("### %s (%s)\n\n", role, time.Unix(msg.CreatedAt, 0).Format("2006-01-02 15:04:05")))
		
		// Add message content
		content.WriteString(msg.Content().String())
		content.WriteString("\n\n")
	}
	
	return os.WriteFile(filename, []byte(content.String()), 0644)
}

func init() {
	exportCmd.Flags().StringP("format", "f", "json", "Export format (json, markdown)")
	rootCmd.AddCommand(exportCmd)
}