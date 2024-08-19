package parser

import (
	"errors"
	"fmt"
	"go-dice/engine/lexer"
	"math/rand"
	"strconv"
)

type Parser struct {
	lexer        lexer.Lexer
	currentToken lexer.Token
	peekToken    lexer.Token
}

// TODO: add some sort of result metadata struct to a result like
// average value (average of all numbers)
// platonic average value (average of all numbers if they were all half of the face_count)
// NOTE: only after a save or BS/WS because that way it's either make it or not make it

// type resultMetadata struct {
// 	result_average int
// 	ideal_average  int
// 	dice_amount    int
// 	dice_type      int
// }

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func NewParser(myLexer lexer.Lexer) *Parser {
	parser := &Parser{lexer: myLexer}
	parser.nextToken()
	parser.nextToken()
	return parser
}

// dicethingwhatever ::= ( basic_dice | result ) { result_modifier }
func (parser *Parser) Parse() []int {
	result := []int{}
	if parser.currentToken.Type == lexer.TokenOpenBrackets {
		result = parser.result()
	} else if parser.currentToken.Type == lexer.TokenNumber {
		result = parser.basic_dice()
	} else {
		panic("unexpected token")
	}

	result = parser.result_modifier(result)

	fmt.Printf("result: %+v\n", result)
	return result
	// parser.result_modifier()
}

// basic_dice should turn into an array of rolled values
func (parser *Parser) basic_dice() []int {
	amount := parser.term()
	if parser.currentToken.Type != lexer.TokenDiceSeparator {
		panic("Expected a dice separator :(")
	}
	parser.nextToken()
	face_count := parser.term()

	results := make([]int, 0)
	for i := 0; i < amount; i++ {
		results = append(results, rand.Intn(face_count)+1)
	}
	return results
}

// term ::= unary { "*" unary}
// just multiplication, and unary is always a positive int so just raw-dog it
func (parser *Parser) term() int {
	// unary
	if parser.currentToken.Type != lexer.TokenNumber {
		panic("Expected a number :(")
	}
	accumulator, err := strconv.Atoi(parser.currentToken.Literal)
	if err != nil {
		panic("Expected a number :(")
	}
	parser.nextToken()

	// { "*" unary }
	for parser.currentToken.Type == lexer.TokenMultiply {
		if parser.peekToken.Type != lexer.TokenNumber {
			panic("Expected a number :(")
		}
		parser.nextToken() // on top of number now
		currentNum, err := strconv.Atoi(parser.currentToken.Literal)
		if err != nil {
			panic("Expected a number :(")
		}
		accumulator *= currentNum
		parser.nextToken()
	}
	// fmt.Println(accumulator)
	return accumulator
}

// result ::= '[' int { "," int } ']'
// [1,2,3]
func (parser *Parser) result() []int {
	results := make([]int, 0)

	// [
	if parser.currentToken.Type != lexer.TokenOpenBrackets {
		panic("Expected an opening bracket")
	}
	parser.nextToken()

	// int
	if parser.currentToken.Type != lexer.TokenNumber {
		panic("Expected a number :(")
	}
	num, err := strconv.Atoi(parser.currentToken.Literal)
	if err != nil {
		panic("Expected a number :(")
	}
	results = append(results, num)

	parser.nextToken()
	// { "," int }
	for parser.currentToken.Type == lexer.TokenComma && parser.currentToken.Type != lexer.TokenCloseBrackets {
		if parser.peekToken.Type != lexer.TokenNumber {
			panic("Expected a number :(")
		}
		parser.nextToken() // on top of number now
		currentNum, err := strconv.Atoi(parser.currentToken.Literal)
		if err != nil {
			panic("Expected a number :(")
		}
		results = append(results, currentNum)
		parser.nextToken()
	}

	if parser.currentToken.Type != lexer.TokenCloseBrackets {
		panic("Expected a closing bracket")
	}
	parser.nextToken()

	return results
}

// result_modifier ::= { number "+" }
func (parser *Parser) result_modifier(result []int) []int {
	if parser.currentToken.Type == lexer.TokenEOF {
		return result
	}

	modifiedResult := result

	// var num int
	var arr []int
	var err error

	for err == nil {
		arr, err = parser.save()
		if err == nil {
			fmt.Println("saving", arr)
			fmt.Println("before", modifiedResult)
			tempResult := make([]int, 0)
			for _, roll := range modifiedResult {
				for _, v := range arr {
					if v == roll {
						tempResult = append(tempResult, roll)
						break
					}
				}
			}
			modifiedResult = tempResult
			fmt.Println("after ", modifiedResult)
			fmt.Println()
		}

		arr, err = parser.reroll()
		if err == nil {
			fmt.Println("rerolling", arr)
			fmt.Println("before", modifiedResult)
			tempResult := make([]int, 0)
			for _, roll := range modifiedResult {
				rerollMe := false
				for _, v := range arr {
					if v == roll {
						rerollMe = true
						break
					}
				}
				if rerollMe == true {
					tempResult = append(tempResult, rand.Intn(6)+1)
				} else {
					tempResult = append(tempResult, roll)
				}
			}
			modifiedResult = tempResult
			fmt.Println("after ", modifiedResult)
			fmt.Println()
		}
	}

	return modifiedResult
}

// save ::= range
func (parser *Parser) save() ([]int, error) {
	return parser.number_range()
}

// reroll ::= "rr" range
func (parser *Parser) reroll() ([]int, error) {
	// create a copy of the parser so I don't have to care about messing it up
	parserCopy := *parser

	if parserCopy.currentToken.Type != lexer.TokenReroll {
		return make([]int, 0), errors.New("not a reroll")
	}

	parser.nextToken()
	return parser.number_range()
}

// number_range ::= number { "+" | "-" | "s" }
func (parser *Parser) number_range() ([]int, error) {
	// create a copy of the parser so I don't have to care about messing it up
	parserCopy := *parser

	if parserCopy.currentToken.Type != lexer.TokenNumber {
		return make([]int, 0), errors.New("not a range")
	}

	num, err := strconv.Atoi(parserCopy.currentToken.Literal)
	if err != nil {
		panic("Expected a number :(")
	}

	switch parserCopy.peekToken.Type {
	case lexer.TokenPlus:
		// 4+ = [4,5,6]
		result := make([]int, 0)
		for i := 1; i <= 6; i++ {
			if i >= num {
				result = append(result, i)
			}
		}

		parser.nextToken()
		parser.nextToken()
		return result, nil
	case lexer.TokenMinus:
		// 4- = [1,2,3,4]
		result := make([]int, 0)
		for i := 1; i <= 6; i++ {
			if i <= num {
				result = append(result, i)
			}
		}

		parser.nextToken()
		parser.nextToken()
		return result, nil
	case lexer.TokenPlural:
		// 4s = [4]
		result := make([]int, 0)
		result = append(result, num)

		parser.nextToken()
		parser.nextToken()
		return result, nil
	default:
		// 4 = [4]
		result := make([]int, 0)
		result = append(result, num)

		parser.nextToken()
		return result, nil
	}
}
