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

// Define keywords
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIndent Checks the keyword table to see if the given identifier is a keyword
// If it is, return the keyword's TokenType constant
// Else, just get back the token.IDENT which is the token type for all user
// defined identifiers
func LookupIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
