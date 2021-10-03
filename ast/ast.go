package ast

import (
	"eon/token"
	"bytes"
	decimal "github.com/shopspring/decimal"
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

// dec
type Dec struct{
	Token token.Token 
	Value Decimal
}

func (il *Dec) expressionNode() {}
func (il *Dec) TokenLiteral() string {
	return il.Token.Literal
}
func (il *Dec) String() string{
	return il.Token.Literal
}

// str
type Str struct{
	Token token.Token 
	Value string
}

func (il *Str) expressionNode() {}
func (il *Str) TokenLiteral() string {
	return il.Token.Literal
}
func (il *Str) String() string{
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

type Group struct{
	Token token.Token // the open delimiter token
	Expressions []Expression 
}

func (g *Group) expressionNode(){}
func (g *Group) TokenLiteral() string {
	return g.Token.Literal
}
func (g *Group) String() string {
	var out bytes.Buffer

	for _, e := range g.Expressions {
		out.WriteString(e.String())
	}

	return out.String()
}
