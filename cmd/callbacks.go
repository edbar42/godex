package cmd

func (a *App) cliHelp() error {
	for _, command := range a.CommandsMap {
		command.toString()
	}
	return nil
}

func (a *App) cliExit() error {
	return nil
}
