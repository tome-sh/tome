package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tome "tome/src"
)

var id string

var getByIdCmd = &cobra.Command{
	Use:   "get",
	Short: "Prints command with given id.",
	Run: func(cmd *cobra.Command, args []string) {
		repo := tome.NewGitRepository(viper.GetString(tome.REPOSITORY_CONFIG_KEY))
		commands, err := repo.GetAll()
		tome.Check(err)

		command, err := tome.FindById(commands, id)
		tome.Check(err)
		fmt.Print(command.Command)
	},
}

func init() {
	getByIdCmd.PersistentFlags().StringVarP(&id, "id", "", "", "UUID of the command")
	err := getByIdCmd.MarkPersistentFlagRequired("id")
	tome.Check(err)

	rootCmd.AddCommand(getByIdCmd)

}
