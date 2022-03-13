package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(">> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if input == "" {
			continue
		}
		if input == "exit;" {
			break
		}
		// tokens, err := TokenizeString(input)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println("tokens", tokens)

		node, err := ParserString(input)
		if err != nil {
			panic(err)
		}
		PrintAST(node)
	}
}
