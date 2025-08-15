package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Manage context paths",
	Long:  `View or manage the context paths used by Crush.`,
}

var listContextCmd = &cobra.Command{
	Use:   "list",
	Short: "List current context paths",
	Long:  `Display the current context paths that are automatically included in prompts.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := setupApp(cmd)
		if err != nil {
			return err
		}
		defer app.Shutdown()

		// Get the config
		cfg := app.Config()

		// List context paths
		if len(cfg.Options.ContextPaths) == 0 {
			fmt.Println("No context paths configured.")
			return nil
		}

		// Create a new tabwriter
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "CONTEXT PATHS\t")

		for _, path := range cfg.Options.ContextPaths {
			fmt.Fprintf(w, "%s\t\n", path)
		}

		// Flush the tabwriter's buffer to ensure all output is printed
		return w.Flush()
	},
}

func init() {
	contextCmd.AddCommand(listContextCmd)
	rootCmd.AddCommand(contextCmd)
}