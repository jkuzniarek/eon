package lexer


import (
	"testing"
	
	"eon/token"
)

func TestNextToken(t *testing.T) {
	input := `five: 5
	ten: 10
	// comment
	/*
	long comment
	*/
	hello:? fn(
		out: 'hello'
	)
	result.hello
	null: <>
	list: {}
	array: []
	try (
		(
		void
		)
		(
			esc
		)
	)
	`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.NAME, "five"},
		{token.SET_VAL, ":"},
		{token.INT, "5"},
		{token.EOL, "\n"},
		{token.NAME, "ten"},
		{token.SET_VAL, ":"},
		{token.INT, "10"},
		{token.EOL, "\n"},
		{token.COMMENT, " comment"},
		{token.COMMENT, "\n	long comment\n	"},
		{token.EOL, "\n"},
		{token.NAME, "hello"},
		{token.SET_WEAK, ":?"},
		{token.FN, "fn"},
		{token.LPAREN, "("},
		{token.EOL, "\n"},
		{token.OUT, "out"},
		{token.SET_VAL, ":"},
		{token.STR, "'hello'"},
		{token.EOL, "\n"},
		{token.RPAREN, ")"},
		{token.EOL, "\n"},
		{token.NAME, "result"},
		{token.DOT, "."},
		{token.NAME, "hello"},
		{token.EOL, "\n"},
		{token.NAME, "null"},
		{token.SET_VAL, ":"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EOL, "\n"},
		{token.NAME, "list"},
		{token.SET_VAL, ":"},
		{token.LCURLY, "{"},
		{token.RCURLY, "}"},
		{token.EOL, "\n"},
		{token.NAME, "array"},
		{token.SET_VAL, ":"},
		{token.LSQUAR, "["},
		{token.RSQUAR, "]"},
		{token.EOL, "\n"},
		{token.TRY, "try"},
		{token.LPAREN, "("},
		{token.EOL, "\n"},
		{token.LPAREN, "("},
		{token.EOL, "\n"},
		{token.VOID, "void"},
		{token.EOL, "\n"},
		{token.RPAREN, ")"},
		{token.EOL, "\n"},
		{token.LPAREN, "("},
		{token.EOL, "\n"},
		{token.ESC, "esc"},
		{token.EOL, "\n"},
		{token.RPAREN, ")"},
		{token.EOL, "\n"},
		{token.RPAREN, ")"},
		{token.EOL, "\n"},
		{token.EOF, ""},
	}

	// L
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q\nrow=%d, col=%d\nlastChar=%q", 
				i, tt.expectedType.ToStr(), tok.Type.ToStr(), l.row, l.col, l.peekLChar())
		}
		
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q\nrow=%d, col=%d\nlastChar=%q", 
				i, tt.expectedLiteral, tok.Literal, l.row, l.col, l.peekLChar())
		}
	}
}