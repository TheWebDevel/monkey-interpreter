package ast

import (
	"bytes"

	"github.com/thewebdevel/monkey-interpreter/token"
)

// Node is an interface
// Every node in our AST has to implement the Node interface
// meaning it has to provide a TokenLiteral() method that
// returns the literal value that it associates with
type Node interface {
	// TokenLiteral() will be used only for debugging and testing
	TokenLiteral() string
	// allows us to print AST nodes for debugging and to compare other AST Nodes
	String() string
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

// ExpressionStatement has two fields, the token field, which every node has
// and the Expression field which holds the expression.
// ast.ExpressionStatement fulfills the ast.Statement Interface, which means
// we can add that to the Statements slice of ast.Program
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// statementNode satisfy the Statement Interface
func (es *ExpressionStatement) statementNode() {}

// TokenLiteral satisfiy the Node Interface
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// IntegerLiteral fulfills the ast.Expression interface.
// Here, value is an int64 type and not a string
// When we build an *ast.IntegerLiteral we have to convert the string in
// /*ast.IntegerLiteral.Token.Literal to an int64 ("5" to 5)
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral satisfiy the Node Interface
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

func (il *IntegerLiteral) String() string { return il.Token.Literal }

// String method creates a buffer and writes the return value of each
// statement's String() method to it. It then returns a buffer of a string.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// With these String() in place, we can now just call String() on *ast.Program
// and get our whole program back as string. This makes the *ast.Program
// easy to test and debug
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String() + " ")
	out.WriteString("=")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

func (i *Identifier) String() string { return i.Value }
