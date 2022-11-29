package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type pertOpts struct {
	Optimistic  int
	Realistic   int
	Pessimistic int
}

// rootCmd represents the base command when called without any subcommands

func main() {
	root := &cobra.Command{
		Use:   "pert-tool",
		Short: "Get PERT distribution from estimate.",
		Long:  `Calculate PERT estimate from optimistic, realistic and pessimistic estimate.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return fmt.Errorf("unknown pert-tool command: %q", args[0])
		},
	}

	root.AddCommand(
		GetPERTEstimate(),
	)
}
