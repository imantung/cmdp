package cmdp

type Command interface {
	Name() string
	Description() string
	Execute(args []string) (output interface{}, err error)
}

type Execution func(args []string) (output interface{}, err error)

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
func (c *command) Execute(args []string) (output interface{}, err error) {
	return c.exec(args)
}
