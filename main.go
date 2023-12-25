package main

import (
	_interface "main/interface"
	"main/utils"
	"os"
)

var Args = os.Args[1:]

func main() {
	if len(Args) == 0 {
		_interface.Start()
	}
	utils.ParseArgs(Args)
}
