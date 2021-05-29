package commands

import (
	"eminentcodex/awsc/cmd/ec2_cmd"

	"github.com/spf13/cobra"
)

func init() {
	r := ec2_cmd.GetEC2Command()
	r.AddCommand(dh)
}

var dh = &cobra.Command{
	Use:   "describe-hosts",
	Short: "Describes the specified Dedicated Hosts or all your Dedicated Hosts",
	Long:  "Describes the specified Dedicated Hosts or all your Dedicated Hosts",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
