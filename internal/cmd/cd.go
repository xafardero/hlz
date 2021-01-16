package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cdCmd)
}

var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Navigate through holaluz repositories",
	Long:  `Navigate through holaluz repositories in github`,
	Run: func(cmd *cobra.Command, args []string) {
		project := ""

		if len(args) > 0 {
			project = args[0]
		}

		fmt.Print(getProjectsDirectory(project))
	},
}
