package ast

// ast node
type Node interface {
	TokenLiteral() string
	String() string
}
