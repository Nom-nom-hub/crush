package cmd

import (
	"fmt"
	"os"

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
		
		// For now, we'll just print a message
		// In a full implementation, we'd parse the file and create a new session
		fmt.Printf("Importing session from %s\n", filename)
		fmt.Println("[In a full implementation, this would import the session from the file.]")
		
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}