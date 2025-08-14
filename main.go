package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MateusVasc/go-cli-repl/utils"
)

func main() {
	for {
		fmt.Print("GoCLI > ")
		scn := bufio.NewScanner(os.Stdin)

		if scn.Scan() {
			input := scn.Text()
			fmt.Printf("Your command was: %s\n", utils.CleanInput(input)[0])
		}
	}
}
