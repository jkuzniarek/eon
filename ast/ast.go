package ast

import (
	"eon/token"
	"bytes"
)

// ast node
type Node interface {
	TokenLiteral() string
}


// type Statement interface {
// 	Node
// 	statementNode()
// }

type Expression interface {
	Node 
	expressionNode()
}

// eventually rename to Root
type Program struct {
	Statements []Expression 
}
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
// func (p *Program) String() string{
// 	var out bytes.Buffer

// 	for _, s := range p.Statements{
// 		out.WriteString(s.String())
// 	}

// 	return out.String()
// }

type AssignmentExpr struct {
	Token token.Token // token indicating assignment type
	Name *Identifier // IDENT token indicating name
	Value Expression // 
}
func (i *AssignmentExpr) expressionNode() {}
func (i *AssignmentExpr) TokenLiteral() string {
	return i.Token.Literal 
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal 
}
// func (i *Identifier) String() string{
// 	return i.Value
// }

type ExpressionStatement struct{
	Token token.Token // the first token of the expression
	Expression Expression
}
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
// func (es *ExpressionStatement) String() string{
// 	if es.Expression != nil {
// 		return es.Expression.String()
// 	}
// 	return ""
// }

type IntegerLiteral struct{
	Token token.Token 
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
// func (il *IntegerLiteral) String() string{
// 	return il.Token.Literal
// }


type PrefixExpression struct {
	Token token.Token // the prefix token, eg !
	Operator string
	Right Expression 
}

func (pe *PrefixExpression) expressionNode(){}
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}
// func (pe *PrefixExpression) String() string {
// 	var out bytes.Buffer

// 	out.WriteString("(")
// 	out.WriteString(pe.Operator)
// 	out.WriteString(pe.Right.String())
// 	out.WriteString(")")

// 	return out.String()
// }

type InfixExpression struct {
	Token token.Token // the operator token, eg +
	Left Expression
	Operator string
	Right Expression
}

func (ie *InfixExpression) expressionNode(){}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}
// func (ie *InfixExpression) String() string {
// 	var out bytes.Buffer

// 	out.WriteString("(")
// 	out.WriteString(ie.Left.String())
// 	out.WriteString(" " + ie.Operator + " ")
// 	out.WriteString(ie.Right.String())
// 	out.WriteString(")")

// 	return out.String()
// }

// consult for others
// type IfExpression struct {
// 	Token token.Token // The 'if' token
// 	Condition Expression
// 	Consequence *BlockStatement
// 	Alternative *BlockStatement
// }

// func (ie *IfExpression) expressionNode(){}
// func (ie *IfExpression) TokenLiteral() string {
// 	return ie.Token.Literal
// }
// func (ie *IfExpression) String() string {
// 	var out bytes.Buffer

// 	out.WriteString("if")
// 	out.WriteString(ie.Condition.String())
// 	out.WriteString(" ")
// 	out.WriteString(ie.Consequence.String())

// 	if ie.Alternative != nil {
// 		out.WriteString("else ")
// 		out.WriteString(ie.Alternative.String())
// 	}

// 	return out.String()
// }

// consult for linked_list, array
// type BlockStatement struct{
// 	Token token.Token // the { token
// 	Statements []Statement 
// }

// func (bs *BlockStatement) statementNode(){}
// func (bs *BlockStatement) TokenLiteral() string {
// 	return bs.Token.Literal
// }
// func (bs *BlockStatement) String() string {
// 	var out bytes.Buffer

// 	for _, s := range bs.Statements {
// 		out.WriteString(s.String())
// 	}

// 	return out.String()
// }
