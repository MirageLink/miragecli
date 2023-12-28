package _interface

import (
	"main/lang"
	"main/utils"
	"strings"
)

type Command struct {
	Name     string
	Function func(args []string) string
}

type Commands struct {
	commandMap map[string]func(args []string) string
}

func NewCommands() *Commands {
	return &Commands{
		commandMap: make(map[string]func(args []string) string),
	}
}

func (c *Commands) AddCommand(name string, function func(args []string) string) {
	c.commandMap[strings.ToLower(name)] = function
}

func (c *Commands) ExecuteCommand(name string, args []string) string {

	if commandFunc, exists := c.commandMap[strings.ToLower(name)]; exists {
		return commandFunc(args)
	}
	return lang.GetMessage("unknown_command", utils.Getfromconfig("lang")) + ": " + name
}
