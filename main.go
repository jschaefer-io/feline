package main

import (
	"fmt"
)

var command string = "test(name string, num int): int{print_r(num + \" Hello \" + name);}"

func main() {
	commandLexer := NewLexer(&command)
	commandLexer.Tokenize()
	for _, token := range commandLexer.tokens {
		fmt.Printf("Type: %d | Value: %s\n", token.Type, token.Value)
	}
}
