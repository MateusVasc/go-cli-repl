package registry

type cliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config) error
}

type Config struct {
	Next     string
	Previous string
}

var Commands map[string]cliCommand

func init() {
	Commands = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    func(c *Config) error { return commandExit(c) },
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    func(c *Config) error { return commandHelp(c) },
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world",
			Callback:    func(c *Config) error { return commandMap(c) },
		},
	}
}
