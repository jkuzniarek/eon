package parser

import (
	"fmt"
	"eon/ast"
	"eon/lexer"
	"testing"
	"eon/token"
)

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
	if len(program.Expressions) != 3 {
		t.Fatalf("program.Expressions does not contain 3 expressions. got=%d", len(program.Expressions))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		expr := program.Expressions[i]
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

func TestIdentifierExpression(t *testing.T) {
	input := "foobar"

	l := lexer.New(input)
	p := New(l, true)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Expressions) != 1 {
		t.Fatalf("program has not enough Expressions. got=%d", len(program.Expressions))
	}
	cmd, ok := program.Expressions[0].(ast.Expression)
	if !ok {
		t.Fatalf("program.Expressions[0] is not ast.Expression. got=%T", program.Expressions[0])
	}

	name, ok := cmd.(*ast.Identifier)
	if !ok {
		t.Fatalf("cmd not *ast.Identifier. got=%T", cmd)
	}
	if name.Value != "foobar" {
		t.Errorf("name.Value not %s. got=%s", "foobar", name.Value)
	}
	if name.TokenLiteral() != "foobar" {
		t.Errorf("name.TokenLiteral not %s. got=%s", "foobar", name.TokenLiteral())
	}
}

func TestUIntegerLiteralExpression(t *testing.T) {
	input := "5"

	l := lexer.New(input)
	p := New(l, true)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Expressions) != 1 {
		t.Fatalf("program has not enough Expressions. got=%d", len(program.Expressions))
	}
	cmd, ok := program.Expressions[0].(ast.Expression)
	if !ok {
		t.Fatalf("program.Expressions[0] is not ast.Expression. got=%T", program.Expressions[0])
	}

	literal, ok := cmd.(*ast.UIntegerLiteral)
	if !ok {
		t.Fatalf("cmd not *ast.UIntegerLiteral. got=%T", cmd)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}
}

func TestSIntegerLiteralExpression(t *testing.T) {
	input := "-5"

	l := lexer.New(input)
	p := New(l, true)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Expressions) != 1 {
		t.Fatalf("program has not enough Expressions. got=%d", len(program.Expressions))
	}
	cmd, ok := program.Expressions[0].(ast.Expression)
	if !ok {
		t.Fatalf("program.Expressions[0] is not ast.Expression. got=%T", program.Expressions[0])
	}

	literal, ok := cmd.(*ast.SIntegerLiteral)
	if !ok {
		t.Fatalf("cmd not *ast.SIntegerLiteral. got=%T", cmd)
	}
	if literal.Value != -5 {
		t.Errorf("literal.Value not %d. got=%d", -5, literal.Value)
	}
	if literal.TokenLiteral() != "-5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "-5", literal.TokenLiteral())
	}
}


// func testUIntegerLiteral(t *testing.T, il ast.Expression, value uint) bool {
// 	integ, ok := il.(*ast.UIntegerLiteral)
// 	if !ok {
// 		t.Errorf("il not *ast.UIntegerLiteral. got=%T", il)
// 		return false
// 	}

// 	if integ.Value != value {
// 		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
// 		return false
// 	}

// 	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
// 		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
// 		return false
// 	}

// 	return true
// }