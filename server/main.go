package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	goql "github.com/ssarangi/goql/goql"
)

type statementType int32

const (
	statementINSERT statementType = iota
	statementSELECT
	statementCREATETABLE
)

const columnUsernameSize = 32
const columnEmailSize = 255

type row struct {
	id       uint32
	username [columnUsernameSize]byte
	email    [columnEmailSize]byte
}

type statement struct {
	stmtType    statementType
	rowToInsert row
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

func prepareSQLCommand(command string) (*statement, goql.PrepareSQLCommandResult) {
	statement := new(statement)
	sqlCommandResult := goql.PrepareSuccess

	// Create statements
	if strings.HasPrefix(command, "create table") == true {
		statement.stmtType = statementCREATETABLE
	} else if strings.HasPrefix(command, "insert") == true {
		statement.stmtType = statementINSERT
		fmt.Sscanf(command, "insert %s %s %s", statement.rowToInsert.id, statement.rowToInsert.username, statement.rowToInsert.email)
	} else if strings.HasPrefix(command, "select") == true {
		statement.stmtType = statementSELECT
	} else {
		sqlCommandResult = goql.PrepareUnrecognizedStatement
	}

	return statement, sqlCommandResult
}

func executeSQLCommand(sqlStatement *statement) {
	switch statement.stmtType {
	case statementINSERT:
		fmt.Println("Insert Statement")
		break
	case statementSELECT:
		fmt.Println("Select Statement")
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
			continue
		}

		// If it's not a metacommand then prepare the command to be fed into the VM
		sqlStatement, sqlCommandResult := prepareSQLCommand(command)
		if sqlCommandResult == goql.PrepareUnrecognizedStatement {
			fmt.Println("Unrecognized SQL command provided")
		}

		executeSQLCommand(sqlStatement)
	}
}
