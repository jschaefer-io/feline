package main

import "errors"

type Item interface{}

type ItemGroup interface {
	add(item Item)
}

type Group struct {
	Items []Item
}

func (group *Group) add(item Item) {
	group.Items = append(group.Items, item)
}

type Scope struct {
	Group
}

type SubParse struct {
	getInstance func() ItemGroup
	endChar     string
	tokenType   TokenType
}

var subParses = [...]SubParse{
	{
		getInstance: func() ItemGroup { return &Scope{} },
		endChar:     "}",
		tokenType:   CurlyBrackets,
	},
	{
		getInstance: func() ItemGroup { return &Group{} },
		endChar:     ")",
		tokenType:   Parenthesis,
	},
}

func testSubParse(token Token, tokenType TokenType, compare string) (bool, bool) {
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

func buildSubParses(token *Token, tokens *Queue) (ItemGroup, error, bool) {
	var con bool
	var subScope bool

	for _, subParse := range subParses {
		con, subScope = testSubParse(*token, subParse.tokenType, subParse.endChar)
		if !con {
			if subScope {
				item, err := NewParse(tokens, subParse.getInstance())
				return item, err, true
			}
			return nil, nil, true
		}
	}
	return nil, nil, false
}

func NewParse(tokens *Queue, item ItemGroup) (ItemGroup, error) {
	for tokens.len() > 0 {
		tokenInterface, _ := tokens.pop()
		token := tokenInterface.(Token)

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
