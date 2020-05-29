package ast

import "github.com/thewebdevel/monkey-interpreter/token"

// Node is an interface
// Every node in our AST has to implement the Node interface
// meaning it has to provide a TokenLiteral() method that
// returns the literal value that it associates with
type Node interface {
	// TokenLiteral() will be used only for debugging and testing
	TokenLiteral() string
}

// The AST we are going to construct consists solely of Nodes
// that are connected to eachother as it's a tree. Some of these
// nodes implement the Statement and some of them implement an
// Expression interface
//
// These interfaces only contain dummy methods called statementNode()
// and expressionNode respectively. They are not strictly necessary
// but help us by guiding the Go compiler and possibly causing it to throw
// erros when we use a Statement where an Expression should've been used
// and vice versa

// Statement is an interface
type Statement interface {
	Node
	statementNode()
}

// Expression is an interface
type Expression interface {
	Node
	expressionNode()
}

// Program node is going to be the root node of every AST our parser produces
// Every valid monkey program is a series of statements. These statements are
// contained in the Program.Statements, whihc is a slice of AST nodes that implements
// he statement interface
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// LetStatement have the fields we need:
// Name to identify the identifier of the binding and Value for the expression
// that produces the value.
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// statementNode satisfy the Statement Interface
func (ls *LetStatement) statementNode() {}

// TokenLiteral satisfiy the Node Interface
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier struct type (which implement the Expression Interfoace)
// is to hold the identifier of the binding. Identifiers in other parts
// of a Monkey program do produce values, eg: let x = valueProduingIdentifier;.
// And to keep the number of different nodes small, we'll use identifier here
// to represent the name in a variable binding and later reuse it, to represent
// an identifier as part of or as complete expression
// Ex: x in let x = 5;
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// statementNode satisfy the Statement Interface
func (i *Identifier) expressionNode() {}

// TokenLiteral satisfiy the Node Interface
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ReturnStatement of type struct that has a token (return) and a ReturnValue
// which is an expression
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

// statementNode satisfy the Statement Interface
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral satisfiy the Node Interface
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
