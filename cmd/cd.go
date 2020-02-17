package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(cdCmd)
}

var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Navigate atraves holaluz repositories",
	Long:  `Navigate atraves holaluz repositories...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(viper.GetString("code_path") + "/" + args[0])
	},
}
