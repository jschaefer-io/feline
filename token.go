package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TokenType uint8

const (
	Stop           TokenType = 0
	Operator       TokenType = 1
	String         TokenType = 2
	Char           TokenType = 3
	Number         TokenType = 4
	Keyword        TokenType = 5
	Whitespace     TokenType = 6
	Parenthesis    TokenType = 7
	CurlyBrackets  TokenType = 8
	SquareBrackets TokenType = 9
	Delimiter      TokenType = 10
)

type Token struct {
	Type  TokenType
	Value interface{}
}

func determineType(char rune) TokenType {
	charString := string(char)
	if strings.Index(" \t\n\r", charString) >= 0 {
		return Whitespace
	}
	if strings.Index("+-*/%", charString) >= 0 {
		return Operator
	}
	if strings.Index("()", charString) >= 0 {
		return Parenthesis
	}
	if strings.Index("{}", charString) >= 0 {
		return CurlyBrackets
	}
	if strings.Index("[]", charString) >= 0 {
		return SquareBrackets
	}
	if char == '"' {
		return String
	}
	if char == ',' {
		return Delimiter
	}
	if char == '\'' {
		return Char
	}
	if char == ';' {
		return Stop
	}
	if strings.Index("0123456789.", charString) >= 0 {
		return Number
	}
	return Keyword
}

func getTokenFromType(tokenType TokenType) (Token, error) {
	token := Token{Type: tokenType}
	return token, nil
}

func collectUntilMatch(needMatch string, count int, input *string, position *int, length *int) (string, int) {
	value := ""
	matchCounter := 0
	var pos int
	var currentChar string
	for pos = *position; pos < *length; pos++ {
		currentChar = string((*input)[pos])
		value += currentChar
		if strings.Index(needMatch, currentChar) >= 0 {
			matchCounter++
			if matchCounter == count {
				break
			}
		}
	}
	return value, pos - *position
}

func collectUntilNoMatch(needMatch string, input *string, position *int, length *int) (string, int) {
	value := ""
	var pos int
	var currentChar string
	for pos = *position; pos < *length; pos++ {
		currentChar = string((*input)[pos])
		if strings.Index(needMatch, currentChar) < 0 {
			pos--
			break
		}
		value += currentChar
	}
	return value, pos - *position
}

func collectUntilChange(tokenType TokenType, input *string, position *int, length *int) (string, int) {
	value := ""
	var pos int
	var currentChar rune
	var currentType TokenType
	for pos = *position; pos < *length; pos++ {
		currentChar = rune((*input)[pos])
		currentType = determineType(currentChar)
		if currentType != tokenType {
			pos--
			break
		}
		value += string(currentChar)
	}
	return value, pos - *position
}

func fillToken(token *Token, input *string, position *int, length *int) (int, error) {
	var value string
	offset := 0
	switch token.Type {
	case Char:
		value, offset = collectUntilMatch("'", 2, input, position, length)
		value = strings.Trim(value, "'")
		length := len(value)
		if length != 1 {
			return offset, errors.New(fmt.Sprintf("can not identify '%s' as type char", value))
		}
		token.Value = rune(value[0])
		break
	case String:
		value, offset = collectUntilMatch("\"", 2, input, position, length)
		token.Value = strings.Trim(value, "\"")
		break
	case Number:
		value, offset = collectUntilChange(token.Type, input, position, length)
		parsedValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return offset, errors.New(fmt.Sprintf("can not identify \"%s\" as type number", value))
		}
		token.Value = parsedValue
	default:
		token.Value, offset = collectUntilChange(token.Type, input, position, length)
		break
	}
	return offset, nil
}

func NextToken(input *string, position *int, length *int) (Token, int, error) {
	currentChar := rune((*input)[*position])
	tokenType := determineType(currentChar)
	token, _ := getTokenFromType(tokenType)
	offset, err := fillToken(&token, input, position, length)
	return token, offset, err
}
