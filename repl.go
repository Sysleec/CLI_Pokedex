package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Repl(cfg *config) {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex >")
		sc.Scan()
		err := sc.Err()
		if err != nil {
			log.Fatal(err)
		}

		words := sc.Text()
		if len(words) == 0 {
			continue
		}

		wordLow := strings.ToLower(words)
		wordsF := strings.Fields(wordLow)

		args := []string{}
		if len(words) > 1 {
			args = wordsF[1:]
		}

		command, ok := Commands()[wordsF[0]]
		if ok {
			err := command.callback(cfg, args...)
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
