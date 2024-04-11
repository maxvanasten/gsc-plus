package lexer

import (
	"strconv"
)

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

	for input[index].Identifier != end_token_identifier && index+1 < len(input) {
		tokens = append(tokens, input[index])
		index += 1
	}

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
			} else {
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
		Content:    l.Buffer,
	})
	l.Buffer = ""
}

func (l *Lexer) GetNextToken() Token {
	currentCharacter := string(l.Input[l.Index])
	l.Index += 1
	token := Token{}

	if !l.BufferIsString {
		// Character is not a whitespace
		if currentCharacter == " " || currentCharacter == "\t" || currentCharacter == "\n" {
			l.AddBufferAs("Identifier")
		} else if currentCharacter == ";" {
			l.AddBufferAs("Identifier")
			token.Identifier = "Terminator"
			token.Content = ";"
		} else if currentCharacter == "=" {
			l.AddBufferAs("Identifier")
			token.Identifier = "Equals"
			token.Content = "="
		} else if currentCharacter == "{" {
			l.AddBufferAs("Identifier")
			token.Identifier = "LCurly"
			token.Content = "{"
		} else if currentCharacter == "}" {
			l.AddBufferAs("Identifier")
			token.Identifier = "RCurly"
			token.Content = "}"
		} else if currentCharacter == "(" {
			l.AddBufferAs("Identifier")
			token.Identifier = "LParen"
			token.Content = "("
		} else if currentCharacter == ")" {
			l.AddBufferAs("Identifier")
			token.Identifier = "RParen"
			token.Content = ")"
		} else if currentCharacter == "," {
			l.AddBufferAs("Identifier")
			token.Identifier = "Comma"
			token.Content = ","
		} else if currentCharacter == "+" || currentCharacter == "-" || currentCharacter == "*" {
			l.AddBufferAs("Identifier")
			token.Identifier = "Operator"
			token.Content = currentCharacter
		} else if currentCharacter == "\"" {
			l.BufferIsString = true
			l.Buffer += currentCharacter
		} else {
			// Character is other character (aka part of identifier)
			l.Buffer += currentCharacter
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
