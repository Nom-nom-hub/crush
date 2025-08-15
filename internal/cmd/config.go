package cmd

import (
	"encoding/json"
	"fmt"

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

		fmt.Printf("Setting %s to %s\n", key, value)
		fmt.Println("[In a full implementation, this would set the configuration value.]")
		
		return nil
	},
}

var editConfigCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the configuration file",
	Long:  `Open the configuration file in the default editor.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the config file path
		// For now, we'll just print a message
		// In a full implementation, we'd determine the actual config file path
		fmt.Println("Opening configuration file in default editor...")
		fmt.Println("[In a full implementation, this would open the configuration file in the default editor.]")
		
		return nil
	},
}

func init() {
	configCmd.AddCommand(viewConfigCmd)
	configCmd.AddCommand(setConfigCmd)
	configCmd.AddCommand(editConfigCmd)
	rootCmd.AddCommand(configCmd)
}