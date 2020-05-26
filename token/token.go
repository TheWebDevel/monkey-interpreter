package token

// A custom type of type string
type TokenType string

// A struct that has
type Token struct {
	Type    TokenType
	Literal string
}

// Different token types in the monkey programming language
// We defined it as constants
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers as Literals
	IDENT = "INDENT" // add, foobar, x, y...
	INT   = "INT"    // 12345

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
