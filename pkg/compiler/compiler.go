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
		}
	}

	return output
}
