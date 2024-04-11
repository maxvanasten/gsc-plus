package compiler

import (
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
		if node.Identifier == "Variable_Declaration" {
			output += GetTabs(level) + node.Content + " = "
			for _, child := range node.Children {
				output += child.Content + " "
			}
			output = strings.TrimSuffix(output, " ")
			output += ";\n"
		} else if node.Identifier == "Function_Declaration" {
			output += GetTabs(level) + node.Content + "("
			for _, argument := range node.Children[0].Children {
				output += argument.Content + ", "
			}
			output = strings.TrimSuffix(output, ", ")
			output += ") {\n"

			for _, scope_node := range node.Children[1].Children {
				scope_output := Compile([]parser.Node{scope_node}, level+1)
				output += scope_output
			}

			output += "}\n\n"
		}
	}

	return output
}
