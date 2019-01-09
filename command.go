package cmdp

type Command interface {
	Name() string
	Description() string
	Execute(ctx interface{}, args []string) (output string, err error)
}

type Execution func(ctx interface{}, args []string) (output string, err error)

type command struct {
	name        string
	description string
	exec        Execution
}

// NewCommand
func NewCommand(name, description string, exec Execution) Command {
	return &command{name: name, description: description, exec: exec}
}

// Name
func (c *command) Name() string {
	return c.name
}

// Description
func (c *command) Description() string {
	return c.description
}

// Execute
func (c *command) Execute(ctx interface{}, args []string) (output string, err error) {
	return c.exec(ctx, args)
}
