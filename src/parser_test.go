package tome

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestZshParser(t *testing.T) {
	simpleFile, _ := filepath.Abs("./test_resources/parser/simple_history")
	stat, _ := os.Stat(simpleFile)
	table := []struct {
		parser Parser
		expected string
	} {
		{NewZshParser(simpleFile), "md src/test_resources/parser"},
		{NewZshParserWithBatchSize(simpleFile, int64(stat.Size() / 2)), "md src/test_resources/parser"},
	}

	for _, tt := range table {
		testname := fmt.Sprintf("Testing fixture: %s", tt.parser)
		t.Run(testname, func(t *testing.T) {
			ans := tt.parser.Parse()
			if ans != tt.expected {
				t.Errorf("got %s, wanted %s", ans, tt.expected)
			}
		})
	}
}
