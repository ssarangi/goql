package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	goql "github.com/ssarangi/goql/goql"
	goql_parser "github.com/ssarangi/goql/goql/parser"
)

var ctx *goql.Context

func showIntro() {
	logo :=
		`
                GGGGGGGGGGGGG                      QQQQQQQQQ      LLLLLLLLLLL
             GGG::::::::::::G                    QQ:::::::::QQ    L:::::::::L
           GG:::::::::::::::G                  QQ:::::::::::::QQ  L:::::::::L
          G:::::GGGGGGGG::::G                 Q:::::::QQQ:::::::Q LL:::::::LL
         G:::::G       GGGGGG   ooooooooooo   Q::::::O   Q::::::Q   L:::::L
        G:::::G               oo:::::::::::oo Q:::::O     Q:::::Q   L:::::L
        G:::::G              o:::::::::::::::oQ:::::O     Q:::::Q   L:::::L
        G:::::G    GGGGGGGGGGo:::::ooooo:::::oQ:::::O     Q:::::Q   L:::::L
        G:::::G    G::::::::Go::::o     o::::oQ:::::O     Q:::::Q   L:::::L
        G:::::G    GGGGG::::Go::::o     o::::oQ:::::O     Q:::::Q   L:::::L
        G:::::G        G::::Go::::o     o::::oQ:::::O  QQQQ:::::Q   L:::::L
         G:::::G       G::::Go::::o     o::::oQ::::::O Q::::::::Q   L:::::L         LLLLLL
          G:::::GGGGGGGG::::Go:::::ooooo:::::oQ:::::::QQ::::::::Q LL:::::::LLLLLLLLL:::::L
           GG:::::::::::::::Go:::::::::::::::o QQ::::::::::::::Q  L::::::::::::::::::::::L
             GGG::::::GGG:::G oo:::::::::::oo    QQ:::::::::::Q   L::::::::::::::::::::::L
                GGGGGG   GGGG   ooooooooooo        QQQQQQQQ::::QQ LLLLLLLLLLLLLLLLLLLLLLLL
                                                           Q:::::Q
                                                            QQQQQQ
        `

	introString := "Welcome to GoQL: A SQL based database developed in GoLang."
	fmt.Println(logo)
	fmt.Println("")
	fmt.Println(introString)
}

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
	showIntro()
	ctx = new(goql.Context)
	reader := bufio.NewReader(os.Stdin)
	for true {
		printPrompt()
		command := readInput(reader)

		// Handle the metacommands first before handling the SQL commands.
		if strings.Compare(string(command[0]), ".") == 0 {
			result := handleMetaCommand(command)
			checkMetaCommandSuccess(command, result)
			continue
		}

		// If it's not a metacommand then prepare the command to be fed into the VM
		stmt, error := goql_parser.NewParser(command).Parse()
		if error != nil {
			fmt.Println(error)
		}
		ctx.Execute(stmt)
	}
}
