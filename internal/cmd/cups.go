package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cupsgenerator "github.com/xafardero/cups-gen"
)

func init() {
	rootCmd.AddCommand(cupsCmd)
}

var cupsCmd = &cobra.Command{
	Use:   "cups",
	Short: "Generate random cups",
	Long:  `Generate random cups for testing purposes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cupsgenerator.Generate())
	},
}
