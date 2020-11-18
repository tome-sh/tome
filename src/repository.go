package tome

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/viper"
)

// Repository has the methods to store shell commands.
type Repository interface {
	Store(cmd Command) error
	GetAll() ([]Command, error)
}

// FileRepository is a basic kind of repository that simply writes to a file.
type FileRepository struct {
	path string
}

// GitRepository is a repository that writes to a file and pushes to git.
type GitRepository struct {
	fileRepository FileRepository
}

// NewFileRepository creates a new FileRepository.
func NewFileRepository(p string) FileRepository {
	return FileRepository{path: p}
}

// NewGitRepository creates a new GitRepository with a nested FileRepository.
func NewGitRepository(p string) GitRepository {
	return GitRepository{fileRepository: FileRepository{path: p}}
}

// Store the given cmd in the FileRepository.
func (r FileRepository) Store(cmd Command) error {
	f, err := os.OpenFile(r.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%s\n", cmd.Serialize())); err != nil {
		return err
	}

	return nil
}

// Get all commands from the repository.
func (r FileRepository) GetAll() ([]Command, error) {
	f, err := os.Open(r.path)

	if err != nil {
		return []Command{}, err
	}
	defer f.Close()

	return Deserialize(f)
}

// Store the given cmd in the GitRepository.
func (r GitRepository) Store(cmd Command) error {
	err := r.fileRepository.Store(cmd)
	if err != nil {
		return err
	}

	if err = r.sync(); err != nil {
		return err
	} else {
		return nil
	}
}

// Pull changes from git backend.
func (r GitRepository) Pull() error {
	repo, err := git.PlainOpen(getDir())
	if err != nil {
		return err
	}

	worktree, err := getWorkTree(repo)
	if err != nil {
		return err
	}

	return worktree.Pull(&git.PullOptions{})
}

// Get all commands from the repository.
func (r GitRepository) GetAll() ([]Command, error) {
	return r.fileRepository.GetAll()
}

func (r GitRepository) sync() error {
	repo, err := git.PlainOpen(getDir())
	if err != nil {
		return err
	}

	worktree, err := getWorkTree(repo)
	if err != nil {
		return err
	}

	err = worktree.Pull(&git.PullOptions{})
	if err != nil && err.Error() != "already up-to-date" {
		return err
	}

	file := getFile()
	_, err = worktree.Add(file)
	if err != nil {
		return err
	}

	userName, err := GetGitConfigSetting("user.name")
	if err != nil {
		return err
	}

	email, err := GetGitConfigSetting("user.email")
	if err != nil {
		return err
	}

	auth := object.Signature{Name: userName, Email: email, When: time.Now()}
	opts := git.CommitOptions{All: false, Author: &auth, Parents: []plumbing.Hash{}}
	_, err = worktree.Commit("Add command", &opts)
	if err != nil {
		return err
	}

	return repo.Push(&git.PushOptions{})
}

func getFile() string {
	path := viper.GetString("repository")
	splits := strings.Split(path, "/")
	return splits[len(splits)-1]
}

func getDir() string {
	path := viper.GetString("repository")
	splits := strings.Split(path, "/")
	return strings.Join(splits[:len(splits)-1], "/")
}

func getWorkTree(repo *git.Repository) (*git.Worktree, error) {
	return repo.Worktree()
}
