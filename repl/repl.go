package repl

import (
	"bufio"
	"fmt"
	"io"
	"eon/lexer"
	"eon/parser"
	"eon/eval"
	"eon/card"
)

const PROMPT = ">_ "

func Shell(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := card.NewEnv()
	inputStack := "" 
	stackDepth := 0
	var inputLog []string

	for {
		if inputStack == "" {
			fmt.Fprintf(out, PROMPT)
		} else {
			offset := "   "
			for stackDepth > 0 {
				offset += "  "
				stackDepth--
			}
			fmt.Fprintf(out, offset)
		}
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		input := scanner.Text()+"\n"
		line := inputStack+input
		// io.WriteString(out, line)
		l := lexer.New(line)
		p := parser.New(l, true)

		program := p.ParseShell()

		if l.Depth > 0 || p.Depth > 0 {
			inputStack = inputStack + input
			stackDepth = p.Depth
			continue
		}else if len(p.Errors()) != 0 {
			inputLog = append(inputLog, inputStack+input)
			inputStack = ""
			stackDepth = 0
			printParserErrors(out, p.Errors(), p.Trace)
			continue
		}else{
			inputLog = append(inputLog, inputStack+input)
			inputStack = ""
			stackDepth = 0
		}

		// print ast for debugging
		// io.WriteString(out, program.String())
		// io.WriteString(out, "\n")

		// eval program
		evaluated := eval.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.String())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string, t string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
	io.WriteString(out, "\tTrace: "+t+"\n")
}