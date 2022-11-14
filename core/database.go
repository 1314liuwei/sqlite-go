package core

import (
	"1314liuwei/sqlite.go/backend"
	"1314liuwei/sqlite.go/compiler"
	"errors"
	"log"
)

type DB interface {
	Execute(input string) ([]interface{}, error)
}

type Database struct {
	compiler *compiler.Compiler
	backend  *backend.Database
}

func Open(name string) (DB, error) {
	open, err := backend.Open(name)
	if err != nil {
		return nil, err
	}
	return &Database{
		compiler: &compiler.Compiler{},
		backend:  open,
	}, nil
}

func (db *Database) Execute(input string) ([]interface{}, error) {
	if input[0] == '.' {
		if err := db.compiler.DoMetaCommand(input); err != nil {
			return nil, errors.New("Unrecognized command: " + input)
		}
	}

	st, err := db.compiler.PrepareStatement(input)
	if err != nil {
		log.Fatal(err)
	}

	err = db.backend.ExecuteStatement(st)
	if err != nil {
		log.Println(err)
	}
	return nil, nil
}
