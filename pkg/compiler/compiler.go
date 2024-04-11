package compiler

import (
	"fmt"
	"strings"

	"github.com/maxvanasten/gsc++/pkg/parser"
)

func GetTabs(level int) string {
	tabs := ""

	for len(tabs) < level {
		tabs += "\t"
	}

	return tabs
}

func Compile(nodes []parser.Node, level int) string {
	output := ""

	for _, node := range nodes {
		fmt.Println("Node: ", node.Identifier)
		if node.Identifier == "Variable_Declaration" || node.Identifier == "Variable_Reassignment" {
			output += GetTabs(level) + node.Content + " = "
			for _, child := range node.Children {
				output += child.Content + " "
			}
			output = strings.TrimSuffix(output, " ")
			output += ";\n"
		} else if node.Identifier == "Function_Declaration" {
			output += "\n" + GetTabs(level) + node.Content + "("
			for _, argument := range node.Children[0].Children {
				if argument.Identifier != "Comma" {
					output += argument.Content + ", "
				}
			}
			output = strings.TrimSuffix(output, ", ")
			output += ")\n{\n"

			for _, scope_node := range node.Children[1].Children {
				scope_output := Compile([]parser.Node{scope_node}, level+1)
				output += scope_output
			}

			output += "}\n"
		} else if node.Identifier == "Function_Call" {
			output += GetTabs(level) + node.Content + "("

			for _, argument := range node.Children {
				output += argument.Content
			}
			output = strings.TrimSuffix(output, ", ")

			output += ");\n"
		} else if node.Identifier == "Return_Statement" {
			output += GetTabs(level) + "return "

			for _, node := range node.Children {
				output += node.Content + " "
			}
			output = strings.TrimSuffix(output, " ")

			output += ";\n"
		} else if node.Identifier == "Include_Statement" {
			output += "#include " + node.Content + ";\n"
		}
	}

	return output
}
