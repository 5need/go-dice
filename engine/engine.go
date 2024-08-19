package engine

import (
	"fmt"
	"os"

	"go-dice/engine/lexer"
	"go-dice/engine/parser"
)

func main() {
	// os.Args is a slice of strings that contains the command-line arguments
	args := os.Args

	// Subsequent elements are the actual command-line arguments
	if len(args) > 1 {
		for _, arg := range args[1:] {
			fmt.Printf("String to parse: %s\n", arg)

			myLexer := lexer.NewLexer(arg)
			myParser := parser.NewParser(*myLexer)
			myParser.Parse()
		}
	} else {
		fmt.Println("No arguments provided.")
	}
}

func RollDice(input string) ([]int, error) {
	myLexer := lexer.NewLexer(input)
	myParser := parser.NewParser(*myLexer)
	result, err := myParser.Parse()
	if err != nil {
		return []int{}, err
	}
	return result, nil
}
