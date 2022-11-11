package main

import (
	"1314liuwei/sqlite.go/compiler"
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
			case compiler.McsSuccess:
				continue
			case compiler.McsUnrecognizedCommand:
				fmt.Println("Unrecognized command: ", input)
			}
		}

		st, ps := compiler.PrepareStatement(input)
		switch ps {
		case compiler.PsSuccess:
		case compiler.PsUnrecognizedState:
			fmt.Printf("Unrecognized keyword at start of '%s'.", input)
		}

		compiler.ExecuteStatement(st)
	}
}
