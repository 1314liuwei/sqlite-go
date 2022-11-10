package main

import (
	"1314liuwei/sqlite.go/consts"
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
		fmt.Printf("db>")
		input, err = inputReader.ReadString('\n')
		if err != nil {
			log.Fatalf("input err: %s", err)
		}
		input = strings.TrimSpace(input)

		if input == consts.Exit {
			return
		} else {
			fmt.Printf("Unrecognized command '%s'. \n", input)
		}
	}
}
