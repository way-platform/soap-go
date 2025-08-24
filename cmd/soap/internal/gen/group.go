package gen

import "github.com/spf13/cobra"

// NewGroup creates a new [cobra.Group] for the gen command.
func NewGroup() *cobra.Group {
	return &cobra.Group{
		ID:    "gen",
		Title: "Code Generation",
	}
}
