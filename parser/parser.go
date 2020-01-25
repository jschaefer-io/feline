package parser

import (
	"errors"
	"github.com/jschaefer-io/feline/data_types"
	"github.com/jschaefer-io/feline/lexer"
)

type SubParse struct {
	getInstance func() ItemGroup
	endChar     string
	tokenType   lexer.TokenType
}

var subParses = [...]SubParse{
	{
		getInstance: func() ItemGroup { return &Scope{} },
		endChar:     "}",
		tokenType:   lexer.CurlyBrackets,
	},
	{
		getInstance: func() ItemGroup { return &Group{} },
		endChar:     ")",
		tokenType:   lexer.Parenthesis,
	},
}

func testSubParse(token lexer.Token, tokenType lexer.TokenType, compare string) (bool, bool) {
	if token.Type == tokenType {
		if token.Value == compare {
			// SubParse should end here
			return false, false
		} else {
			// SubParse should start here
			return false, true
		}
	}
	// no action required
	return true, false
}

func buildSubParses(token *lexer.Token, tokens *data_types.Queue) (ItemGroup, error, bool) {
	var con bool
	var subScope bool

	for _, subParse := range subParses {
		con, subScope = testSubParse(*token, subParse.tokenType, subParse.endChar)
		if !con {
			if subScope {
				item, err := Parse(tokens, subParse.getInstance())
				return item, err, true
			}
			return nil, nil, true
		}
	}
	return nil, nil, false
}

func Parse(tokens *data_types.Queue, item ItemGroup) (ItemGroup, error) {
	for tokens.Len() > 0 {
		tokenInterface, _ := tokens.Pop()
		token := tokenInterface.(lexer.Token)

		parse, err, useParse := buildSubParses(&token, tokens)
		if err != nil {
			return nil, err
		}
		if useParse {
			if parse != nil {
				item.add(parse)
				continue
			}
			return item, nil
		}
		item.add(token)
	}
	return nil, errors.New("unexpected EOF")
}
