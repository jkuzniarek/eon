package ast

import (
	"eon/token"
	"bytes"
)

// ast node
type Node interface {
	TokenLiteral() string
	String() string
}

// remove?
type Command interface {
	Node
	commandNode()
}

type Expression interface {
	Node 
	expressionNode()
}

// eventually rename to Root
type Program struct {
	Commands []Expression 
}
func (p *Program) TokenLiteral() string {
	if len(p.Commands) > 0 {
		return p.Commands[0].TokenLiteral()
	} else {
		return ""
	}
}
func (p *Program) String() string{
	var out bytes.Buffer

	for _, s := range p.Commands{
		out.WriteString(s.String())
	}

	return out.String()
}

type AssignmentExpr struct {
	Token token.Token // token indicating assignment type
	Name *Identifier // NAME token indicating name
	Value Expression // 
}
func (ae *AssignmentExpr) expressionNode() {}
func (ae *AssignmentExpr) TokenLiteral() string {
	return ae.Token.Literal 
}
func (ae *AssignmentExpr) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Name.String())
	out.WriteString(ae.TokenLiteral() + " ")

	if ae.Value != nil {
		out.WriteString(ae.Value.String())
	}

	return out.String()
}

type Identifier struct {
	Token token.Token // usually the token.NAME token
	Value string
}
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal 
}
func (i *Identifier) String() string{
	return i.Value
}

// unsigned int
type UInt struct{
	Token token.Token 
	Value uint
}

func (il *UInt) expressionNode() {}
func (il *UInt) TokenLiteral() string {
	return il.Token.Literal
}
func (il *UInt) String() string{
	return il.Token.Literal
}

// signed int
type SInt struct{
	Token token.Token 
	Value int
}

func (il *SInt) expressionNode() {}
func (il *SInt) TokenLiteral() string {
	return il.Token.Literal
}
func (il *SInt) String() string{
	return il.Token.Literal
}

type Infix struct {
	Token token.Token // the operator token, eg +
	Left Expression
	Operator string
	Right Expression
}

func (ie *Infix) expressionNode(){}
func (ie *Infix) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *Infix) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// consult for others
// type IfExpression struct {
// 	Token token.Token // The 'if' token
// 	Condition Expression
// 	Consequence *BlockCommand
// 	Alternative *BlockCommand
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
// type BlockCommand struct{
// 	Token token.Token // the { token
// 	Commands []Command 
// }

// func (bs *BlockCommand) commandNode(){}
// func (bs *BlockCommand) TokenLiteral() string {
// 	return bs.Token.Literal
// }
// func (bs *BlockCommand) String() string {
// 	var out bytes.Buffer

// 	for _, s := range bs.Commands {
// 		out.WriteString(s.String())
// 	}

// 	return out.String()
// }
