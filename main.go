package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// token.Value.(float64)

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

/*
func tokenlist(...list int)[]int{
	return list
}

func startswith tokenlist  endswith tokenlist
	-> erst startswith
	-> dann sub-parse process
	-> endswith


# Expression besteht aus:
- Term
	Token (Value)
	group/Expression
	functionCall
- Operator (Type)

# variable set besteht aus:
let (keyword)
xyz (keyword)
type (keyword)
= Operator
group/Expression
STOP (token)

# function define besteht aus:
keyword
group
	keyword
	type keyword
	komma delimiter
scope

# function call besteht aus:
keyword
group
	keyword
	type keyword
	komma delimiter
*/

func prepareFile(path string) (File, error) {
	file := NewFile(path)
	err := file.Prepare()
	if err != nil {
		return file, err
	}
	return file, nil
}
