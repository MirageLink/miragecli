package main

import (
	_interface "main/interface"
	"main/scripts"
	"main/utils"
	"os"
)

var Args = os.Args[1:]

func main() {
	utils.LoadConfig()
	scripts.InitScripts()
	if len(Args) == 0 {
		_interface.Start()
	}
	utils.ParseArgs(Args)
}
