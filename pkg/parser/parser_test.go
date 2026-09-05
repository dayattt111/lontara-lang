package parser

import (
	"testing"

	"github.com/dayattt111/lontara-lang/pkg/ast"
	"github.com/dayattt111/lontara-lang/pkg/lexer"
)

func TestTaroiStatements(t *testing.T) {
	input := `
taroi x = 5;
taroi y = 10;
taroi foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements tidak memuat 3 statement. didapat=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if stmt.TokenLiteral() != "taroi" {
			t.Fatalf("TokenLiteral() bukan 'taroi'. didapat=%q", stmt.TokenLiteral())
		}
		taroiStmt, ok := stmt.(*ast.TaroiStatement)
		if !ok {
			t.Fatalf("stmt bukan *ast.TaroiStatement. didapat=%T", stmt)
		}
		if taroiStmt.Name.Value != tt.expectedIdentifier {
			t.Fatalf("taroiStmt.Name.Value bukan %s. didapat=%s", tt.expectedIdentifier, taroiStmt.Name.Value)
		}
	}
}

func TestOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"1 + 2 * 3;",
			"(1 + (2 * 3))",
		},
		{
			"rekko (x < y) { x; } sangadinna { y; };",
			"rekko(x < y) x sangadinna y",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("diharapkan=%q, didapat=%q", tt.expected, actual)
		}
	}
}

func TestSikiExpression(t *testing.T) {
	input := `siki (x < y) { x; }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements tidak memuat 1 statement. didapat=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] bukan *ast.ExpressionStatement. didapat=%T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.SikiExpression)
	if !ok {
		t.Fatalf("stmt.Expression bukan *ast.SikiExpression. didapat=%T", stmt.Expression)
	}

	if exp.TokenLiteral() != "siki" {
		t.Fatalf("exp.TokenLiteral() bukan 'siki'. didapat=%q", exp.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser memiliki %d error", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
