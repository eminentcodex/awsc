package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	EndpointURL string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&EndpointURL, "endpoint-url", "", "", "Override command’s default URL with the given URL.")
}

var rootCmd = &cobra.Command{
	Use:   "awscli",
	Short: "AWSCLI is a clone of the official aws cli command.",
	Long:  `It's a replacement for the official python aws cli tool`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func GetRoot() *cobra.Command {
	return rootCmd
}
