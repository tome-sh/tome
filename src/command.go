package tome

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/scanner"
	"time"

	"github.com/google/uuid"
)

// Command is a tome Command with its metadata
type Command struct {
	Id      uuid.UUID
	Time    time.Time
	Author  string
	Tags    []string
	Command string
}

// Helper enum Type for better readability in Deserialize
type CommandFields int

const (
	Id = iota
	Time
	Author
	Tags
	Cmd
)

func NewCommand(
	id uuid.UUID,
	time time.Time,
	author string,
	tags []string,
	command string,
) Command {
	return Command{
		Id:      id,
		Time:    time,
		Author:  author,
		Tags:    tags,
		Command: command,
	}
}

//String() is the string representation of a Command.
func (c Command) String() string {
	t := strings.Join(c.Tags, ":")
	return fmt.Sprintf("%s;%d;%s;%s;%s", c.Id, c.Time.Unix(), c.Author, t, c.EscapedCommandString())
}

func (c Command) Serialize() string {
	t := strings.Join(c.Tags, ":")
	return fmt.Sprintf(": %s;%d;%s;%s;%s", c.Id, c.Time.Unix(), c.Author, t, c.Command)
}

func (c Command) EscapedCommandString() string {
	return strings.ReplaceAll(c.Command, "\n", "\\")
}

func Deserialize(reader io.Reader) ([]Command, error) {
	var s scanner.Scanner
	s.Init(reader)
	s.Whitespace ^= 1<<'\t' | 1<<'\n' | 1<<' '

	lastRune := ' '
	newCommand := true
	fieldIndex := 0

	s.IsIdentRune = func(ch rune, i int) bool {
		oldLastRune := lastRune
		lastRune = ch

		if ch == ':' && oldLastRune == '\n' {
			newCommand = true
			return false
		} else if newCommand && i == 2 {
			newCommand = false
			fieldIndex = 0
			return false
		} else if ch == ';' && fieldIndex < Cmd {
			if i == 0 {
				fieldIndex++
			}
			return false
		} else if ch == scanner.EOF {
			return false
		}
		return true
	}

	var items []Command

	var id uuid.UUID
	var timestamp int64
	var author string
	var tags []string
	var command string
	var lastError error

	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		if token == ';' {
			continue
		}
		text := s.TokenText()
		if text == ": " {
			lastError = nil
			continue
		}

		if lastError != nil {
			continue
		}

		switch fieldIndex {
		case Id:
			id, lastError = uuid.Parse(text)
		case Time:
			timestamp, lastError = strconv.ParseInt(text, 10, 64)
		case Author:
			author = text
		case Tags:
			if len(text) > 0 {
				tags = strings.Split(text, ":")
			} else {
				tags = nil
			}
		case Cmd:
			command = strings.TrimRight(text, "\n")
			if len(command) <= 0 {
				lastError = fmt.Errorf("command with id '%s' is empty", id)
			} else {
				items = append(items, NewCommand(id, time.Unix(timestamp, 0), author, tags, command))
			}
		}

		if lastError != nil {
			Logger.Print(lastError)
		}
	}

	return items, nil
}

func FindById(commands []Command, id string) (Command, error) {
	for _, c := range commands {
		if c.Id.String() == id {
			return c, nil
		}
	}

	return Command{}, fmt.Errorf("command with id '%s' does not exist", id)
}
