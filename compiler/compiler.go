package compiler

import (
	"1314liuwei/sqlite.go/consts"
	"1314liuwei/sqlite.go/core"
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

func (t *Compiler) PrepareStatement(cmd string) (core.UserTableStatement, error) {
	if strings.HasPrefix(strings.ToLower(cmd), "insert") {
		var row core.UserTableRow

		_, err := fmt.Sscanf(cmd, "insert %d %s %s", &row.ID, &row.Username, &row.Email)
		if err != nil {
			return core.UserTableStatement{}, err
		}

		return core.UserTableStatement{Type: consts.StInsert, Row: row}, nil
	}

	if strings.HasPrefix(strings.ToLower(cmd), "select") {
		return core.UserTableStatement{Type: consts.StSelect}, nil
	}

	return core.UserTableStatement{}, errors.New("unrecognized prepare state")
}
