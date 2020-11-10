package tome

import (
	"os"
	"strings"
)

var defaultBatchSizeKb int64 = 512 * 1000

// Parser interface

type Parser interface {
	Parse() string
}

// Zsh parser

type ZshParser struct {
	path string
	batchSize int64
}

func (p ZshParser) Parse() string {
	line := readSecondToLastLine(p.path, p.batchSize)
	splits := strings.Split(line, ";")
	return splits[len(splits) - 1]
}

func NewZshParser(p string) Parser {
	return ZshParser{path: p, batchSize: defaultBatchSizeKb}
}

// NewZshParserWithBatchSize creates a zsh parser with custom batchSize.
func NewZshParserWithBatchSize(p string, batchSize int64) Parser {
	return ZshParser{path: p, batchSize: batchSize}
}

// Helpers

func readSecondToLastLine(filePath string, batchSize int64) string {
	file, err := os.Open(filePath)
	Check(err)
	defer file.Close()

	stat, err := os.Stat(filePath)
	Check(err)

	size := stat.Size()
	buf := make([]byte, batchSize)
	start := size - batchSize
	if (size < batchSize) {
		buf = make([]byte, size)
		start = 0
	}

	_, err = file.ReadAt(buf, start)
	Check(err)

	splits := strings.Split(string(buf), "\n")

	lastTwoLines := splits[len(splits) - 3 :]
	return lastTwoLines[0]
}
