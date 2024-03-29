package parser

import (
	tk "eon/token"
)



func getPrecedence(token tk.TokenType) int {
	switch token {
	case tk.EVAL_OPERATOR:
		return EQUALS
	case tk.ASSIGN_OPERATOR:
		return ASSIGN
	case tk.ACCESS_OPERATOR:
		return CALL
	default:
		return LOWEST
	}
}

func isHexChar(ch byte) bool{
	if 48 <= ch && ch <= 57 {
		return true
	} else if 65 <= ch && ch <= 70 {
		return true
	}
	return false
}

func isIntChar(ch byte) bool{
	if 48 <= ch && ch <= 57 {
		return true
	}
	return false
}

func isBinChar(ch byte) bool{
	if 48 <= ch && ch <= 49 {
		return true
	}
	return false
}

func hexToByte(src string) byte {
	i := 0
	out := make([]byte, 2)
	for i < 2 {
		switch src[i] {
		case 48:
			out[i] = 0
		case 49:
			out[i] = 1
		case 50:
			out[i] = 2
		case 51:
			out[i] = 3
		case 52:
			out[i] = 4
		case 53:
			out[i] = 5
		case 54:
			out[i] = 6
		case 55:
			out[i] = 7
		case 56:
			out[i] = 8
		case 57:
			out[i] = 9
		case 65:
			out[i] = 10
		case 66:
			out[i] = 11
		case 67:
			out[i] = 12
		case 68:
			out[i] = 13
		case 69:
			out[i] = 14
		case 70:
			out[i] = 15
		}
		i += 1
	}
	return (src[0]*16)+src[1]
}

func decToByte(src string) byte {
	i := 0
	out := make([]byte, 3)
	for i < 3 {
		switch src[i] {
		case 48:
			out[i] = 0
		case 49:
			out[i] = 1
		case 50:
			out[i] = 2
		case 51:
			out[i] = 3
		case 52:
			out[i] = 4
		case 53:
			out[i] = 5
		case 54:
			out[i] = 6
		case 55:
			out[i] = 7
		case 56:
			out[i] = 8
		case 57:
			out[i] = 9
		}
		i += 1
	}
	return (src[0]*100)+(src[1]*10)+src[2]
}

func binToByte(src string) byte {
	i := 0
	out := make([]byte, 8)
	for i < 8 {
		switch src[i] {
		case 48:
			out[i] = 0
		case 49:
			out[i] = 1
		}
		i += 1
	}
	return (src[0]*128)+(src[1]*64)+(src[2]*32)+(src[3]*16)+(src[4]*8)+(src[5]*4)+(src[6]*2)+src[7]
}