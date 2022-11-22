package eval

import (
	"eon/ast"
	"eon/card"
	tk "eon/token"
)

var (
	VOID = &card.Void{}
	ANTIVOID = &card.Void{Value: VOID}
)

func Eval(node ast.Node, env *card.Env) card.Card {
	switch node := node.(type) {

	// eval entry point
	case *ast.Program:
		return evalExpressions(node.Expressions, env)

	// expressions
	case *ast.UInt:
		return &card.UInt{ Value: node.Value}
	case *ast.Name:
		return evalName(node, env)
	case *ast.Infix:
		switch tk.GetOperatorType(node.Operator) {
		case tk.EVAL_OPERATOR:
			left := Eval(node.Left, env)
			if isError(left) {
				return left
			}
			right := Eval(node.Right, env)
			if isError(right) {
				return right
			}
			return evalInfixEval(node.Operator, left, right)
		case tk.ASSIGN_OPERATOR:
			val := Eval(node.Right, env)
			if isError(val) {
				return val
			}
			return evalInfixAssign(node.Left, node.Operator, val, env)

		// case tk.ACCESS_OPERATOR: // TODO

		}
	case *ast.Group:
		switch node.Token.Type {
		// case tk.CPAREN:
			
		case tk.HPAREN:
			return &card.Function{Body: node, Env: env}
		}

	}

	return nil
}

func evalExpressions(exprs []ast.Node, env *card.Env) card.Card {
	var result card.Card

	for _, expression := range exprs {
		result = Eval(expression, env)

		switch result := result.(type) {
		// classical return statement which includes stopping execution
		// case *card.ReturnValue:
		// 	return result.Value
		case *card.Error:
			return result
		}
	}
	return result
}

func evalInfixEval(operator string, left card.Card, right card.Card) card.Card {
	switch {
	case left.IRType() == card.UINT && right.IRType() == card.UINT:
		return evalInfixUInt(operator, left, right)

	// pointer based comparisons
	case operator == "==":
		return boolToVoid(left == right)
	case operator == "!=":
		return boolToVoid(left != right)
	case left.IRType() != right.IRType():
		return newError("type mismatch: %s %s %s", left.IRType().String(), operator, right.IRType().String())
	default:
		return newError("unknown operator: %s %s %s", left.IRType().String(), operator, right.IRType().String())
	}
}

func evalInfixAssign(left ast.Node, operator string, val card.Card, env *card.Env) card.Card {
	switch operator {
	case ":":
		switch left.(type) { // TODO: add ifelse for nested data structures and accessors
		case *ast.Name:
			return env.Set(left.(*ast.Name).Value, val) 
		default:
			return newError("unhandled name node: %s %s %s", left.String(), operator, val.IRType().String())
		}
		
		return ANTIVOID
	default:
		return newError("unknown operator: %s %s %s", left.String(), operator, val.IRType().String())
	}
}

func evalInfixUInt(operator string, left card.Card, right card.Card) card.Card {
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
		return newError("unknown operator: %s %s %s", left.IRType().String(), operator, right.IRType().String())
	}
}

func evalName(node *ast.Name, env *card.Env) card.Card {
	val, ok := env.Get(node.Value)
	if !ok {
		return newError("name not found: " + node.Value)
	}
	// TODO: use code to adapt evalInfixAccess to nested data structures
	return val
}