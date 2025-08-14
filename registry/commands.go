package registry

import (
	"fmt"
	"net/http"
	"os"
)

func commandExit(_ *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(_ *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, val := range Commands {
		fmt.Printf("%v: %s\n", val.Name, val.Description)
	}

	return nil
}

func commandMap(c *Config) error {
	fPage := "https://pokeapi.co/api/v2/location-area"
	if c.Next == "" {
		res, err := http.Get(fPage)
		if err != nil {
			return err
		}
	} else {
		res, err := http.Get(c.Next)
		if err != nil {
			return err
		}
	}

	return nil
}
