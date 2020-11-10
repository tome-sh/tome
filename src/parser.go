package tome

import (
	"os"
	"strings"
)

var batchSizeKb int64 = 512 * 1000

// Parser interface

type Parser interface {
	Parse() string
}

// Zsh parser

type ZshParser struct {
	path string
}

func (p ZshParser) Parse() string {
	line := readSecondToLastLine(p.path)
	splits := strings.Split(line, ";")
	return splits[len(splits) - 1]
}

func NewZshParser(p string) Parser {
	return ZshParser{path: p}
}

// Helpers

func readSecondToLastLine(filePath string) string {
	file, err := os.Open(filePath)
	Check(err)
	defer file.Close()

	buf := make([]byte, batchSizeKb)
	stat, err := os.Stat(filePath)
	Check(err)
	start := stat.Size() - batchSizeKb
	_, err = file.ReadAt(buf, start)
	Check(err)

	splits := strings.Split(string(buf), "\n")

	lastTwoLines := splits[len(splits) - 3 :]
	return lastTwoLines[0]
}
