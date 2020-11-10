package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tome/src"
)

// lastCmd represents the last command
var lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Put last command from history into tome.",
	Run: func(cmd *cobra.Command, args []string) {
		parser := tome.NewZshParser(viper.GetString(historyFilePathConfigKey))
		command := parser.Parse()
		repo := tome.NewFileRepository(viper.GetString(repositoryConfigKey))
		_, err := repo.Store(command)
		tome.Check(err)
		fmt.Printf("Stored command: %s\n", command)
	},
}

func init() {
	rootCmd.AddCommand(lastCmd)
}

