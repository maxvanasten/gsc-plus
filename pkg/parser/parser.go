package parser

import (
	// "fmt"

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

		if token.Identifier == "VariableKeyword" && index+3 < len(tokens) {
			// NOTE: Variable declaration
			// let VARNAME = VARVALUE;
			// let VARNAME = VARVALUE1 +/-/* VARVALUE2 ...;
			var_name := tokens[index+1].Content
			// fmt.Println("New variable declaration", tokens[index+1].Content)
			var_tokens, new_index := lexer.GetTokensBetween(index+3, "Terminator", tokens)
			index = new_index
			parsed_var_tokens := ParseTokens(var_tokens)
			// for _, var_token := range var_tokens {
			// 	fmt.Println("VAR_TOKEN:", var_token)
			// }

			// fmt.Println("PARSED_VAR_TOKENS:")
			// for _, parsed_var_token := range parsed_var_tokens {
			// 	fmt.Println(parsed_var_token)
			// }

			nodes = append(nodes, Node{
				Identifier: "Variable_Declaration",
				Content:    var_name,
				Children:   parsed_var_tokens,
			})
		} else if token.Identifier == "FunctionKeyword" {
			function_name := tokens[index+1].Content

			arguments, new_index := lexer.GetTokensBetween(index+3, "RParen", tokens)
			index = new_index
			parsed_arguments := ParseTokens(arguments)

			scope, new_index := lexer.GetTokensBetween(index+2, "RCurly", tokens)
			index = new_index
			parsed_scope := ParseTokens(scope)

			nodes = append(nodes, Node{
				Identifier: "Function_Declaration",
				Content:    function_name,
				Children: []Node{
					{
						Identifier: "Arguments",
						Children:   parsed_arguments,
					},
					{
						Identifier: "Scope",
						Children:   parsed_scope,
					},
				},
			})
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
