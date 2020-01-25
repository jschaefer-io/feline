package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	file, fileErr := prepareFile("./testfile.feline")
	if fileErr != nil {
		panic(fileErr)
	}
	scope, parseError := parseFile(file)
	if parseError != nil {
		panic(parseError)
	}

	bytes, _ := json.MarshalIndent(scope, "", "\t")
	fmt.Println(string(bytes))
}

func parseFile(file File) (Scope, error) {
	queue := tokenListToParseQueue(file.GetTokenList())
	itemGroup, err := NewParse(&queue, &Scope{})
	if err != nil {
		return Scope{}, err
	}
	if queue.len() > 0 {
		return Scope{}, errors.New("a subparsing group has not been closed properly")
	}
	var scope *Scope = itemGroup.(*Scope)
	return *scope, nil

}

func tokenListToParseQueue(tokens []Token) Queue {
	queue := Queue{}
	for _, token := range tokens {
		queue.push(token)
	}
	queue.push(Token{
		Type:  CurlyBrackets,
		Value: "}",
	})
	return queue
}

func prepareFile(path string) (File, error) {
	file := NewFile(path)
	err := file.Prepare()
	if err != nil {
		return file, err
	}
	return file, nil
}
