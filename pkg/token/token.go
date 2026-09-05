package token

import (
	"github.com/dayattt111/lontara-lang/pkg/dictionary"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	// Spesial
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier & Literal
	IDENT  = "IDENT"  // nama variabel / fungsi (misal: umur, x)
	INT    = "INT"    // 12345
	STRING = "STRING" // "halo dunia"

	// Operator Aritmetika
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// Operator Pembanding
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="
	LTE    = "<="
	GTE    = ">="

	// Delimiter
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Kata Kunci (Keywords) - Bugis Latin & Aksara Lontara
	JAMAGAU    = "JAMAGAU"    // fungsi (func / function)
	TAROI      = "TAROI"      // deklarasi variabel (let / var)
	REKKO      = "REKKO"      // if
	SANGADINNA = "SANGADINNA" // else
	TONGENG    = "TONGENG"    // true
	BANNA      = "BANNA"      // false
	PAUI       = "PAUI"       // cetak (print / println)
	LISU       = "LISU"       // return
	SIKI       = "SIKI"       // perulangan (while / loop)
)

// LookupIdent memeriksa apakah sebuah kata adalah keyword (terdaftar di pkg/dictionary) atau identifier biasa
func LookupIdent(ident string) TokenType {
	if tokType, ok := dictionary.LookupKeyword(ident); ok {
		return TokenType(tokType)
	}
	return IDENT
}

// New membuat token baru dengan informasi posisi baris dan kolom
func New(tokenType TokenType, ch string, line, col int) Token {
	return Token{
		Type:    tokenType,
		Literal: ch,
		Line:    line,
		Column:  col,
	}
}
