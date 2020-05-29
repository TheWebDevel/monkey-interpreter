package parser

import (
	"fmt"

	"github.com/thewebdevel/monkey-interpreter/ast"
	"github.com/thewebdevel/monkey-interpreter/lexer"
	"github.com/thewebdevel/monkey-interpreter/token"
)

// Parser has 3 fields: l, curToken and peekToken.
// l is a pointer to an instance of the lexer on which
// we repeatedly call the NextToken() to get the next token from the input
// curToken and peekToken are exactly like the two "pointers" or lexer has:
// position and readPosition. But instead of pointing to the character,
// they point to the next TOKEN.
type Parser struct {
	l *lexer.Lexer
	// An error field which is a slice of strings
	errors []string

	// We need to look at the curToken, which is the current token under examination
	// to decide what to do next, we also need peekToken for the decision if
	// curToken doesn't give us enough information.
	// Ex: Consider the single line containing `5;` then curToken is a token.INT
	// and we need the peekToken to decide whether we are at the end of the line
	// of if we are at just start of the arithmatic expression
	curToken  token.Token
	peekToken token.Token
}

// New function returns an intial Parser that has a lexer, errors, curToken and the peekToken
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two token so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken is a helper method that advances curToken and peekToken
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram consturcts the AST
func (p *Parser) ParseProgram() *ast.Program {
	// Construct the root node of the AST which is an *ast.Program
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Iterate every token in the input until it encounters an EOF token
	// It does this by repeatedly calling p.NextToken(), which advances
	// both p.curToken and p.peekToken. In every iteration, it calls
	// parseStatement, which parses a statement as the name suggests.
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		// If parseStatement returned something other than nil, a ast.Statement
		// it's returned value is added to Statements slice of the AST root node
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	// When nothin is left to parse, the *ast.Program root note is returned
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// parseLetStatement constructs an *ast.LetStatement with the token it's
// currently sitting (a token.LET token) and then advances the tokens
// while making assertions about the next token with calls to expectPeek.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// First it expects a token.IDENT token, which it then uses to construct
	// an *ast.Identifier node
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Then it excects an equal sign
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// Then it jumps over the expression following the equal sign until it
	// encounters a semi-colon
	// TODO: we're skipping the expression until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement constructs an ast,ReturnStatement, with the current
// token it's sitting on as Token. It then brings the parsr in place for the
// expression that comes next by calling nextToken()
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// Then it jumps over the expression until it encounters a semi-colon
	// TODO: we're skipping the expression until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Check if the curToken type is equal to the type in parameter
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// Check if the peekToken type is equal to the type in parameter
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek enforces the correctness of the order of tokens
// by checking the type of the next token.
func (p *Parser) expectPeek(t token.TokenType) bool {
	// It checks the type of peekToken, only if it's correct, it advances the
	// tokens by calling nextToken()
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

// Errors will check if the parser has encountered any errors
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError is used to add an error to errors when the type of peekToken
// does not match the expectation.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
