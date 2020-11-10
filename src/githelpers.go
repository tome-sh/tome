package tome

import (
	"os/exec"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/viper"
)

func Sync() error {
	repo, err := git.PlainOpen(getDir())
	if (err != nil) {
		return err
	}

	worktree, err := getWorkTree(repo)
	if (err != nil) {
		return err
	}

	err = worktree.Pull(&git.PullOptions{})
	if (err != nil && err.Error() != "already up-to-date") {
		return err
	}

	file := getFile()
	_, err = worktree.Add(file)
	if (err != nil) {
		return err
	}

	userName, err := GetGitConfigSetting("user.name")
	if (err != nil) {
		return err
	}

	email, err := GetGitConfigSetting("user.email")
	if (err != nil) {
		return err
	}

	auth := object.Signature{Name: userName, Email: email, When: time.Now()}
	opts := git.CommitOptions{All: false, Author: &auth, Parents: []plumbing.Hash{}}
	_, err = worktree.Commit("Add command", &opts)
	if (err != nil) {
		return err
	}

	return repo.Push(&git.PushOptions{})
}

func GetGitConfigSetting(key string) (string, error) {
	app := "git"

	arg0 := "config"
	arg1 := "--includes"
	arg2 := "--get"

	cmd := exec.Command(app, arg0, arg1, arg2, key)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}

func getFile() string {
	path := viper.GetString("repository")
	splits := strings.Split(path, "/")
	return splits[len(splits) - 1]
}

func getDir() string {
	path := viper.GetString("repository")
	splits := strings.Split(path, "/")
	return strings.Join(splits[:len(splits) - 1], "/")
}

func getWorkTree(repo *git.Repository) (*git.Worktree, error){
	return repo.Worktree()
}
