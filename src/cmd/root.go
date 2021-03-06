package cmd

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tome "tome/src"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tome",
	Short: "Share shell spells.",
	Long: `Share shell spells with other wizards.
	This application aspires to be a shared zsh/bash history.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	tome.Check(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.tome{.yaml|.json})",
	)
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "set to true to see stack traces")
	tome.Check(viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug")))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	tome.Check(err)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with name ".tome" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tome")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		tome.Logger.Println("no suitable config file found in: ", home)
	}

	requireParam(tome.SHELL_TYPE_CONFIG_KEY)
	requireParam(tome.REPOSITORY_CONFIG_KEY)
}

func requireParam(configKey string) {
	if !viper.IsSet(configKey) {
		tome.Check(fmt.Errorf("missing required config parameter: %s", configKey))
	}
}
