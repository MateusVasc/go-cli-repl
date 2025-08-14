package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MateusVasc/go-cli-repl/registry"
	"github.com/MateusVasc/go-cli-repl/utils"
)

func main() {
	for {
		fmt.Print("GoCLI > ")
		scn := bufio.NewScanner(os.Stdin)

		if scn.Scan() {
			input := scn.Text()
			val := utils.CleanInput(input)[0]

			command, ok := registry.Commands[val]
			if !ok {
				fmt.Printf("Unknown command %s\n", val)
			} else {
				err := command.Callback()

				if err != nil {
					fmt.Printf("Error while running command %s: %v\n", command.Name, err)
				}
			}

		}
	}
}
