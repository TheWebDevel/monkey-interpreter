package parser

import (
	"testing"

	"github.com/thewebdevel/monkey-interpreter/ast"
	"github.com/thewebdevel/monkey-interpreter/lexer"
)

// We are checking as many fields of the AST nodes as possible
// and make sure nothing is missing
func TestLetStatements(t *testing.T) {
	// We provide monkey source code as an input
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383; `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("prgram.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatements(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatements(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral is not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.Letstatement.got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value is not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
