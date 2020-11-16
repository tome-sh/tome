package tome

import (
	"fmt"
	"strings"
)

// Command is a tome command.
type Command struct {
	Author    string
	Command   string
	Timestamp int64
	Tags      []string
}

// NewFileRepository creates a new FileRepository.
func NewCommand(author string, command string, timestamp int64, tags []string) Command {
	return Command{
		Author:    author,
		Command:   command,
		Timestamp: timestamp,
		Tags:      tags,
	}
}

//String() is the string representation of a command.
func (c Command) String() string {
	t := strings.Join(c.Tags, ":")
	return fmt.Sprintf("%d;%s;%s;%s", c.Timestamp, c.Author, t, c.Command)
}
