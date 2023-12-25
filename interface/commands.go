package _interface

func funny(args []string) string {
	return "lmao get good"
}

func initializeCommands(commands *Commands) {
	commands.AddCommand("im", funny)
}
