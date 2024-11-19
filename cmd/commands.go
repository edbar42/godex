package cmd

// Describes a cli command for the app
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Load commands map to the main app singleton
func (a *App) loadCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    a.cliHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex app",
			callback:    a.cliExit,
		},
	}
}
