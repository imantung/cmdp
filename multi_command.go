package cmdp

import "fmt"

type MultiCommand interface {
	Command
	Register(cmd Command) (err error)
	RegisterCommand(name, description string, exec Execution) (cmd Command, err error)
	GetCommand(name string) (cmd Command, isExist bool)
}

type multiCommand struct {
	command
	cmdMap map[string]Command
}

func NewMultiCommand(name, description string) MultiCommand {
	return &multiCommand{
		command: command{name: name, description: description},
		cmdMap:  make(map[string]Command),
	}
}

func (m *multiCommand) Register(cmd Command) (err error) {
	name := cmd.Name()
	if _, ok := m.cmdMap[name]; ok {
		err = fmt.Errorf("%s is already exist", name)
		return
	}
	m.cmdMap[name] = cmd
	return
}

func (m *multiCommand) RegisterCommand(name, description string, exec Execution) (cmd Command, err error) {
	cmd = NewCommand(name, description, exec)
	err = m.Register(cmd)
	return
}

func (m *multiCommand) GetCommand(name string) (cmd Command, isExist bool) {
	cmd, isExist = m.cmdMap[name]
	return
}

func (m *multiCommand) Execute(ctx interface{}, args []string) (output string, err error) {
	if len(args) < 1 {
		err = fmt.Errorf("missing arguments")
		return
	}

	cmdName := args[0]
	cmd, ok := m.cmdMap[cmdName]
	if !ok {
		err = fmt.Errorf("Command '%s' is not found", cmdName)
		return
	}
	return cmd.Execute(ctx, args[1:])
}
