package main

import (
	"1314liuwei/sqlite.go/compiler"
	"1314liuwei/sqlite.go/consts"
	"1314liuwei/sqlite.go/core"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, Sqlite-go!")

	var (
		err         error
		input       string
		inputReader = bufio.NewReader(os.Stdin)
	)

	db, err := core.Open("./db.gdb")
	if err != nil {
		return
	}
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

		if input[0] == '.' {
			switch compiler.DoMetaCommand(input) {
			case consts.McsSuccess:
				continue
			case consts.McsUnrecognizedCommand:
				fmt.Println("Unrecognized command: ", input)
			}
		}

		st, err := compiler.PrepareStatement(input)
		if err != nil {
			log.Fatal(err)
		}

		err = db.ExecuteStatement(st)
		if err != nil {
			log.Println(err)
		}
	}
}
