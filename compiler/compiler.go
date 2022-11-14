package compiler

import (
	"1314liuwei/sqlite.go/backend"
	"1314liuwei/sqlite.go/consts"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Compiler struct{}

func (t *Compiler) GetStdinInput() string {
	var (
		inputReader = bufio.NewReader(os.Stdin)
		input       string
		err         error
	)

	for {
		fmt.Printf("db> ")
		input, err = inputReader.ReadString('\n')
		if err != nil {
			log.Fatalf("input err: %s", err)
		}
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}
	}
}

func (t *Compiler) DoMetaCommand(cmd string) error {
	switch cmd {
	case ".exit":
		os.Exit(0)
		return nil
	default:
		return errors.New("unrecognized meta command")
	}
}

func (t *Compiler) PrepareStatement(cmd string) (backend.UserTableStatement, error) {
	if strings.HasPrefix(strings.ToLower(cmd), "insert") {
		var row backend.UserTableRow

		_, err := fmt.Sscanf(cmd, "insert %d %s %s", &row.ID, &row.Username, &row.Email)
		if err != nil {
			return backend.UserTableStatement{}, err
		}

		return backend.UserTableStatement{Type: consts.StInsert, Row: row}, nil
	}

	if strings.HasPrefix(strings.ToLower(cmd), "select") {
		return backend.UserTableStatement{Type: consts.StSelect}, nil
	}

	return backend.UserTableStatement{}, errors.New("unrecognized prepare state")
}
