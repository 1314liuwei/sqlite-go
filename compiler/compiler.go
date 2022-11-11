package compiler

import (
	"fmt"
	"os"
	"strings"
)

type MetaCommandState int

const (
	McsSuccess MetaCommandState = iota
	McsUnrecognizedCommand
)

type PrepareState int

const (
	PsSuccess PrepareState = iota
	PsUnrecognizedState
)

type StatementType int

const (
	StUnknow StatementType = iota
	StInsert
	StSelect
)

func DoMetaCommand(cmd string) MetaCommandState {
	switch cmd {
	case ".exit":
		os.Exit(0)
		return McsSuccess
	default:
		return McsUnrecognizedCommand
	}
}

func PrepareStatement(cmd string) (StatementType, PrepareState) {
	stMap := map[string]StatementType{
		"insert": StInsert,
		"select": StSelect,
	}

	for s, statementType := range stMap {
		if strings.HasPrefix(strings.ToLower(cmd), s) {
			return statementType, PsSuccess
		}
	}
	return StUnknow, PsUnrecognizedState

}

func ExecuteStatement(state StatementType) {
	switch state {
	case StInsert:
		fmt.Println("This is where we would do an insert.")
	case StSelect:
		fmt.Println("This is where we would do a select.")
	}
}
