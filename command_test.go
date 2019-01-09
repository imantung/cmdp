package cmdp_test

import (
	"fmt"
	"testing"

	"github.com/imantung/cmdp"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	cmd := cmdp.NewCommand("some-name", "some-description", func(ctx interface{}, args []string) (string, error) {
		return "some-output", fmt.Errorf("some-error")
	})

	assert.Equal(t, cmd.Name(), "some-name")
	assert.Equal(t, cmd.Description(), "some-description")

	out, err := cmd.Execute(nil, nil)
	assert.Equal(t, out, "some-output")
	assert.EqualError(t, err, "some-error")
}
