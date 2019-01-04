package cmdp

import "fmt"

type MultiCommand interface {
	Command
	Register(cmd Command) (err error)
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

func (m *multiCommand) Execute(args []string) (output string, err error) {
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
	return cmd.Execute(args[1:])
}
