package main

import (
	"errors"
	"fmt"
	"github.com/jschaefer-io/feline/ast"
	"github.com/jschaefer-io/feline/ast/literals"
	"github.com/jschaefer-io/feline/ast/operators"
	"github.com/jschaefer-io/feline/data_types"
	"github.com/jschaefer-io/feline/files"
	"github.com/jschaefer-io/feline/lexer"
	"github.com/jschaefer-io/feline/parser"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		panic("No input file specified")
	}
	file, fileErr := prepareFile(os.Args[1])
	if fileErr != nil {
		panic(fileErr)
	}
	scope, parseError := parseFile(file)
	if parseError != nil {
		panic(parseError)
	}

	//bytes, _ := json.MarshalIndent(scope, "", "\t")
	//fmt.Println(string(bytes))

	buildAst(&scope)
}

func buildAst(scope *parser.Scope) {
	a := literals.NewBooleanLiteral(true)
	b := literals.NewCharLiteral('C')
	op := operators.Addition{}
	exp := ast.BinaryExpression{&op, &a, &b}

	res, err := exp.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func parseFile(file files.File) (parser.Scope, error) {
	queue := tokenListToParseQueue(file.GetTokenList())
	itemGroup, err := parser.Parse(&queue, &parser.Scope{})
	if err != nil {
		return parser.Scope{}, err
	}
	if queue.Len() > 0 {
		return parser.Scope{}, errors.New("a subparsing group has not been closed properly")
	}
	var scope = itemGroup.(*parser.Scope)
	return *scope, nil

}

func tokenListToParseQueue(tokens []lexer.Token) data_types.Queue {
	queue := data_types.Queue{}
	for _, token := range tokens {
		queue.Push(token)
	}
	queue.Push(lexer.Token{
		Type:  lexer.CurlyBrackets,
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
