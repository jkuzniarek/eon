package parser

import (
	"eon/ast"
)

func (p *Parser) parseBytes() ast.Node {
	lit := &ast.Byt{Token: p.curToken}
	src := p.curToken.Literal
	var value []byte
	ch := src[1]
	position := 2
	count := 0
	switch ch {
	case 'x':
		for src[position] != '\\' {
			if isHexChar(src[position]) && isHexChar(src[position+1]) {
				count++
				position = position + 2
			} else if isHexChar(src[position]) && !isHexChar(src[position+1]) {
				msg := fmt.Sprintf("could not parse %q as hexadecimal byte string", p.curToken.Literal)
				p.errors = append(p.errors, msg)
				return nil
			} else {
				position++
			}
		}
		value := make([]byte, count) 
		count = 0
		position = 2
		for src[position] != '\\' {
			if isHexChar(src[position]) && isHexChar(src[position+1]) {
				value[count] = hexToByte(src[position:(position+2)])
				count++
				position = position + 2
			} else {
				position++
			}
		}
	case 'd':
		i := 0
		for src[position] != '\\' {
			if isIntChar(src[position]){
				for i < 3 {
					// ensure digit range only encompasses digits of ints between 000-255
					if i == 0 && 48 <= src[position] && src[position] <= 50 {
						i++
						position++
					} else if (
						i == 1 && 
						((48 <= src[position-1] && src[position-1] <= 49) || 
							(src[position-1] == 50 && 48 <= src[position] && src[position] <= 53))){
						i++
						position++
					} else if (
						i == 2 && 
						((48 <= src[position-2] && src[position-2] <= 49) || 
							(src[position-2] == 50 && ((48 <= src[position-1] && src[position-1] <= 52) ||
								(src[position-1] == 53 && 48 <= src[position] && src[position] <= 53) )))){
						i++
						position++
					} else {
						msg := fmt.Sprintf("could not parse %q as decimal byte string", p.curToken.Literal)
						p.errors = append(p.errors, msg)
						return nil
					}
				}
				i = 0
				count++
			} else {
				position++
			}
		}
		value := make([]byte, count)
		count = 0
		position = 2
		for src[position] != '\\' {
			if isIntChar(src[position]){
				value[count] = decToByte(src[position:(position+3)])
				count++
				position = position + 3
			} else {
				position++
			}
		}
	case 'b':
		i := 0
		for src[position] != '\\' {
			if isBinChar(src[position]){
				i++
				position++
				for i < 8 {
					if isBinChar(src[position]) {
						i++
						position++
					} else {
						msg := fmt.Sprintf("could not parse %q as binary byte string", p.curToken.Literal)
						p.errors = append(p.errors, msg)
						return nil
					}
				}
				i = 0
				count++
			} else {
				position++
			}
		}
		value := make([]byte, count)
		count = 0
		position = 2
		for src[position] != '\\' {
			if isBinChar(src[position]){
				value[count] = binToByte(src[position:(position+8)])
				count++
				position = position + 8
			} else {
				position++
			}
		}
	default:
		msg := fmt.Sprintf("could not parse %q as byte string", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
	}