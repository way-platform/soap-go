package call

import "github.com/spf13/cobra"

// NewGroup creates a new [cobra.Group] for the call command.
func NewGroup() *cobra.Group {
	return &cobra.Group{
		ID:    "network",
		Title: "Network Operations",
	}
}
