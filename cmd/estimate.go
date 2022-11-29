package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func GetPERTEstimate() *cobra.Command {
	var opts pertOpts
	cmd := &cobra.Command{
		Use:   "calculate",
		Short: "Get pert estimate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return calculateEstimate(cmd.Context(), opts)
		},
	}

	cmd.Flags().IntVarP(&opts.Optimistic, "optimistic", "o", 0, "Optimistic estimate from user")
	cmd.Flags().IntVarP(&opts.Realistic, "realistic", "r", 0, "Realistic estimate from user")
	cmd.Flags().IntVarP(&opts.Pessimistic, "pessimistic", "p", 0, "Pessimistic estimate from user")

	return cmd

}

func calculateEstimate(ctx context.Context, opts pertOpts) error {
	const pertNumber int = 6 // TODO: Change name
	const pertCount int = 4  // TODO: Change name

	pertResult := (opts.Optimistic + pertNumber*opts.Realistic + opts.Pessimistic) / pertNumber

	// TODO: Maybe need output result another
	fmt.Printf("PERT estimate is %d", pertResult)

	return nil
}
