package ast

import (
	"eon/token"
	"testing"
)

func TestString(t *testing.T){
	program := &Program{
		Expressions: []Expression{
			&AssignmentExpr{
				Token: token.Token{ Type: token.SET_VAL, Literal: ":"},
				Name: &Identifier{
					Token: token.Token{ Type:token.NAME, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{ Type: token.NAME, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "myVar: anotherVar" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}