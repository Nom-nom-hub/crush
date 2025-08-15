package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var providersCmd = &cobra.Command{
	Use:   "providers",
	Short: "List available providers and models",
	Long:  `Display a list of all configured providers and their available models.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := setupApp(cmd)
		if err != nil {
			return err
		}
		defer app.Shutdown()

		// Get the config
		cfg := app.Config()

		// Create a new tabwriter
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "PROVIDER ID\tPROVIDER NAME\tMODEL ID\tMODEL NAME\tCONTEXT WINDOW\t")

		// List configured providers and their models
		for providerID, providerConfig := range cfg.Providers.Seq2() {
			for _, model := range providerConfig.Models {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t\n", 
					providerID, 
					providerConfig.Name, 
					model.ID, 
					model.Name, 
					model.ContextWindow)
			}
		}

		// Flush the tabwriter's buffer to ensure all output is printed
		return w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(providersCmd)
}