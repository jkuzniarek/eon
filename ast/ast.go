package ast

import (
	tk "eon/token"
	"bytes"
	ssDec "github.com/shopspring/decimal"
)

// ast node
type Node interface {
	TokenLiteral() string
	String() string
}

type Expression interface {
	Node 
	expressionNode()
}

// eventually rename to Root
type Program struct {
	Expressions []Expression 
}
func (p *Program) TokenLiteral() string {
	if len(p.Expressions) > 0 {
		return p.Expressions[0].TokenLiteral()
	} else {
		return ""
	}
}
func (p *Program) String() string{
	var out bytes.Buffer

	for _, s := range p.Expressions{
		out.WriteString(s.String())
	}

	return out.String()
}

type Card struct{
	Token tk.Token // the open delimiter token
	Type string // the type literal
	Size Expression // the array size specifier
	Index []Expression // name and infix assign expressions
	Body Expression // card body expression
}

func (c *Card) expressionNode(){}
func (c *Card) TokenLiteral() string {
	return c.Token.Literal
}
func (c *Card) String() string {
	var out bytes.Buffer
	iLen := len(c.Index)
	out.WriteString("<")

	if c.Type != "" {
		out.WriteString(c.Type)
	}

	if iLen != 0 {
		if iLen == 1 {
			out.WriteString(" ")
			out.WriteString(c.Index[0].String())
			out.WriteString("\n")
		}else{
			out.WriteString("\n")
			for _, e := range c.Index {
				out.WriteString(e.String())
				out.WriteString("\n")
			}
		}
	}

	if c.Body != nil {
		out.WriteString("/")
		out.WriteString(c.Body.String())
	}

	out.WriteString(">")
	return out.String()
}

type Input struct {
	Left Expression
	Input Expression
}

func (i *Input) expressionNode(){}
func (i *Input) TokenLiteral() string {
	return i.Left.TokenLiteral()
}
func (i *Input) String() string {
	var out bytes.Buffer

	out.WriteString(i.Left.String())
	out.WriteString(" ")
	out.WriteString(i.Input.String())
	
	return out.String()
}

type Infix struct {
	Token tk.Token // the operator token, eg +
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
	Token tk.Token // the open delimiter token
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

type Name struct {
	Token tk.Token // usually the tk.NAME token
	Value string
}
func (i *Name) expressionNode() {}
func (i *Name) TokenLiteral() string {
	return i.Token.Literal 
}
func (i *Name) String() string{
	return i.Value
}

// unsigned int
type UInt struct{
	Token tk.Token 
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
	Token tk.Token 
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
	Token tk.Token 
	Value ssDec.Decimal
}

func (il *Dec) expressionNode() {}
func (il *Dec) TokenLiteral() string {
	return il.Token.Literal
}
func (il *Dec) String() string{
	return il.Value.String()
}

// str
type Str struct{
	Token tk.Token 
	Value string
}

func (il *Str) expressionNode() {}
func (il *Str) TokenLiteral() string {
	return il.Token.Literal
}
func (il *Str) String() string{
	return il.Token.Literal
}

// byt
type Byt struct{
	Token tk.Token 
	Value []byte
}

func (il *Byt) expressionNode() {}
func (il *Byt) TokenLiteral() string {
	return il.Token.Literal
}
func (il *Byt) String() string{
	return il.Token.Literal
}

// comment
type Comment struct{
	Token tk.Token 
	Value string
	Multiline bool
}

func (n *Comment) expressionNode() {}
func (n *Comment) TokenLiteral() string {
	return n.Token.Literal
}
func (n *Comment) String() string{
	var out bytes.Buffer
	if n.Multiline {
		out.WriteString("/*")
		out.WriteString(n.Value)
		out.WriteString("*/")
	}else{
		out.WriteString("//")
		out.WriteString(n.Value)
		out.WriteString("\n")
	}
	return out.String()
}