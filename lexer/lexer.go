package lexer

import (
	"github.com/thewebdevel/monkey-interpreter/token"
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

// New function will use read char so that our *Lexer is in a fully working state
// before anyone calls NextToken()
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

// NextToken function will look at the current character under examination (l.ch)
// and return a token depending on which character it is
func (l *Lexer) NextToken() token.Token {
	// Declare a variable tok of type token.Token
	var tok token.Token

	// Skip whiteslace
	l.skipWhitespace()

	// Based on the Character under examination
	// return the appropriate chracter
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
	default:
		// A default case to check for identifier whenever l.ch is not recognized
		// Recognize if the current char is a letter. If so, it needs to read the
		// rest of the identifier/keyword until it encounters a one letter character
		// After reading the identifier/keyword we need to find out if its an
		// identifier or keyword, so we can correct token.TokenType
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			// Check if the identifier is a keyword and assign the type appropriately
			tok.Type = token.LookupIndent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		// If we don't know how to handle current chracter then we declare it as token.ILLEGAL
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

// This function helps us with initializing the tokens for NextToken()
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Reads an identifier and advances our lexer position until it encounters a non-letter-character
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// Check if the current char is a letter
func isLetter(ch byte) bool {
	// We check if the current char falls in the alphabet range of both upper and lower case
	// We also consider '_' as letter and allow it in identifier and keyword
	// This means that we can use variables like foo_baar. Other program like
	// ruby allows ? and ! so this is the place where we can sneak it in
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// readNumber is exactly same as the readIdentifier except it's use of isDigit
// instead of isLetter
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// Check if the current character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Skip whitespace when lexing as it does not have any meaning other than seperating tokens
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
