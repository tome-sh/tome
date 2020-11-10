package tome

import (
	"fmt"
	"os"
)

type Repository interface {
	Store(cmd string) (bool, error)
}

type FileRepository struct {
	path string
}

func NewFileRepository(p string) Repository {
	return FileRepository{path: p}
}

func (r FileRepository) Store(cmd string) (bool, error) {
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
