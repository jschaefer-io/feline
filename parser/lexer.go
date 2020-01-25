package parser

import (
	"errors"
)

type Lexer struct {
	Position int
	Command  string
	Tokens   []Token
	Length   int
}

func NewLexer(command *string) Lexer {
	return Lexer{
		Position: -1,
		Command:  *command,
		Length:   len(*command),
	}
}

func (lexer *Lexer) GetCharAt(position int) (rune, error) {
	if position < 0 || position >= lexer.Length {
		return ' ', errors.New("undefined char index")
	}
	return rune(lexer.Command[position]), nil
}

func (lexer *Lexer) GetTokenAt(position int) (Token, error) {
	if position < 0 || position >= len(lexer.Tokens) {
		return Token{}, errors.New("undefined token index")
	}
	return lexer.Tokens[position], nil
}

func (lexer *Lexer) handlePosition() error {
	token, offset, err := NextToken(&lexer.Command, &lexer.Position, &lexer.Length)
	if token.Type != Whitespace {
		lexer.Tokens = append(lexer.Tokens, token)
	}
	lexer.Position += offset
	return err
}

func (lexer *Lexer) increment() bool {
	lexer.Position++
	return lexer.Position < lexer.Length
}

func (lexer *Lexer) Tokenize() error {
	for lexer.increment() {
		err := lexer.handlePosition()
		if err != nil {
			return err
		}
	}
	return nil
}
