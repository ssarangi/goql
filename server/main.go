package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ssarangi/goql/goql"
)

func handleMetaCommand(command string) goql.MetaCommandResult {
	if strings.Compare(command, ".exit") == 0 {
		os.Exit(0)
	}

	return goql.MetaCommandUnrecognized
}

func checkMetaCommandSuccess(command string, metaCommandResult goql.MetaCommandResult) {
	if metaCommandResult == goql.MetaCommandUnrecognized {
		fmt.Printf("Unrecognized MetaCommand encountered: %s\n", command)
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
			result := handleMetaCommand(command)
			checkMetaCommandSuccess(command, result)
		}
		// If it's not a metacommand then prepare the command to be fed into the VM
	}
}
