package lexer

type TokenType string

const (
	TokenEOF           TokenType = "EOF"
	TokenIllegal       TokenType = "ILLEGAL"
	TokenIdent         TokenType = "IDENT"
	TokenNumber        TokenType = "NUMBER"
	TokenEquals        TokenType = "EQUALS"
	TokenPlus          TokenType = "PLUS"
	TokenMinus         TokenType = "MINUS"
	TokenPlural        TokenType = "PLURAL"
	TokenMultiply      TokenType = "MULTIPLY"
	TokenDiceSeparator TokenType = "DICE_SEPARATOR"
	TokenOpenBrackets  TokenType = "OPEN_BRACKETS"
	TokenCloseBrackets TokenType = "CLOSE_BRACKETS"
	TokenComma         TokenType = "COMMA"
	TokenReroll        TokenType = "REROLL"
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar() // initialize by reading the first character
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0 // Indicates end of input
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

func isWhiteSpace(ch byte) bool {
	return (ch == ' ' || ch == '\t' || ch == '\n')
}

func (lexer *Lexer) skipWhiteSpace() {
	for isWhiteSpace(lexer.ch) {
		lexer.readChar()
	}
}

func (lexer *Lexer) NextToken() Token {
	lexer.skipWhiteSpace()

	var tok Token

	switch lexer.ch {
	case '=':
		tok = newToken(TokenEquals, lexer.ch)
	case '+':
		tok = newToken(TokenPlus, lexer.ch)
	case '-':
		tok = newToken(TokenMinus, lexer.ch)
	case '[':
		tok = newToken(TokenOpenBrackets, lexer.ch)
	case ']':
		tok = newToken(TokenCloseBrackets, lexer.ch)
	case ',':
		tok = newToken(TokenComma, lexer.ch)
	case '*':
		tok = newToken(TokenMultiply, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = TokenEOF
	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			switch tok.Literal {
			case "d":
				tok.Type = TokenDiceSeparator
			case "x":
				tok.Type = TokenMultiply
			case "s":
				tok.Type = TokenPlural
			case "rr":
				tok.Type = TokenReroll
			default:
				tok.Type = TokenIdent
			}
			return tok
		} else if isDigit(lexer.ch) {
			tok.Literal = lexer.readNumber()
			tok.Type = TokenNumber
			return tok
		} else {
			tok = newToken(TokenIllegal, lexer.ch)
		}
	}

	lexer.readChar()
	return tok
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}
