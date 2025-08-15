package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testModelCmd = &cobra.Command{
	Use:   "test-model [provider] [model] [prompt]",
	Short: "Test a specific model",
	Long:  `Quickly test a specific model with a prompt without starting a full session.`,
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		providerID := args[0]
		modelID := args[1]
		prompt := args[2]

		fmt.Printf("Testing model %s/%s with prompt: %s\n\n", providerID, modelID, prompt)
		fmt.Println("[In a full implementation, this would test the specified model with the given prompt.]")
		
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testModelCmd)
}