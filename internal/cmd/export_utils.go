package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/crush/internal/message"
	"github.com/charmbracelet/crush/internal/session"
)

// ExportAsJSON exports session data to a JSON file
func ExportAsJSON(filename string, session session.Session, messages []message.Message) error {
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