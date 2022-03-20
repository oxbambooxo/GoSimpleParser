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
		// 调试 lexer
		// tokens, err := TokenizeString(input)
		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println("tokens", tokens)
		// }

		// 调试 parser
		// node, err := ParserString(input)
		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	PrintAST(node)
		// }

		// 调试 eval
		value, err := EvalString(input)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	}
}
