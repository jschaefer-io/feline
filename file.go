package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type File struct {
	basename string
	path     string
	lines    []line
}

type line struct {
	number int
	tokens []Token
}

func NewFile(filePath string) File {
	absPath, _ := filepath.Abs(filePath)
	basename := path.Base(filePath)
	return File{
		basename: basename,
		path:     absPath,
	}
}

func handleEscapedChars(content *string) string {
	newString := *content
	newString = strings.ReplaceAll(newString, "\\a", "\a")
	newString = strings.ReplaceAll(newString, "\\b", "\b")
	newString = strings.ReplaceAll(newString, "\\f", "\f")
	newString = strings.ReplaceAll(newString, "\\n", "\n")
	newString = strings.ReplaceAll(newString, "\\r", "\r")
	newString = strings.ReplaceAll(newString, "\\t", "\t")
	newString = strings.ReplaceAll(newString, "\\v", "\v")
	newString = strings.ReplaceAll(newString, "\\\\", "\\")
	newString = strings.ReplaceAll(newString, "\\\"", "\"")
	return newString
}

func (file *File) addLine(command *string, number int) error {
	newLine := line{
		number: number,
	}
	escapedCommand := handleEscapedChars(command)
	commandLexer := NewLexer(&escapedCommand)
	err := commandLexer.Tokenize()
	if err != nil {
		errorText := fmt.Sprintf("parse error at line %d in %s\n", newLine.number, file.path)
		errorText += err.Error()
		return errors.New(errorText)
	}
	for _, token := range commandLexer.tokens {
		newLine.tokens = append(newLine.tokens, token)
	}
	if len(newLine.tokens) > 0 {
		file.lines = append(file.lines, newLine)
	}
	return nil
}

func (file *File) GetTokenList() []Token {
	var tokens []Token
	for _, line := range file.lines {
		tokens = append(tokens, line.tokens...)
	}
	return tokens
}

func (file *File) Prepare() error {
	openFile, fileError := os.OpenFile(file.path, os.O_RDONLY, 0755)
	if fileError != nil {
		return fileError
	}
	scanner := bufio.NewScanner(openFile)

	// Read line by line
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()
		parseError := file.addLine(&line, lineNumber)
		if parseError != nil {
			return parseError
		}
	}

	scannerError := scanner.Err()
	if scannerError != nil {
		return scannerError
	}
	return nil
}
