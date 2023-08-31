package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Repl() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex >")
		sc.Scan()
		err := sc.Err()
		if err != nil {
			log.Fatal(err)
		}

		word := sc.Text()
		if len(word) == 0 {
			continue
		}

		word = strings.ToLower(word)
		word = strings.Fields(word)[0]

		command, ok := Commands()[word]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Invalid command")
			continue
		}
	}
}
