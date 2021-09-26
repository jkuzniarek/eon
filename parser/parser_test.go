package parser

import (
	// "fmt"
	"eon/ast"
	"eon/lexer"
	"testing"
)

func TestAssignmentExpressions(t *testing.T) {
	input := `
x: 5
y: 10
foobar: 838383
`
	l := lexer.New(input)
	p := New(l, true)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 expressions. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		expr := program.Statements[i]
		if !testAssignmentExpression(t, expr, tt.expectedIdentifier){
			return
		}
	}
}

func testAssignmentExpression(t *testing.T, e ast.Expression, name string) bool{
	switch e.TokenLiteral() {
	case ":":
	case "::":
	case ":?":
	case ":&":
	default:
		return false
	}

	assExpr, ok := e.(*ast.AssignmentExpr)
	if !ok {
		t.Errorf("e not *ast.AssignmentExpr. got=%T", e)
		return false 
	}

	if assExpr.Name.Value != name {
		t.Errorf("assExpr.Name.Value not '%s'. got=%s", name, assExpr.Name.Value)
		return false 
	}

	if assExpr.Name.TokenLiteral() != name {
		t.Errorf("assExpr.Name.TokenLiteral not '%s'. got=%s", name, assExpr.Name.TokenLiteral())
		return false 
	}
	return true 
}

func checkParserErrors(t *testing.T, p *Parser){
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}