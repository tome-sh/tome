package tome

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var defaultBatchSizeKb int64 = 512 * 1000

// Parser is the generic interface to parse shell histories.
type Parser interface {
	Parse(author string) Command
	ParseWithTags(author string, tags []string) Command
}

// Command is a tome command.
type Command struct {
	author    string
	tags      []string
	timestamp int64
	command   string
}

//String() is the string representation of a command.
func (c Command) String() string {
	t := strings.Join(c.tags, ":")
	return fmt.Sprintf("%d;%s;%s;%s", c.timestamp, c.author, t, c.command)
}

// ZshParser is the zsh implementation of parser interface.
type ZshParser struct {
	path      string
	batchSize int64
}

// Parse the second to last line (this will be effectively the last command, as
//`tome last` will be put into the history before we read it)
func (p ZshParser) Parse(author string) Command {
	return p.ParseWithTags(author, []string{})
}

// ParseWithTags parses the second to last line (this will be effectively
// the last command, as `tome last` will be put into the history before we
// read it) and attaches `tags`.
func (p ZshParser) ParseWithTags(author string, tags []string) Command {
	return Command{
		author:    author,
		timestamp: time.Now().Unix(),
		tags:      tags,
		command:   p.getCmd(),
	}
}

func (p ZshParser) getCmd() string {
	line := readSecondToLastLine(p.path, p.batchSize)
	splits := strings.Split(line, ";")
	return splits[len(splits)-1]
}

// NewZshParser creates a zsh parser.
func NewZshParser(p string) Parser {
	return ZshParser{path: p, batchSize: defaultBatchSizeKb}
}

// NewZshParserWithBatchSize creates a zsh parser with custom batchSize.
func NewZshParserWithBatchSize(p string, batchSize int64) Parser {
	return ZshParser{path: p, batchSize: batchSize}
}

func readSecondToLastLine(filePath string, batchSize int64) string {
	file, err := os.Open(filePath)
	Check(err)
	defer file.Close()

	stat, err := os.Stat(filePath)
	Check(err)

	size := stat.Size()
	buf := make([]byte, Min(batchSize, size))
	start := Max(size-batchSize, 0)

	_, err = file.ReadAt(buf, start)
	Check(err)

	splits := strings.Split(string(buf), "\n")

	lastTwoLines := splits[len(splits)-3:]
	return lastTwoLines[0]
}
