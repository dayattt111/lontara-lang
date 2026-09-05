package lexer

import (
	"testing"

	"github.com/dayattt111/lontara-lang/pkg/token"
)

func TestNextTokenBugisLatin(t *testing.T) {
	input := `taroi lima = 5;
taroi sepulo = 10;

jamagau tambah(x, y) {
    lisu x + y;
};

taroi aselleng = tambah(lima, sepulo);
rekko (aselleng >= 15) {
    paui("sitinaja");
} sangadinna {
    paui("kurang");
}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.TAROI, "taroi"},
		{token.IDENT, "lima"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.TAROI, "taroi"},
		{token.IDENT, "sepulo"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.JAMAGAU, "jamagau"},
		{token.IDENT, "tambah"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		{token.LISU, "lisu"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.TAROI, "taroi"},
		{token.IDENT, "aselleng"},
		{token.ASSIGN, "="},
		{token.IDENT, "tambah"},
		{token.LPAREN, "("},
		{token.IDENT, "lima"},
		{token.COMMA, ","},
		{token.IDENT, "sepulo"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.REKKO, "rekko"},
		{token.LPAREN, "("},
		{token.IDENT, "aselleng"},
		{token.GTE, ">="},
		{token.INT, "15"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		{token.PAUI, "paui"},
		{token.LPAREN, "("},
		{token.STRING, "sitinaja"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},
		{token.SANGADINNA, "sangadinna"},
		{token.LBRACE, "{"},

		{token.PAUI, "paui"},
		{token.LPAREN, "("},
		{token.STRING, "kurang"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tipe token salah. diharapkan=%q, didapat=%q (literal=%q)",
				i, tt.expectedType, tok.Type, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal salah. diharapkan=%q, didapat=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenAksaraLontara(t *testing.T) {
	// Uji coba kata kunci asli beraksara Lontara
	input := `ᨈᨑᨚᨕᨗ angka = 10;
ᨑᨙᨀᨚ (angka > 5) {
    ᨄᨕᨘᨕᨗ("salama");
}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.TAROI, "ᨈᨑᨚᨕᨗ"},
		{token.IDENT, "angka"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.REKKO, "ᨑᨙᨀᨚ"},
		{token.LPAREN, "("},
		{token.IDENT, "angka"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		{token.PAUI, "ᨄᨕᨘᨕᨗ"},
		{token.LPAREN, "("},
		{token.STRING, "salama"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("lontara tests[%d] - tipe token salah. diharapkan=%q, didapat=%q (literal=%q)",
				i, tt.expectedType, tok.Type, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("lontara tests[%d] - literal salah. diharapkan=%q, didapat=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestEOFIdentifiersAndNumbers(t *testing.T) {
	testCases := []struct {
		input           string
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{"tongeng", token.TONGENG, "tongeng"},
		{"banna", token.BANNA, "banna"},
		{"10", token.INT, "10"},
		{"5", token.INT, "5"},
		{"ᨈᨑᨚᨕᨗ", token.TAROI, "ᨈᨑᨚᨕᨗ"},
	}

	for _, tt := range testCases {
		l := New(tt.input)
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Errorf("input %q: tipe token salah. diharapkan=%q, didapat=%q", tt.input, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("input %q: literal salah. diharapkan=%q, didapat=%q", tt.input, tt.expectedLiteral, tok.Literal)
		}

		eofTok := l.NextToken()
		if eofTok.Type != token.EOF {
			t.Errorf("input %q: diharapkan EOF, didapat=%q", tt.input, eofTok.Type)
		}
	}
}
