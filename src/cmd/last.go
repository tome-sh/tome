package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tome/src"
)

var tags []string

// lastCmd represents the last command
var lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Put last command from history into tome.",
	Run: func(cmd *cobra.Command, args []string) {
		parser := tome.NewZshParser(viper.GetString(tome.HISTORY_FILE_PATH_CONFIG_KEY))
		name, err := tome.GetGitConfigSetting("user.name")
		tome.Check(err)
		command := parser.ParseWithTags(name, tags)

		repo := tome.NewGitRepository(viper.GetString(tome.REPOSITORY_CONFIG_KEY))
		err = repo.Store(command)
		tome.Check(err)

		fmt.Printf("Stored command: %s\n", command.String())
	},
}

func init() {
	lastCmd.PersistentFlags().StringSliceVarP(&tags, "tags", "t", tags, "tags for this command")
	rootCmd.AddCommand(lastCmd)
}
