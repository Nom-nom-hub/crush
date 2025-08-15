package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/charmbracelet/crush/internal/llm/tools"
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
			{tools.BashToolName, "Execute bash commands"},
			{tools.ViewToolName, "View file contents"},
			{tools.EditToolName, "Edit file contents"},
			{tools.MultiEditToolName, "Edit multiple files"},
			{tools.WriteToolName, "Write new files"},
			{tools.FetchToolName, "Fetch content from URLs"},
			{tools.DownloadToolName, "Download files from URLs"},
			{tools.GlobToolName, "Find files matching a pattern"},
			{tools.GrepToolName, "Search for patterns in files"},
			{tools.LSToolName, "List directory contents"},
			{tools.SourcegraphToolName, "Search code with Sourcegraph"},
			{tools.DiagnosticsToolName, "Get LSP diagnostics for files"},
			{"Agent", "Launch a new agent with a subset of tools"},
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