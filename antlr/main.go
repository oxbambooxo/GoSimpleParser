package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// value, err := EvalString(" ")
	// fmt.Println(value, err)
	// return

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

		// 调试 eval
		value, err := EvalString(input)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	}
}
