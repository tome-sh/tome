package tome

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

var singleLineCommandString = "e3246abc-29a6-11eb-8af2-c3d824a00352;1605292765;Oskar Jung;tag;fg"
var singleLineCommandSerialized = ": e3246abc-29a6-11eb-8af2-c3d824a00352;1605292765;Oskar Jung" +
	";tag;fg"
var singleLineCommand = Command{
	Id:      uuid.MustParse("e3246abc-29a6-11eb-8af2-c3d824a00352"),
	Time:    time.Unix(1605292765, 0),
	Author:  "Oskar Jung",
	Tags:    []string{"tag"},
	Command: "fg",
}
var multiLineCommandSerialized = ": 20f6e9aa-29a7-11eb-924f-5f3b38e417e4;1605290038;Oskar Jung;" +
	"multiline;echo \"asd\\\nfgh\""
var multiLineCommandString = "20f6e9aa-29a7-11eb-924f-5f3b38e417e4;1605290038;Oskar Jung;" +
	"multiline;echo \"asd\\\\fgh\""
var multiLineCommand = Command{
	Id:      uuid.MustParse("20f6e9aa-29a7-11eb-924f-5f3b38e417e4"),
	Time:    time.Unix(1605290038, 0),
	Author:  "Oskar Jung",
	Tags:    []string{"multiline"},
	Command: "echo \"asd\\\nfgh\"",
}

func TestCommand_String(t *testing.T) {
	table := []struct {
		testName string
		input    Command
		expected string
	}{
		{"return command", singleLineCommand, singleLineCommandString},
		{"escape multiline command", multiLineCommand, multiLineCommandString},
	}

	for _, tt := range table {
		t.Run(
			"should"+tt.testName, func(t *testing.T) {
				result := tt.input.String()
				if result != tt.expected {
					t.Errorf("'%s' does not equal '%s'", result, tt.expected)
				}
			},
		)
	}
}

func TestParseCommands(t *testing.T) {

	invalidCommand := ": 1605292765;Oskar Jung;fg"
	invalidUUIDCommand := ": uuid;1605292765;Oskar Jung;tag;fg"
	invalidTimestampCommand := ": 3766fa42-29d7-11eb-ac1d-3f49909e858b;160529276b;Oskar Jung;tag;fg"

	table := []struct {
		testName string
		input    string
		expected []Command
	}{
		{"parse empty file without error", "", nil},
		{"parse command", singleLineCommandSerialized, []Command{singleLineCommand}},
		{"ignore invalid line", invalidCommand, nil},
		{"parse multiline command", multiLineCommandSerialized, []Command{multiLineCommand}},
		{
			"parse multiple commands",
			strings.Join(
				[]string{
					multiLineCommandSerialized,
					invalidCommand,
					singleLineCommandSerialized,
				}, "\n",
			),
			[]Command{multiLineCommand, singleLineCommand},
		},
		{
			"parse only commands that start with newline",
			singleLineCommandSerialized + " " + singleLineCommandSerialized,
			[]Command{
				{
					Id:      singleLineCommand.Id,
					Time:    time.Unix(1605292765, 0),
					Author:  "Oskar Jung",
					Tags:    []string{"tag"},
					Command: "fg " + singleLineCommandSerialized,
				},
			},
		},
		{
			"parse commands containing a semicolon",
			singleLineCommandSerialized + "; test",
			[]Command{
				{
					Id:      singleLineCommand.Id,
					Time:    singleLineCommand.Time,
					Author:  singleLineCommand.Author,
					Tags:    singleLineCommand.Tags,
					Command: singleLineCommand.Command + "; test",
				},
			},
		},
		{
			"skip commands with uuid parsing errors",
			invalidUUIDCommand,
			nil,
		},
		{
			"skip commands with time parsing errors",
			invalidTimestampCommand,
			nil,
		},
	}

	for _, tt := range table {
		t.Run(
			"should "+tt.testName, func(t *testing.T) {
				result, err := Deserialize(strings.NewReader(tt.input))

				if err != nil {
					t.Errorf("Failed to parse '%s'. Error: %s", result, err.Error())
				}

				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("'%s' did not match expected result '%s'", result, tt.expected)
				}
			},
		)
	}
}
