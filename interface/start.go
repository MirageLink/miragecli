package _interface

import (
	"bufio"
	"fmt"
	"main/lang"
	"main/utils"
	"os"
	"strings"
)

var CommandsObj *Commands

func Start() {
	// TODO: show version on startup
	fmt.Println(lang.GetMessage("welcome", utils.Getfromconfig("lang")))
	CommandsObj := NewCommands()
	initializeCommands(CommandsObj)
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print(lang.GetMessage("PS1", utils.Getfromconfig("lang")))
		scanner.Scan()
		input := scanner.Text()
		inputArray := strings.Fields(input)
		command := inputArray[0]
		args := inputArray[1:]
		if strings.TrimSpace(command) == "exit" {
			os.Exit(0)
		}
		if strings.TrimSpace(command) == "" {
			continue
		}
		fmt.Println(CommandsObj.ExecuteCommand(command, args))
	}
	os.Exit(0)
}
