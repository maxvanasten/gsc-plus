package lexer

import (
	"strconv"
	"strings"

	"github.com/maxvanasten/gsc++/pkg/debug"
)

var d = debug.Debugger{Name: "Lexer", Level: "debug"}

type Lexer struct {
	Input          []byte
	Index          int
	Buffer         string
	BufferIsString bool
	Tokens         []Token
}

type Token struct {
	Identifier string
	Content    string
}

func GetTokensBetween(index int, end_token_identifier string, input []Token) ([]Token, int) {
	tokens := []Token{}

	// while identifiers are not equal and more tokens exist
	for input[index].Identifier != end_token_identifier && index+1 < len(input) {
		// add token to list
		tokens = append(tokens, input[index])
		index += 1
	}

	tokens = append(tokens, input[index])
	return tokens, index
}

func (l *Lexer) Tokenize() []Token {
	// Get all tokens in input
	for l.Index+1 < len(l.Input) {
		l.Tokens = append(l.Tokens, l.GetNextToken())
	}
	// Add last buffer as identifier if not empty
	if l.Buffer != "" && l.Buffer != " " && l.Buffer != "\t" && l.Buffer != "\n" {
		l.AddBufferAs("Identifier")
	}

	// Sanitize tokens by removing all empty tokens
	new_tokens := []Token{}
	for _, token := range l.Tokens {
		if token.Identifier != "" {
			// Identifier is not empty
			if token.Identifier == "Identifier" {
				// Token is an identifier
				if token.Content != "" && token.Content != " " {
					// Content is not empty

					// If identifier can be parsed as number, change token identifier
					_, err := strconv.ParseFloat(token.Content, 32)
					if err == nil {
						token.Identifier = "Number"
					}

					// Parse keywords
					switch token.Content {
					case "let":
						token.Identifier = "VariableKeyword"
					case "fn":
						token.Identifier = "FunctionKeyword"
					case "return":
						token.Identifier = "ReturnKeyword"
					}

					new_tokens = append(new_tokens, token)
				}
			} else if token.Identifier != "Newline" && token.Identifier != "Whitespace" && token.Identifier != "Tab" {
				new_tokens = append(new_tokens, token)
			}

		}
	}

	l.Tokens = new_tokens
	return l.Tokens
}

// Add current buffer as token with identifier, clear current buffer
func (l *Lexer) AddBufferAs(identifier string) {
	l.Tokens = append(l.Tokens, Token{
		Identifier: identifier,
		Content:    strings.TrimSpace(l.Buffer),
	})
	l.Buffer = ""
}

type SpecialCharacter struct {
	Identifier string
	Content    string
}

var SPECIALCHARACTERS = []SpecialCharacter{
	{
		Identifier: "Whitespace",
		Content:    " ",
	},
	{
		Identifier: "Tab",
		Content:    "\t",
	},
	{
		Identifier: "Newline",
		Content:    "\n",
	},
	{
		Identifier: "Terminator",
		Content:    ";",
	},
	{
		Identifier: "Equals",
		Content:    "=",
	},
	{
		Identifier: "LCurly",
		Content:    "{",
	},
	{
		Identifier: "RCurly",
		Content:    "}",
	},
	{
		Identifier: "LParen",
		Content:    "(",
	},
	{
		Identifier: "RParen",
		Content:    ")",
	},
	{
		Identifier: "Comma",
		Content:    ",",
	},
	{
		Identifier: "PlusOperator",
		Content:    "+",
	},
	{
		Identifier: "MinusOperator",
		Content:    "-",
	},
	{
		Identifier: "MultOperator",
		Content:    "*",
	},
}

func (l *Lexer) GetNextToken() Token {
	currentCharacter := string(l.Input[l.Index])
	l.Index += 1
	token := Token{}

	if !l.BufferIsString {
		is_special := false
		for _, character := range SPECIALCHARACTERS {
			if currentCharacter == character.Content {
				is_special = true
				l.AddBufferAs("Identifier")
				token.Identifier = character.Identifier
				token.Content = character.Content
			}
		}

		if currentCharacter == "\"" {
			l.BufferIsString = true
			l.Buffer += currentCharacter
		} else {
			// Character is other character (aka part of identifier)
			if !is_special {
				l.Buffer += currentCharacter
			}
		}
	} else {
		l.Buffer += currentCharacter
		if currentCharacter == "\"" {
			l.AddBufferAs("String")
			l.BufferIsString = false
		}
	}

	return token
}
