package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uritau/sshaws/pkg/cmd/login"
)

var name string
var env string
var app string
var region string
var user string
var sshParam bool

func init() {
	sshCmd.Flags().StringVarP(&env, "env", "e", "*", "env")
	sshCmd.Flags().StringVarP(&app, "app", "a", "*", "app")
	sshCmd.Flags().StringVarP(&region, "region", "r", "eu-west-1", "region")
	sshCmd.Flags().StringVarP(&user, "user", "u", "", "user")
	sshCmd.Flags().StringVarP(&name, "name", "n", "*", "name")
	sshCmd.Flags().BoolVarP(&sshParam, "ssh", "s", false, "ssh")
	rootCmd.AddCommand(sshCmd)
}

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh to your aws instances.",
	Long:  `select and enter to your aws instances via ssh.`,
	Run: func(cmd *cobra.Command, args []string) {

		name := cmd.Flag("name").Value.String()

		if len(args) > 0 {
			name = args[0]
		}

		login.NewLogin(
			name,
			cmd.Flag("region").Value.String(),
			cmd.Flag("user").Value.String(),
			false,
			false,
			false,
		)
	},
}
