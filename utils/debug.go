package utils

import (
	"fmt"
	"main/lang"
	"os"
)

func Debug(Args []string) {
	switch Args[0] {
	case "sendmessage":
		fmt.Println(lang.GetMessage(Args[1], Args[2]))
		os.Exit(0)
	}
}
