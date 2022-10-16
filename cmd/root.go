package cmd

import (
	"os"

	"github.com/Weidaicheng/len/pkg/length"
	"github.com/Weidaicheng/len/pkg/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "len TEXT",
	Short:                 "Get length",
	Long:                  "Get length of a gavin string",
	Version:               version.Version,
	Args:                  cobra.MinimumNArgs(1),
	SilenceErrors:         true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		length.Print(args)
	},
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
