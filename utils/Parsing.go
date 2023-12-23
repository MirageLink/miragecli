package utils

func ParseArgs(Args []string) {
	if Args[0] == "debug" {
		Debug(Args[1:])
	}
}
