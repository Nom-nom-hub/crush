package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import [file]",
	Short: "Import a session",
	Long:  `Import a session from a file.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]
		
		// Check if file exists
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", filename)
		}
		
		// Read the file
		data, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}
		
		// Determine file format based on extension
		var importedData map[string]interface{}
		if strings.HasSuffix(filename, ".json") {
			if err := json.Unmarshal(data, &importedData); err != nil {
				return fmt.Errorf("failed to parse JSON: %w", err)
			}
		} else {
			return fmt.Errorf("unsupported file format: %s", filename)
		}
		
		// For now, we'll just print a message with the imported data
		// In a full implementation, we'd parse the data and create a new session
		fmt.Printf("Importing session from %s\n", filename)
		
		// Print some basic information about the imported session
		if sessionData, ok := importedData["session"].(map[string]interface{}); ok {
			if title, ok := sessionData["Title"].(string); ok {
				fmt.Printf("Session Title: %s\n", title)
			}
			if id, ok := sessionData["ID"].(string); ok {
				fmt.Printf("Session ID: %s\n", id)
			}
		}
		
		fmt.Println("[In a full implementation, this would import the session from the file and create a new session in the database.]")
		
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}