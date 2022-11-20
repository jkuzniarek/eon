package eval

import (
	"eon/ast"
	"eon/card"
)

var (
	VOID = &card.Void{}
	ANTIVOID = &card.Void{Value: VOID}
)

func Eval(node ast.Node) card.Card {
	switch node := node.(type) {

	// eval entry points
	case *ast.Program:
		return evalExpressions(node.Expressions)

	// expressions
	case *ast.UInt:
		return &card.UInt{ Value: node.Value}
	case *ast.Infix:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfix(node.Operator, left, right)

	}

	return nil
}

func evalExpressions(exprs []ast.Node) card.Card {
	var result card.Card

	for _, expression := range exprs {
		result = Eval(expression)
	}
	return result
}

func evalInfix( operator string, left card.Card, right card.Card) card.Card {
	switch {
	case left.IRType() == card.UINT && right.IRType() == card.UINT:
		return evalUIntInfix(operator, left, right)

	// pointer based comparisons
	case operator == "==":
		return boolToVoid(left == right)
	case operator == "!=":
		return boolToVoid(left != right)
	default:
		return VOID // TODO: update to handle errors
	}
}

func evalUIntInfix(operator string, left card.Card, right card.Card) card.Card {
	leftVal := left.(*card.UInt).Value
	rightVal := right.(*card.UInt).Value

	switch operator {
		// math is done via function calls not operators (except for increment/decrement :+ :- )
	// case "+":
	// 	return &card.UInt{Value: leftVal + rightVal}
	// case "-":
	// 	return &card.UInt{Value: leftVal - rightVal}
	// case "*":
	// 	return &card.UInt{Value: leftVal * rightVal}
	// case "/":
	// 	return &card.UInt{Value: leftVal / rightVal}
	case "<":
		return boolToVoid(leftVal < rightVal)
	case ">":
		return boolToVoid(leftVal > rightVal)
	case "<=":
		return boolToVoid(leftVal <= rightVal)
	case ">=":
		return boolToVoid(leftVal >= rightVal)
	case "==":
		return boolToVoid(leftVal == rightVal)
	case "!=":
		return boolToVoid(leftVal != rightVal)
	default:
		return VOID
	}
}

