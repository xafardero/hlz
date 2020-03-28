package cmd

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "hlz",
		Short: "A toolbox for holaluz engineers.",
		Long: `hlz is a CLI aplication for holaluz eengineers.
This application is a tool to improve bla bla bla.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Error finding the home direcory:", err.Error())
	}
	viper.SetConfigName("hlz")  // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(home)   // call multiple times to add many search paths

	err = viper.ReadInConfig() // Find and read the config file

	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error config file: %s", err))
	}
}
