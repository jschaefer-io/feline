package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jschaefer-io/feline/data_types"
	"github.com/jschaefer-io/feline/files"
	"github.com/jschaefer-io/feline/parser"
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

func parseFile(file files.File) (parser.Scope, error) {
	queue := tokenListToParseQueue(file.GetTokenList())
	itemGroup, err := parser.NewParse(&queue, &parser.Scope{})
	if err != nil {
		return parser.Scope{}, err
	}
	if queue.Len() > 0 {
		return parser.Scope{}, errors.New("a subparsing group has not been closed properly")
	}
	var scope *parser.Scope = itemGroup.(*parser.Scope)
	return *scope, nil

}

func tokenListToParseQueue(tokens []parser.Token) data_types.Queue {
	queue := data_types.Queue{}
	for _, token := range tokens {
		queue.Push(token)
	}
	queue.Push(parser.Token{
		Type:  parser.CurlyBrackets,
		Value: "}",
	})
	return queue
}

func prepareFile(path string) (files.File, error) {
	file := files.New(path)
	err := file.Prepare()
	if err != nil {
		return file, err
	}
	return file, nil
}
