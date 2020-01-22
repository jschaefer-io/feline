package main

import (
	"errors"
)

type Lexer struct {
	position int
	command  string
	tokens   []Token
	length   int
}

func NewLexer(command *string) Lexer {
	return Lexer{
		position: -1,
		command:  *command,
		length:   len(*command),
	}
}

func (lexer *Lexer) GetCharAt(position int) (rune, error) {
	if position < 0 || position >= lexer.length {
		return ' ', errors.New("undefined char index")
	}
	return rune(lexer.command[position]), nil
}

func (lexer *Lexer) GetTokenAt(position int) (Token, error) {
	if position < 0 || position >= len(lexer.tokens) {
		return Token{}, errors.New("undefined token index")
	}
	return lexer.tokens[position], nil
}


func (lexer *Lexer) handlePosition() error {
	token, offset := NextToken(&lexer.command, &lexer.position, &lexer.length)
	lexer.tokens = append(lexer.tokens, token)
	lexer.position += offset
	return nil
}

func (lexer *Lexer) increment() bool {
	lexer.position++
	return lexer.position < lexer.length
}

func (lexer *Lexer) Tokenize() {
	for lexer.increment() {
		_ = lexer.handlePosition()
	}
}
