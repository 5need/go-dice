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

// program ::= modifier { modifier }
func (parser *Parser) Parse() ([]int, error) {
	result := []int{}
	var err error
	for err == nil {
		result, err = parser.modifier(result)
	}

	return result, nil
}

// roll ::= ( dice | array )
func (parser *Parser) roll() ([]int, error) {
	result, err := parser.dice()
	if err == nil {
		return result, nil
	}
	result, err = parser.array()
	if err == nil {
		return result, nil
	}

	return make([]int, 0), errors.New("Expected a roll")
}

// dice ::= number "d" number
func (parser *Parser) dice() ([]int, error) {
	amountOfDice, err := parser.number()
	if err != nil {
		return make([]int, 0), errors.New("Expected a number")
	}
	if parser.currentToken.Type != lexer.TokenDiceSeparator {
		return make([]int, 0), errors.New("Expected a dice separator")
	}
	parser.nextToken()
	faceCount, err := parser.number()
	if err != nil {
		return make([]int, 0), errors.New("Expected a number")
	}

	results := make([]int, 0)

	for i := 0; i < amountOfDice; i++ {
		results = append(results, rand.Intn(faceCount)+1)
	}

	return results, nil
}

// number ::= an int, dummy
func (parser *Parser) number() (int, error) {
	if parser.currentToken.Type != lexer.TokenNumber {
		return 0, errors.New("Expected a number")
	}
	num, err := strconv.Atoi(parser.currentToken.Literal)
	if err != nil {
		return 0, errors.New("Something stupid happened")
	}
	parser.nextToken()

	return num, nil
}

// array ::= "[" { number { "," number } } "]"
func (parser *Parser) array() ([]int, error) {
	results := make([]int, 0)

	if parser.currentToken.Type != lexer.TokenOpenBrackets {
		return []int{}, errors.New("Expected an opening bracket")
	}
	parser.nextToken()

	// in case of empty array
	if parser.currentToken.Type == lexer.TokenCloseBrackets {
		parser.nextToken()
		return []int{}, nil
	}

	num, err := parser.number()
	if err != nil {
		return []int{}, err
	}
	results = append(results, num)

	// { "," number }
	{
		for parser.currentToken.Type == lexer.TokenComma && parser.currentToken.Type != lexer.TokenCloseBrackets {
			if parser.peekToken.Type != lexer.TokenNumber {
				return []int{}, errors.New("Expected a number")
			}
			parser.nextToken() // on top of number now
			currentNum, err := parser.number()
			if err != nil {
				return []int{}, err
			}

			results = append(results, currentNum)
		}
	}

	parser.nextToken()
	return results, nil
}

// modifier ::= ( add_roll | remove_roll | reroll )
func (parser *Parser) modifier(result []int) ([]int, error) {
	newResult, err := parser.add_roll(result)
	if err == nil {
		fmt.Println("add_roll")
		fmt.Println(result)
		fmt.Println(newResult)
		fmt.Println()
		return newResult, nil
	}
	newResult, err = parser.remove_roll(result)
	if err == nil {
		fmt.Println("remove_roll")
		fmt.Println(result)
		fmt.Println(newResult)
		fmt.Println()
		return newResult, nil
	}
	newResult, err = parser.reroll(result)
	if err == nil {
		fmt.Println("reroll")
		fmt.Println(result)
		fmt.Println(newResult)
		fmt.Println()
		return newResult, nil
	}
	newResult, err = parser.save_roll(result)
	if err == nil {
		fmt.Println("save_roll")
		fmt.Println(result)
		fmt.Println(newResult)
		fmt.Println()
		return newResult, nil
	}

	return result, errors.New("Expected a modifier")
}

// add_roll ::= "+" roll
func (parser *Parser) add_roll(result []int) ([]int, error) {
	if parser.currentToken.Type != lexer.TokenPlus {
		return result, errors.New("Expected a +")
	}
	parser.nextToken()
	toAdd, err := parser.roll()
	if err != nil {
		return result, err
	}
	return append(result, toAdd...), nil
}

// remove_roll ::= "-" roll
func (parser *Parser) remove_roll(result []int) ([]int, error) {
	if parser.currentToken.Type != lexer.TokenMinus {
		return result, errors.New("Expected a -")
	}
	parser.nextToken()

	removeByIndex, err := parser.select_range()
	if err == nil {
		newResult := []int{}
		for _, value := range result {
			amountToRemoveLeft := removeByIndex[value-1]
			if amountToRemoveLeft != -1 {
				if amountToRemoveLeft > 0 {
					removeByIndex[value-1]--
				} else {
					newResult = append(newResult, value)
				}
			}
		}
		return newResult, nil
	}

	return result, nil
}

// save_roll ::= "sv" select_range
func (parser *Parser) save_roll(result []int) ([]int, error) {
	if parser.currentToken.Type != lexer.TokenSave {
		return result, errors.New("Expected an sv")
	}
	parser.nextToken()

	saveByIndex, err := parser.select_range()
	fmt.Println(saveByIndex)
	if err == nil {
		newResult := []int{}
		for _, value := range result {
			amountToSaveLeft := saveByIndex[value-1]
			if amountToSaveLeft == -1 {
				newResult = append(newResult, value)
			} else if amountToSaveLeft > 0 {
				newResult = append(newResult, value)
				saveByIndex[value-1]--
			}
		}
		return newResult, nil
	}

	return result, nil
}

// plural ::= number "s"
func (parser *Parser) plural() (int, error) {
	if parser.peekToken.Type != lexer.TokenPlural {
		return 0, errors.New("Expected a plural")
	}
	num, err := parser.number()
	if err != nil {
		return 0, err
	}
	parser.nextToken()
	return num, nil
}

// range_up ::= number "+"
func (parser *Parser) range_up() (int, error) {
	if parser.peekToken.Type != lexer.TokenPlus {
		return 0, errors.New("Expected a +")
	}
	num, err := parser.number()
	if err != nil {
		return 0, err
	}
	parser.nextToken()
	return num, nil
}

// range_down ::= number "-"
func (parser *Parser) range_down() (int, error) {
	if parser.peekToken.Type != lexer.TokenMinus {
		return 0, errors.New("Expected a -")
	}
	num, err := parser.number()
	if err != nil {
		return 0, err
	}
	parser.nextToken()
	return num, nil
}

// select_range ::= ( plural | range_up | range_down | array )
func (parser *Parser) select_range() ([]int, error) {
	selectRange := []int{0, 0, 0, 0, 0, 0}
	// [ 2, 0, 0, 0, 0, 0 ] means select two 1s
	// [ 0, 0, 4, -1, 0, 0 ] means select four 3s and ALL 4s

	num, err := parser.plural()
	if err == nil {
		selectRange[num-1] = -1
		return selectRange, nil
	}
	num, err = parser.range_up()
	if err == nil {
		for index := range selectRange {
			if num <= index+1 {
				selectRange[index] = -1
			}
		}
		return selectRange, nil
	}
	num, err = parser.range_down()
	if err == nil {
		for index := range selectRange {
			if num >= index+1 {
				selectRange[index] = -1
			}
		}
		return selectRange, nil
	}
	nums := []int{}
	nums, err = parser.array()
	if err == nil {
		selectCounts := []int{0, 0, 0, 0, 0, 0}

		for _, value := range nums {
			selectCounts[value-1]++
		}
		return selectCounts, nil
	}

	return selectRange, errors.New("Expected a select_range")
}

// reroll ::= "rr" select_range
func (parser *Parser) reroll(result []int) ([]int, error) {
	if parser.currentToken.Type != lexer.TokenReroll {
		return result, errors.New("Expected a rr")
	}
	parser.nextToken()
	rerollByIndex, err := parser.select_range()
	if err == nil {
		newResult := []int{}
		for _, value := range result {
			amountToRerollLeft := rerollByIndex[value-1]
			newValue := value
			if amountToRerollLeft == -1 {
				newValue = rand.Intn(6) + 1
			} else if amountToRerollLeft > 0 {
				rerollByIndex[value-1]--
				newValue = rand.Intn(6) + 1
			}
			newResult = append(newResult, newValue)
		}

		return newResult, nil
	}

	return result, nil
}
