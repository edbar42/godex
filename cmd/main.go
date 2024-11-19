// Package cmd contains all executable code for the Pokédex app
package cmd

import "fmt"

// Application singleton to handle dependencies
type App struct {
	CommandsMap map[string]cliCommand
}

func main() {
	var app App
	fmt.Println("Welcome, Pokémon trainer!")
	app = App{
		CommandsMap: app.loadCommands(),
	}
}
