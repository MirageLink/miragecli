package main

import (
	"main/utils"
	"os"
)

var Args = os.Args[1:]

func main() {
	utils.ParseArgs(Args)
	//if !utils.IsServerUp() {
	//	guh, guh2 := lang.GetMessage("gun", "uwu")
	//	fmt.Println(guh, guh2)
	//}
}
