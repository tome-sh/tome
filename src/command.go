package tome

import (
	"fmt"
	"strings"
)

// Command is a tome Command with its metadata
type Command struct {
	Timestamp int64
	Author    string
	Tags      []string
	Command   string
}

func NewCommand(timestamp int64, author string, tags []string, command string) Command {
	return Command{
		Timestamp: timestamp,
		Author:    author,
		Tags:      tags,
		Command:   command,
	}
}

//String() is the string representation of a Command.
func (c Command) String() string {
	t := strings.Join(c.Tags, ":")
	return fmt.Sprintf("%d;%s;%s;%s", c.Timestamp, c.Author, t, c.Command)
}
