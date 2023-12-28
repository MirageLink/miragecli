package utils

import (
	"fmt"
	"main/lang"
	"main/scripts"
)

func ParseArgs(Args []string) {
	switch Args[0] {
	case "debug":
		Debug(Args[1:])
	case "run":
		runScript(Args[1])
	}
}

func runScript(name string) {
	if containsString(scripts.Scripts, name) {
		scripts.ExecuteScript(name)
	} else {
		fmt.Println(lang.GetMessage("script_not_found", Getfromconfig("lang")))
		return
	}
}

func containsString(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}
