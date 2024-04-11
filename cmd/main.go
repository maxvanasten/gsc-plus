package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/maxvanasten/gsc++/pkg/lexer"
)

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

	fmt.Println("FILENAME:", file_name)

	input_bytes, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Println("Error while reading file:", err)
		os.Exit(1)
	}

	lexer := lexer.Lexer{
		Input: input_bytes,
	}

	tokens := lexer.Tokenize()

	fmt.Println("Found", len(tokens), "tokens:")
	for _, token := range tokens {
		fmt.Println("[", token.Identifier, "]: >", token.Content, "<")
	}
}
