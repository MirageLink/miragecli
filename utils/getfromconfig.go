package utils

import "strings"

func Getfromconfig(name string) string {
	lines := strings.Split(ConfigFileContent, "\n")
	for _, b := range lines {
		keyword := strings.Split(b, " ")
		if keyword[0] == name {
			return keyword[1]
		}
	}
	if name == "lang" {
		return "uwu"
	}
	return ""
}
