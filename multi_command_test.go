package cmdp_test

import (
	"strings"
	"testing"

	"github.com/imantung/cmdp"
	"github.com/stretchr/testify/assert"
)

func TestMultiCommand_New(t *testing.T) {
	cmd := cmdp.NewMultiCommand("some-name", "some-description")
	assert.Equal(t, cmd.Name(), "some-name")
	assert.Equal(t, cmd.Description(), "some-description")
}

func TestMultiCommand_Register_SameCommandName(t *testing.T) {
	cmd := cmdp.NewMultiCommand("some-name", "some-description")
	err := cmd.Register(cmd)
	assert.Nil(t, err)

	err = cmd.Register(cmd)
	assert.EqualError(t, err, "some-name is already exist")
}

func TestMultiCommand_RegisterCommand(t *testing.T) {
	multi := cmdp.NewMultiCommand("root", "root")
	cmd, err := multi.RegisterCommand(
		"some-name",
		"some-description",
		func(ctx interface{}, args []string) (string, error) {
			return "output", nil
		},
	)
	assert.Nil(t, err)

	got, isExist := multi.GetCommand("some-name")

	assert.Equal(t, isExist, true)
	assert.Equal(t, cmd.Name(), "some-name")
	assert.Equal(t, got, cmd)
}

func TestMultiCommand_Execute_NoArguments(t *testing.T) {
	cmd := cmdp.NewMultiCommand("some-name", "some-description")
	_, err1 := cmd.Execute(nil, nil)
	_, err2 := cmd.Execute(nil, []string{})

	assert.EqualError(t, err1, "missing arguments")
	assert.EqualError(t, err2, "missing arguments")
}

func TestMultiCommand_Execute_CommandNotFound(t *testing.T) {
	cmd := cmdp.NewMultiCommand("some-name", "some-description")
	_, err := cmd.Execute(nil, []string{"some-command"})

	assert.EqualError(t, err, "Command 'some-command' is not found")
}

func TestMultiCommand_Execute(t *testing.T) {
	cmd := cmdp.NewMultiCommand("some-name", "some-description")
	cmd.Register(cmdp.NewCommand(
		"a",
		"description-a",
		func(ctx interface{}, args []string) (string, error) {
			return strings.Join(args, " "), nil
		}))
	cmd.Register(cmdp.NewCommand(
		"b",
		"description-b",
		func(ctx interface{}, args []string) (string, error) {
			return strings.Join(args, "-"), nil
		}))

	output, err := cmd.Execute(nil, []string{"a", "1", "2", "3"})
	assert.Nil(t, err)
	assert.Equal(t, output, "1 2 3")

	output, err = cmd.Execute(nil, []string{"b", "1", "2", "3"})
	assert.Nil(t, err)
	assert.Equal(t, output, "1-2-3")
}
