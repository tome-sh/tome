package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tome/src"
)

var historyFilePathConfigKey = "zshHistoryFile"

// lastCmd represents the last command
var lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Put last command from history into tome.",
	Run: func(cmd *cobra.Command, args []string) {
		p := tome.NewZshParser(viper.GetString(historyFilePathConfigKey))
		fmt.Println(p.Parse())
	},
}

func init() {
	rootCmd.AddCommand(lastCmd)
}

