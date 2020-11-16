package tome

import (
	"testing"
)

var singleLineCommandString = ": 1605292765;Oskar Jung;tag;fg"
var singleLineCommand = NewCommand(1605292765, "Oskar Jung", []string{"tag"}, "fg")
var multiLineCommandString = ": 1605290038;Oskar Jung;multiline;echo \"asd\\\nfgh\""
var multiLineCommand = NewCommand(1605290038, "Oskar Jung", []string{"multiline"},
	"echo \"asd\\\nfgh\"")

func TestCommand_String(t *testing.T) {
	table := []struct {
		testName string
		input    Command
		expected string
	}{
		{"serialize command", singleLineCommand, singleLineCommandString},
		{"serialize multiline command", multiLineCommand, multiLineCommandString},
	}

	for _, tt := range table {
		t.Run("should"+tt.testName, func(t *testing.T) {
			result := tt.input.String()
			if result != tt.expected {
				t.Errorf("'%s' does not equal '%s'", result, tt.expected)
			}
		})
	}
}
