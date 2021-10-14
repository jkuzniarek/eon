package parser

import (
	"eon/ast"
	"eon/lexer"
	tk "eon/token"
	"strconv"
	"fmt"
)


func (p *Parser) parseSInt() ast.Expression {
	lit := &ast.SInt{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 10, 0)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as signed integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = int(value)
	return lit
}

func (p *Parser) parseUInt() ast.Expression {
	lit := &ast.UInt{Token: p.curToken}
	value, err := strconv.ParseUint(p.curToken.Literal, 10, 0)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as unsigned integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = uint(value)
	return lit
}

func (p *Parser) parseDec() ast.Expression {
	lit := &ast.Dec{Token: p.curToken}
	value, err := decimal.NewFromString(p.curToken.Literal)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as decimal", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseStr() ast.Expression {
// TODO
	lit := &ast.Str{Token: p.curToken}
	src := p.curToken.Literal
	srcLen := len(src)
	ch := src[0]
	position := 1
	strStart := 1
	strEnd := 1
	another := true
	value := ""
	for another {
		strStart = position
		strEnd = position
		for src[position] != ch {
			position++
			strEnd++
		}
		value = append(value, src[strStart:strEnd])
		if position == (srcLen-1) {
			another = false
		} else if position < (srcLen-1) && src[position+1] == ch {
			value = append(value, string(ch))
			position = position + 2
		}
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseBytes() ast.Expression {
lit := &ast.Byt{Token: p.curToken}
src := p.curToken.Literal
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
	if count == 0 {
		return [...]byte{}
	}
	value := [count]byte 
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
		if isDecChar(src[position]){
			for i < 3; i++ {
				// ensure digit range only encompasses digits of ints between 000-255
				if i == 0 && (48 <= src[position] <= 50){
					i++
					position++
				} else if (
					i == 1 && 
					(48 <= src[position-1] <= 49 || 
						(src[position-1] == 50 && (48 <= src[position] <= 53)))
				){
					i++
					position++
				} else if (
					i == 2 && 
					(48 <= src[position-2] <= 49 || 
						(src[position-2] == 50 && (48 <= src[position-1] <= 52 ||
							(src[position-1] == 53 && (48 <= src[position] <= 53)) )))
				){
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
	if count == 0 {
		return [...]byte{}
	}
	value := [count]byte 
	count = 0
	position = 2
	for src[position] != '\\' {
		if isDecChar(src[position]){
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
			for i < 8; i++ {
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
	if count == 0 {
		return [...]byte{}
	}
	value := [count]byte 
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