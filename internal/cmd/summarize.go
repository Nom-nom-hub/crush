package cmd

import (
	"fmt"

	"github.com/charmbracelet/crush/internal/message"
	"github.com/spf13/cobra"
)

var summarizeCmd = &cobra.Command{
	Use:   "summarize [session-id]",
	Short: "Summarize a session",
	Long:  `Generate a summary of a session's conversation.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sessionID := args[0]

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

		// For now, we'll just print a message
		// In a full implementation, we'd generate a summary using an LLM
		fmt.Printf("Summarizing session: %s\n\n", session.Title)
		fmt.Printf("Session ID: %s\n", session.ID)
		fmt.Printf("Message count: %d\n\n", len(messages))
		
		// Print first and last few messages as a preview
		if len(messages) > 0 {
			fmt.Println("First message:")
			fmt.Println(formatMessage(messages[0]))
			
			if len(messages) > 1 {
				fmt.Println("\nLast message:")
				fmt.Println(formatMessage(messages[len(messages)-1]))
			}
		}
		
		fmt.Println("\n[In a full implementation, this would generate a summary of the session using an LLM.]")
		
		return nil
	},
}

func formatMessage(msg message.Message) string {
	var role string
	switch msg.Role {
	case message.User:
		role = "User"
	case message.Assistant:
		role = "Assistant"
	default:
		role = string(msg.Role)
	}
	
	content := msg.Content().String()
	if len(content) > 200 {
		content = content[:200] + "..."
	}
	
	return fmt.Sprintf("[%s] %s", role, content)
}

func init() {
	rootCmd.AddCommand(summarizeCmd)
}