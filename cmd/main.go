package main

import (
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

		_, err = db.Execute(input)
		if err != nil {
			log.Println(err)
		}
	}
}
