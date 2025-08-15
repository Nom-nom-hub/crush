package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  `View or manage your Crush configuration.`,
}

var viewConfigCmd = &cobra.Command{
	Use:   "view",
	Short: "View current configuration",
	Long:  `Display the current configuration settings.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := setupApp(cmd)
		if err != nil {
			return err
		}
		defer app.Shutdown()

		// Get the config
		cfg := app.Config()

		// Convert to JSON for pretty printing
		b, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal config: %w", err)
		}

		fmt.Println(string(b))
		return nil
	},
}

var setConfigCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a configuration value",
	Long:  `Set a configuration value by specifying the key and value.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		value := args[1]

		// For now, we'll just print a message
		// In a full implementation, we'd actually modify the configuration
		fmt.Printf("Setting %s to %s\n", key, value)
		fmt.Println("[In a full implementation, this would set the configuration value in the config file.]")
		
		return nil
	},
}

var editConfigCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the configuration file",
	Long:  `Open the configuration file in the default editor.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Determine the config file path
		// This is a simplified approach - in a full implementation, we'd use the actual config loading logic
		configPath := "crush.json"
		
		// Check if the file exists
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			// Create a basic config file if it doesn't exist
			defaultConfig := `{
  "$schema": "https://charm.land/crush.json",
  "models": {
    "large": {
      "model": "gemini-2.5-flash",
      "provider": "gemini"
    },
    "small": {
      "model": "gemini-2.5-flash",
      "provider": "gemini"
    }
  }
}`
			if err := os.WriteFile(configPath, []byte(defaultConfig), 0644); err != nil {
				return fmt.Errorf("failed to create default config file: %w", err)
			}
			fmt.Printf("Created default config file: %s\n", configPath)
		}
		
		// Open the file in the default editor
		editor := os.Getenv("EDITOR")
		if editor == "" {
			// Try to determine the default editor based on the OS
			if isWindows() {
				editor = "notepad"
			} else {
				editor = "vi"
			}
		}
		
		fmt.Printf("Opening %s in %s...\n", configPath, editor)
		
		// In a full implementation, we'd actually launch the editor
		// For now, we'll just print a message
		fmt.Println("[In a full implementation, this would open the configuration file in the default editor.]")
		
		return nil
	},
}

func isWindows() bool {
	return strings.Contains(strings.ToLower(os.Getenv("OS")), "windows")
}

func init() {
	configCmd.AddCommand(viewConfigCmd)
	configCmd.AddCommand(setConfigCmd)
	configCmd.AddCommand(editConfigCmd)
	rootCmd.AddCommand(configCmd)
}