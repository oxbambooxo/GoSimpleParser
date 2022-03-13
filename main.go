package main

import "fmt"

func main() {

	tokens, err := TokenizeString("a=1; b=a+(2*3)")
	fmt.Println(tokens)
	fmt.Println(err)
}
