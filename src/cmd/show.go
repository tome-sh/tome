package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tome/src"
)

// lastCmd represents the last command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all commands from tome.",
	Run: func(cmd *cobra.Command, args []string) {
		repo := tome.NewGitRepository(viper.GetString(tome.REPOSITORY_CONFIG_KEY))
		err := tome.Pull()
		if err != nil && err.Error() != "already up-to-date" {
			tome.Check(err)
		}

		lines, err := repo.GetAll()
		tome.Check(err)

		fmt.Printf("%s\n", strings.Join(lines, "\n"))
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
