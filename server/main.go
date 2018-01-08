package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func doMetaCommand(command string) {

}

func readInput(reader *bufio.Reader) string {
	command, _ := reader.ReadString('\n')
	command = strings.TrimRight(command, "\n")
	return command
}

func printPrompt() {
	fmt.Print("goql> ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		printPrompt()
		command := readInput(reader)
		if strings.Compare(command, ".exit") == 0 {
			fmt.Println("Exiting GoQL...")
			os.Exit(0)
		} else {
			fmt.Printf("Unrecognized command '%s'.\n", command)
		}
	}
}
