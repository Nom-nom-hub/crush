package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "View usage and costs",
	Long:  `View token usage and associated costs for sessions or overall usage.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := setupApp(cmd)
		if err != nil {
			return err
		}
		defer app.Shutdown()

		// Get all sessions to calculate total usage
		sessions, err := app.Sessions.List(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to list sessions: %w", err)
		}

		var totalPromptTokens, totalCompletionTokens int64
		var totalCost float64

		// Calculate totals
		for _, session := range sessions {
			totalPromptTokens += session.PromptTokens
			totalCompletionTokens += session.CompletionTokens
			totalCost += session.Cost
		}

		// Create a new tabwriter
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "METRIC\tVALUE\t")
		fmt.Fprintf(w, "Total Sessions\t%d\t\n", len(sessions))
		fmt.Fprintf(w, "Total Prompt Tokens\t%d\t\n", totalPromptTokens)
		fmt.Fprintf(w, "Total Completion Tokens\t%d\t\n", totalCompletionTokens)
		fmt.Fprintf(w, "Total Cost\t$%.4f\t\n", totalCost)

		// Flush the tabwriter's buffer to ensure all output is printed
		return w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(usageCmd)
}