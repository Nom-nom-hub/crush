package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "Manage prompt templates",
	Long:  `List, create, or use prompt templates.`,
}

var listTemplatesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all prompt templates",
	Long:  `Display a list of all saved prompt templates.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// For now, we'll just print a message
		// In a full implementation, we'd list the templates from a file or database
		fmt.Println("Prompt Templates:")
		fmt.Println("=================")
		
		// Sample templates
		templates := []struct {
			Name        string
			Description string
		}{
			{"code-review", "Template for code review requests"},
			{"bug-report", "Template for bug reports"},
			{"feature-request", "Template for feature requests"},
			{"architectural-design", "Template for architectural design discussions"},
		}
		
		// Create a new tabwriter
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "NAME\tDESCRIPTION\t")
		
		for _, template := range templates {
			fmt.Fprintf(w, "%s\t%s\t\n", template.Name, template.Description)
		}
		
		// Flush the tabwriter's buffer to ensure all output is printed
		return w.Flush()
	},
}

var createTemplateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new prompt template",
	Long:  `Create a new prompt template with the given name.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		
		// For now, we'll just print a message
		// In a full implementation, we'd create a new template file or database entry
		fmt.Printf("Creating new template: %s\n", name)
		fmt.Println("[In a full implementation, this would create a new template file or database entry.]")
		
		return nil
	},
}

var useTemplateCmd = &cobra.Command{
	Use:   "use [name]",
	Short: "Use a prompt template",
	Long:  `Use a prompt template to generate a prompt.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		
		// For now, we'll just print a message
		// In a full implementation, we'd load the template and use it
		fmt.Printf("Using template: %s\n", name)
		fmt.Println("[In a full implementation, this would load the template and use it to generate a prompt.]")
		
		return nil
	},
}

func init() {
	templatesCmd.AddCommand(listTemplatesCmd)
	templatesCmd.AddCommand(createTemplateCmd)
	templatesCmd.AddCommand(useTemplateCmd)
	rootCmd.AddCommand(templatesCmd)
}