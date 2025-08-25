package doc

import "github.com/spf13/cobra"

// NewGroup creates a new [cobra.Group] for the doc command.
func NewGroup() *cobra.Group {
	return &cobra.Group{
		ID:    "doc",
		Title: "Documentation",
	}
}
