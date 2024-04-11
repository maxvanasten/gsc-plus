package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/maxvanasten/gsc++/pkg/compiler"
	"github.com/maxvanasten/gsc++/pkg/debug"
	"github.com/maxvanasten/gsc++/pkg/lexer"
	"github.com/maxvanasten/gsc++/pkg/parser"
)

var d = debug.Debugger{Name: "Main", Level: "debug"}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./gscpp path/to/input.gpp")
		os.Exit(2)
	}

	file_path := os.Args[1]

	file_name_splits := strings.Split(file_path, "/")
	file_name := file_name_splits[len(file_name_splits)-1]
	file_name = strings.TrimSuffix(file_name, ".gpp")

	output_file_path := "./output/gsc/" + file_name + ".gsc"

	fmt.Println("Attempting to compile:", file_name)

	input_bytes, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Println("Error while reading file:", err)
		os.Exit(1)
	}

	lexer := lexer.Lexer{
		Input: input_bytes,
	}

	tokens := lexer.Tokenize()

	fmt.Println("Found", len(tokens), "tokens")

	nodes := parser.ParseTokens(tokens)

	output := compiler.Compile(nodes, 0)

	fmt.Printf("Writing compiled code to ./output/gsc/%v.gsc\n", file_name)
	os.WriteFile(output_file_path, []byte(output), 0666)

	ast_json, err := json.Marshal(nodes)
	if err != nil {
		fmt.Println("There was an error converting the AST to JSON:", err)
		os.Exit(3)
	}

	fmt.Printf("Writing AST to ./output/ast/%v.json\n", file_name)
	os.WriteFile("./output/ast/"+file_name+".json", ast_json, 0666)
}
