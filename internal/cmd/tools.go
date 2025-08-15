package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Manage tools",
	Long:  `View or manage the tools available to Crush.`,
}

var listToolsCmd = &cobra.Command{
	Use:   "list",
	Short: "List available tools",
	Long:  `Display the tools that are available for use by Crush.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// For now, we'll just list the built-in tools
		// In the future, we could enhance this to show dynamically loaded tools as well
		
		// Create a new tabwriter
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "TOOL NAME\tDESCRIPTION\t")

		// List built-in tools
		builtinTools := []struct {
			Name        string
			Description string
		}{
			{"bash", "Execute bash commands"},
			{"view", "View file contents"},
			{"edit", "Edit file contents"},
			{"multiedit", "Edit multiple files"},
			{"write", "Write new files"},
			{"fetch", "Fetch content from URLs"},
			{"download", "Download files from URLs"},
			{"glob", "Find files matching a pattern"},
			{"grep", "Search for patterns in files"},
			{"ls", "List directory contents"},
			{"sourcegraph", "Search code with Sourcegraph"},
			{"diagnostics", "Get LSP diagnostics for files"},
			{"agent", "Launch a new agent with a subset of tools"},
			{"debugger", "Launch a specialized debugging agent"},
			{"architect", "Launch a specialized architecture agent"},
		}

		for _, tool := range builtinTools {
			fmt.Fprintf(w, "%s\t%s\t\n", tool.Name, tool.Description)
		}

		// Flush the tabwriter's buffer to ensure all output is printed
		return w.Flush()
	},
}

func init() {
	toolsCmd.AddCommand(listToolsCmd)
	rootCmd.AddCommand(toolsCmd)
}