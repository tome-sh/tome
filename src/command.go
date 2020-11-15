package tome

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
)

// Command is a tome Command with its metadata
type Command struct {
	Timestamp int64
	Author    string
	Tags      []string
	Command   string
}

var itemRegex = regexp.MustCompile("(?s)(?P<timestamp>[[:digit:]]+);(?P<author>[^;]+);(?P<tags>(?:[[:word:]]+:?)*);(?P<command>.+)")

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
	return fmt.Sprintf(": %d;%s;%s;%s", c.Timestamp, c.Author, t, c.Command)
}

func ParseCommands(reader io.Reader) ([]Command, error) {
	var s scanner.Scanner
	s.Init(reader)
	s.Whitespace ^= 1<<'\t' | 1<<'\n' | 1<<' '

	lastRune := '\n'
	s.IsIdentRune = func(ch rune, i int) bool {
		oldLastRune := lastRune
		lastRune = ch
		return !(ch == ':' && oldLastRune == '\n' || ch == ' ' && i == 0 || ch == scanner.EOF)
	}

	var items []Command

	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		if token == ':' || token == ' ' {
			continue
		}
		text := s.TokenText()

		match := itemRegex.FindStringSubmatch(text)

		if match == nil {
			log.Printf("failed to parse '%s' as Command", text)
			continue
		}

		timestamp, err := strconv.ParseInt(match[itemRegex.SubexpIndex("timestamp")], 10, 64)
		if err != nil {
			log.Print(err.Error())
			continue
		}

		author := match[itemRegex.SubexpIndex("author")]
		tagsString := match[itemRegex.SubexpIndex("tags")]

		var tags []string
		if len(tagsString) > 0 {
			tags = strings.Split(tagsString, ":")
		}

		command := strings.TrimRight(match[itemRegex.SubexpIndex("command")], "\n")
		items = append(items, NewCommand(timestamp, author, tags , command))

	}

	return items, nil

}