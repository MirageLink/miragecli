package utils

import (
	"fmt"
	"main/lang"
	"os"
)

func Debug(Args []string) {
	switch Args[0] {
	case "sendmessage":
		guh1, guh2 := lang.GetMessage(Args[1], Args[2])
		fmt.Println(guh1, guh2)
		os.Exit(0)
	}
}
