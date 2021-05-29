package ec2_cmd

import (
	"eminentcodex/awsc/cmd"

	"github.com/spf13/cobra"
)

func init() {
	r := cmd.GetRoot()
	r.AddCommand(ec2Command)
}

var ec2Command = &cobra.Command{
	Use:   "ec2",
	Short: "AWS EC2 command to interact with EC2 service",
	Long:  "AWS EC2 command to interact with EC2 service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func GetEC2Command() *cobra.Command {
	return ec2Command
}
