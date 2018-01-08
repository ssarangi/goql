package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handleMetaCommand(command string) {
	if strings.Compare(command, ".exit") == 0 {
		os.Exit(0)
	} else {
		fmt.Printf("Unrecognized command '%s'.\n", command)
	}
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

		// Handle the metacommands first before handling the SQL commands.
		if strings.Compare(string(command[0]), ".") == 0 {
			handleMetaCommand(command)
		}
	}
}
