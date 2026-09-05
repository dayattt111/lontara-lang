package token

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
)

// Kamus kata kunci Bugis Latin & Aksara Lontara
var keywords = map[string]TokenType{
	// Versi Latin
	"jamagau":    JAMAGAU,
	"taroi":      TAROI,
	"rekko":      REKKO,
	"sangadinna": SANGADINNA,
	"tongeng":    TONGENG,
	"banna":      BANNA,
	"paui":       PAUI,
	"lisu":       LISU,

	// Versi Aksara Lontara (Unicode U+1A00 - U+1A1F)
	"ᨍᨆᨁᨕᨘ":  JAMAGAU,    // jamagau
	"ᨈᨑᨚᨕᨗ":    TAROI,      // taroi
	"ᨑᨙᨀᨚ":      REKKO,      // rekko
	"ᨔᨂᨉᨗᨊ":    SANGADINNA, // sangadinna
	"ᨈᨚᨂᨙ":      TONGENG,    // tongeng
	"ᨅᨊ":        BANNA,      // banna
	"ᨄᨕᨘᨕᨗ":    PAUI,       // paui
	"ᨒᨗᨔᨘ":      LISU,       // lisu
}

// LookupIdent memeriksa apakah sebuah kata adalah keyword atau identifier biasa
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
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
