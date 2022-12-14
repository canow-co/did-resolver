package cmd

import (
	"github.com/spf13/cobra"
)

func GetRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "canow-did-resolver",
		Short: "Did resolver for the canow method",
	}

	rootCmd.AddCommand(getServeCmd(), getPrintConfigCmd())

	return rootCmd
}
