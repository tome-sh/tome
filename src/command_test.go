package tome

import (
	"reflect"
	"strings"
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

func TestParseCommands(t *testing.T) {

	singleLineCommandString := ": 1605292765;Oskar Jung;tag;fg"
	singleLineCommand := NewCommand(1605292765, "Oskar Jung", []string{"tag"}, "fg")
	multiLineCommandString := ": 1605290038;Oskar Jung;multiline;echo \"asd\\\nfgh\""
	multiLineCommand := NewCommand(1605290038, "Oskar Jung", []string{"multiline"},
		"echo \"asd\\\nfgh\"")
	invalidCommandString := ": 1605292765;Oskar Jung;fg"

	table := []struct {
		testName string
		input    string
		expected []Command
	}{
		{"parse empty file without error", "", nil},
		{"parse command", singleLineCommandString, []Command{singleLineCommand}},
		{"ignore invalid line", invalidCommandString, nil},
		{"parse multiline command", multiLineCommandString, []Command{multiLineCommand}},
		{"parse multiple commands",
			strings.Join([]string{multiLineCommandString, invalidCommandString, singleLineCommandString}, "\n"),
			[]Command{multiLineCommand, singleLineCommand}},
		{"parse only commands that start with newline",
			singleLineCommandString + " " + singleLineCommandString,
			[]Command{NewCommand(1605292765, "Oskar Jung", []string{"tag"},
				"fg "+singleLineCommandString)},
		},
	}

	for _, tt := range table {
		t.Run("should "+tt.testName, func(t *testing.T) {
			result, err := ParseCommands(strings.NewReader(tt.input))

			if err != nil {
				t.Errorf("Failed to parse '%s'. Error: %s", result, err.Error())
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("'%s' did not match expected result '%s'", result, tt.expected)
			}
		})
	}
}
