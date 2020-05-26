package lexer

import (
	"monkey-interpreter/token"
)

// Lexer is of type struct. `readPosition` always points to the next char of our input
// Which helps us to peek while `positon` points to the character
// in the input coressponds to the `ch` byte
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// This function gives us the next character and advances our postion in
// the input string.
func (l *Lexer) readChar() {
	// It first checks whether we reached our end of input, if that's the case,
	// it sets l.ch to 0 which is the ASCII code for "NUL" and signifies
	// either we haven't read anything end or end of file for us.
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// But if we haven't it sets ch with the next character by accessing
		// l.input[l.readPositon]
		l.ch = l.input[l.readPosition]
	}

	// l.position is updated to the just used l.redPosition
	// l.readPositon is incremented by 1
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
