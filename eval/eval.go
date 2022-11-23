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

	// usual entry point
	case *ast.Program:
		return evalExpressions(node.Expressions, env)

	// expressions
	case *ast.UInt:
		return &card.UInt{ Value: node.Value}
	case *ast.Name:
		return evalName(node, env, nil)
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
	case *ast.Input:
		val := Eval(node.Input, env)
		if isError(val) {
			return val
		}

		// left is name/infix, right is expression
		switch node.Left.(type) {
		case *ast.Name:
			return evalName(node.Left.(*ast.Name), env, val)
			// TODO: build evalInfixAccess
		// case *ast.Infix:
		// 	return evalInfixAccess(node.Left.(*ast.Infix), env, val)
		default:
			return newError("unhandled input node: %s %s", node.String(), val.IRType().String())
		}
	}

	return nil
}

func evalExpressions(exprs []ast.Node, env *card.Env) card.Card {
	var result card.Card

	for _, expression := range exprs {
		result = Eval(expression, env)

		switch result := result.(type) {
		case *card.OutVal:
			return result.Value
		case *card.Error:
			return result
		}
	}
	return result // this line should only apply when run in the shell
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
	switch left.(type) { // TODO: add ifelse for nested data structures and accessors
	case *ast.Name:
		if left.(*ast.Name).Token.Type == tk.KEYWORD {
			return VOID;
		}
		switch operator {
		case ":":
			return env.Set(left.(*ast.Name).Value, val)
		default:
			return newError("unknown operator: %s %s %s", left.String(), operator, val.IRType().String())
		}
	default:
		return newError("unhandled name node: %s %s %s", left.String(), operator, val.IRType().String())
	}

	return ANTIVOID
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

func evalName(node *ast.Name, env *card.Env, input card.Card) card.Card {
	// check if name is keyword
	if input == nil && node.Token.Type == tk.KEYWORD {
		switch node.Value {
		case "void":
			out, ok := env.Get("out")
			if !ok { // TODO: convert this to return a VOID
				out = nil
			}
			return &card.OutVal{Value: out}
		// case "esc":
		// 	return &card.Esc
		default:
			return newError("unhandled keyword node: %s", node.String())
		}
	}
	//
	val, ok := env.Get(node.Value)
	if !ok { 
		val = VOID
	}

	if val.IRType() == card.FUNCTION {
		return applyFunction(val, nil, input)
	}
	// TODO: use code to adapt evalInfixAccess to nested data structures
	return val
}

func applyFunction(fn card.Card, l_arg card.Card, r_arg card.Card) card.Card {
	function, ok := fn.(*card.Function)
	if !ok {
		// ?TODO: convert to return a VOID
		return newError("not a function: %s", fn.IRType().String())
	}

	extendedEnv := extendFunctionEnv(function, nil, r_arg)
	evaluated := Eval(function.Body, extendedEnv)
	return unwrapOutVal(evaluated)
}

// TODO: add l_arg param, consider changing name from src -> argl
func extendFunctionEnv(fn *card.Function, l_arg card.Card, r_arg card.Card) *card.Env {
	env := card.NewChildEnv(fn.Env)

	if l_arg == nil {
		env.Set("argl", VOID)
	}else{
		env.Set("argl", l_arg)
	}
	if r_arg == nil {
		env.Set("argr", VOID)
	}else{
		env.Set("argr", r_arg)
	}
	env.Set("out", ANTIVOID)

	return env
}

func unwrapOutVal(obj card.Card) card.Card {
	if outVal, ok := obj.(*card.OutVal); ok {
		return outVal.Value
	}

	return obj
}