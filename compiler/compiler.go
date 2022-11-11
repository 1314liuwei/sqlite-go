package compiler

import (
	"1314liuwei/sqlite.go/table"
	"errors"
	"fmt"
	"os"
	"strings"
)

type MetaCommandState int

const (
	McsSuccess MetaCommandState = iota
	McsUnrecognizedCommand
)

type StatementType int

const (
	StUnknown StatementType = iota
	StInsert
	StSelect
)

type UserTableStatement struct {
	Type StatementType
	Row  table.UserTableRow
}

func DoMetaCommand(cmd string) MetaCommandState {
	switch cmd {
	case ".exit":
		os.Exit(0)
		return McsSuccess
	default:
		return McsUnrecognizedCommand
	}
}

func PrepareStatement(cmd string) (UserTableStatement, error) {
	if strings.HasPrefix(strings.ToLower(cmd), "insert") {
		var row table.UserTableRow

		_, err := fmt.Sscanf(cmd, "insert %d %s %s", &row.ID, &row.Username, &row.Email)
		if err != nil {
			return UserTableStatement{}, err
		}

		return UserTableStatement{Type: StInsert, Row: row}, nil
	}

	if strings.HasPrefix(strings.ToLower(cmd), "select") {
		return UserTableStatement{Type: StSelect}, nil
	}

	return UserTableStatement{}, errors.New("unrecognized prepare state")
}

func ExecuteStatement(state UserTableStatement) error {
	switch state.Type {
	case StInsert:
		fmt.Println("exec insert!")
		err := table.Table().ExecuteInsert(state.Row)
		if err != nil {
			return err
		}
	case StSelect:
		fmt.Println("exec select!")
		err := table.Table().ExecuteSelect()
		if err != nil {
			return err
		}
	}
	return nil
}
