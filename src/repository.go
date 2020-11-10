package tome

import (
	"fmt"
	"os"
)

// Repository has the methods to store shell commands.
type Repository interface {
	Store(cmd Command) (bool, error)
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
func NewFileRepository(p string) Repository {
	return FileRepository{path: p}
}

// NewGitRepository creates a new GitRepository with a nested FileRepository.
func NewGitRepository(p string) Repository {
	return GitRepository{fileRepository: FileRepository{path: p}}
}

// Store the given cmd in the FileRepository.
func (r FileRepository) Store(cmd Command) (bool, error) {
	f, err := os.OpenFile(r.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%s\n", cmd)); err != nil {
		return false, err
	}

	return true, nil
}

// Store the given cmd in the GitRepository.
func (r GitRepository) Store(cmd Command) (bool, error) {
	_, err := r.fileRepository.Store(cmd)
	if (err != nil) {
		return false, err
	}

	if err = Sync(); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
