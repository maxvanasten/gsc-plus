package compiler

import (
	"strings"

	"github.com/maxvanasten/gsc++/pkg/debug"
	"github.com/maxvanasten/gsc++/pkg/parser"
)

var d = debug.Debugger{Name: "Compiler", Level: "debug"}

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
		d.Log("debug", "Node: "+node.Identifier+"["+node.Content+"]")
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
			if len(node.Children[0].Children) > 1 {
				output = strings.TrimSuffix(output, ", ), ")
				output += ")"
			}
			output = strings.TrimSuffix(output, ", ")
			output += "\n{\n"

			for _, scope_node := range node.Children[1].Children {
				scope_output := Compile([]parser.Node{scope_node}, level+1)
				output += scope_output
			}

			output += "}\n"
		} else if node.Identifier == "Method_Call" {
			output += GetTabs(level) + node.Children[0].Content + " "
			if node.Content == "true" {
				output += "thread "
			}
			func_call := Compile([]parser.Node{node.Children[1]}, 0)
			func_call = strings.ReplaceAll(func_call, "));\n", ");\n")
			output += func_call
		} else if node.Identifier == "Function_Call" {
			output += GetTabs(level) + node.Content + "("

			for _, argument := range node.Children {
				output += argument.Content
			}
			output = strings.TrimSuffix(output, ", ")

			output += ";\n"
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
