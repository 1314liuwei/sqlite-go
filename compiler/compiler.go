package compiler

import (
	"1314liuwei/sqlite.go/consts"
	"1314liuwei/sqlite.go/database"
	"errors"
	"fmt"
	"os"
	"strings"
)

func DoMetaCommand(cmd string) consts.MetaCommandState {
	switch cmd {
	case ".exit":
		os.Exit(0)
		return consts.McsSuccess
	default:
		return consts.McsUnrecognizedCommand
	}
}

func PrepareStatement(cmd string) (database.UserTableStatement, error) {
	if strings.HasPrefix(strings.ToLower(cmd), "insert") {
		var row database.UserTableRow

		_, err := fmt.Sscanf(cmd, "insert %d %s %s", &row.ID, &row.Username, &row.Email)
		if err != nil {
			return database.UserTableStatement{}, err
		}

		return database.UserTableStatement{Type: consts.StInsert, Row: row}, nil
	}

	if strings.HasPrefix(strings.ToLower(cmd), "select") {
		return database.UserTableStatement{Type: consts.StSelect}, nil
	}

	return database.UserTableStatement{}, errors.New("unrecognized prepare state")
}
