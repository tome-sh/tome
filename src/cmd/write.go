package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tome "tome/src"
)

var tags []string

// writeCmd writes commands into tome
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write command from history into tome.",
	Run: func(cmd *cobra.Command, args []string) {
		author, err := tome.GetUserName()
		tome.Check(err)

		text, err := ioutil.ReadAll(os.Stdin)
		tome.Check(err)

		if len(text) <= 0 {
			tome.Check(fmt.Errorf("cannot add empty command"))
		}
		command := tome.NewCommand(uuid.New(), time.Now(), author, tags, string(text))

		repo := tome.NewGitRepository(viper.GetString(tome.REPOSITORY_CONFIG_KEY))
		err = repo.Store(command)
		tome.Check(err)
	},
}

func init() {
	writeCmd.PersistentFlags().StringSliceVarP(&tags, "tags", "t", tags, "tags for this command")

	rootCmd.AddCommand(writeCmd)
}
