package parser

import (
	"fmt"

	"github.com/maxvanasten/gsc++/pkg/lexer"
)

type Parser struct {
	Tokens []lexer.Token
	Index  int
	Nodes  []Node
}

type Node struct {
	Identifier string
	Content    string
	Children   []Node
}

func ParseTokens(tokens []lexer.Token) []Node {
	nodes := []Node{}
	index := 0
	for index < len(tokens) {
		token := tokens[index]
		fmt.Println("Starting at", index, "(", token.Identifier, ") (", token.Content, ")")

		if token.Identifier == "VariableKeyword" && index+3 < len(tokens) {
			// NOTE: Variable declaration
			// let VARNAME = VARVALUE;
			// let VARNAME = VARVALUE1 +/-/* VARVALUE2 ...;
			fmt.Println("New variable declaration", tokens[index+1].Content)
			var_tokens, new_index := lexer.GetTokensBetween(index+3, "Terminator", tokens)

			for _, var_token := range var_tokens {
				fmt.Println("VAR_TOKEN:", var_token)
			}

			parsed_var_tokens := ParseTokens(var_tokens)
			fmt.Println("PARSED_VAR_TOKENS:")
			for _, parsed_var_token := range parsed_var_tokens {
				fmt.Println(parsed_var_token)
			}

			nodes = append(nodes, Node{
				Identifier: "Variable_Declaration",
				Content:    tokens[index+1].Content,
				Children:   parsed_var_tokens,
			})

			fmt.Println("New index: ", index+new_index)
			index = new_index
		} else if token.Identifier == "LParen" || token.Identifier == "RParen" || token.Identifier == "Identifier" || token.Identifier == "String" || token.Identifier == "Number" || token.Identifier == "PlusOperator" || token.Identifier == "MinusOperator" || token.Identifier == "MultOperator" {
			nodes = append(nodes, Node{
				Identifier: token.Identifier,
				Content:    token.Content,
			})
			index += 1
		} else {
			index += 1
		}
	}

	return nodes
}
