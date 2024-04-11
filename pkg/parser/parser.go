package parser

import (
	// "fmt"

	"fmt"

	"github.com/maxvanasten/gsc++/pkg/debug"
	"github.com/maxvanasten/gsc++/pkg/lexer"
)

var d = debug.Debugger{Name: "Parser", Level: "debug"}

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

		d.Log("debug", "Identifier: "+token.Identifier+", Content: "+token.Content)

		if token.Identifier == "VariableKeyword" && index+3 < len(tokens) {
			d.Log("debug", "MATCHING: Variable Declaration")
			// NOTE: Variable declaration
			// let VARNAME = VARVALUE;
			// let VARNAME = VARVALUE1 +/-/* VARVALUE2 ...;
			var_name := tokens[index+1].Content
			// d.Log("New variable declaration", tokens[index+1].Content)
			var_tokens, new_index := lexer.GetTokensBetween(index+3, "Terminator", tokens)
			index = new_index
			parsed_var_tokens := ParseTokens(var_tokens)
			for _, var_token := range var_tokens {
				d.Log("debug", "VAR_TOKEN: "+var_token.Identifier+"["+var_token.Content+"]")
			}

			d.Log("debug", "PARSED_VAR_TOKENS:")
			for _, parsed_var_token := range parsed_var_tokens {
				d.Log("debug", "PARSED_VAR_TOKEN: "+parsed_var_token.Identifier+"["+parsed_var_token.Content+"]")
			}

			nodes = append(nodes, Node{
				Identifier: "Variable_Declaration",
				Content:    var_name,
				Children:   parsed_var_tokens,
			})
		} else if token.Identifier == "FunctionKeyword" {
			d.Log("debug", "MATCHING: Function Declaration")
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
		} else if token.Identifier == "Identifier" {
			if tokens[index+1].Identifier == "LParen" && index+2 < len(tokens) {
				d.Log("debug", "MATCHING: Function Call")
				function_name := token.Content

				arguments, new_index := lexer.GetTokensBetween(index+2, "RParen", tokens)
				index = new_index
				parsed_arguments := ParseTokens(arguments)

				nodes = append(nodes, Node{
					Identifier: "Function_Call",
					Content:    function_name,
					Children:   parsed_arguments,
				})
			} else if token.Content == "#include" && index+1 < len(tokens) {
				d.Log("debug", "MATCHING: Include Statement")
				nodes = append(nodes, Node{
					Identifier: "Include_Statement",
					Content:    tokens[index+1].Content,
				})
				index += 2
			} else if tokens[index+1].Identifier == "Equals" && index+2 < len(tokens) {
				d.Log("debug", "MATCHING: Variable Reassignment")
				ass_tokens, new_index := lexer.GetTokensBetween(index+2, "Terminator", tokens)
				index = new_index
				parsed_ass_tokens := ParseTokens(ass_tokens)

				nodes = append(nodes, Node{
					Identifier: "Variable_Reassignment",
					Content:    token.Content,
					Children:   parsed_ass_tokens,
				})
			} else if tokens[index+1].Identifier == "Greater" && index+2 < len(tokens) {
				d.Log("debug", "MATCHING: Method Call")
				object := token.Content
				parsed_func_call := []Node{}
				thread := "false"

				if tokens[index+2].Identifier == "Greater" {
					func_tokens, new_index := lexer.GetTokensBetween(index+2, "Terminator", tokens)
					index = new_index
					parsed_func_call = ParseTokens(func_tokens)
					thread = "true"
				} else {
					func_tokens, new_index := lexer.GetTokensBetween(index+1, "Terminator", tokens)
					index = new_index
					parsed_func_call = ParseTokens(func_tokens)
				}

				d.Log("debug", "PARSED_FUNC_CALL: ")
				for _, node := range parsed_func_call {
					d.Log("debug", node.Identifier+"["+node.Content+"]")
				}

				node := Node{
					Identifier: "Method_Call",
					Content:    thread,
					Children: []Node{
						{
							Identifier: "Object_Name",
							Content:    object,
						},
					},
				}

				// TODO: Fix function calls without arguments not working correctly.
				// NOTE: Might have something to do with indexing in the parser.
				for _, node := range parsed_func_call {
					fmt.Println("------------------", node)
				}

				node.Children = append(node.Children, parsed_func_call...)

				nodes = append(nodes, node)

			} else {
				nodes = append(nodes, Node{
					Identifier: "Identifier",
					Content:    token.Content,
				})
				index += 1
			}
		} else if token.Identifier == "ReturnKeyword" && index+2 < len(tokens) {
			return_tokens, new_index := lexer.GetTokensBetween(index+1, "Terminator", tokens)
			index = new_index
			parsed_return_tokens := ParseTokens(return_tokens)

			nodes = append(nodes, Node{
				Identifier: "Return_Statement",
				Children:   parsed_return_tokens,
			})
		} else if token.Identifier == "LParen" || token.Identifier == "RParen" || token.Identifier == "Identifier" || token.Identifier == "String" || token.Identifier == "Number" || token.Identifier == "PlusOperator" || token.Identifier == "MinusOperator" || token.Identifier == "MultOperator" || token.Identifier == "Comma" {
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
