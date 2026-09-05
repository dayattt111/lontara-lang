package lexer

import (
	"unicode"
	"unicode/utf8"

	"github.com/dayattt111/lontara-lang/pkg/token"
)

type Lexer struct {
	input        string
	position     int  // posisi karakter saat ini dalam bytes
	readPosition int  // posisi baca berikutnya dalam bytes
	ch           rune // karakter saat ini yang sedang diperiksa
	line         int  // nomor baris aktif
	col          int  // nomor kolom aktif
}

// New menginisialisasi Lexer baru
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
		line:  1,
		col:   0,
	}
	l.readChar()
	return l
}

// readChar membaca satu karakter Unicode (rune) berikutnya dari input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF (End of File)
	} else {
		r, width := utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.ch = r
		l.position = l.readPosition
		l.readPosition += width
		l.col++
	}
}

// peekChar mengintip karakter berikutnya tanpa memajukan posisi pembacaan
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
	return r
}

// NextToken mengambil satu token berikutnya
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	curLine := l.line
	curCol := l.col

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.EQ, literal, curLine, curCol)
		} else {
			tok = token.New(token.ASSIGN, string(l.ch), curLine, curCol)
		}
	case '+':
		tok = token.New(token.PLUS, string(l.ch), curLine, curCol)
	case '-':
		tok = token.New(token.MINUS, string(l.ch), curLine, curCol)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.NOT_EQ, literal, curLine, curCol)
		} else {
			tok = token.New(token.BANG, string(l.ch), curLine, curCol)
		}
	case '/':
		tok = token.New(token.SLASH, string(l.ch), curLine, curCol)
	case '*':
		tok = token.New(token.ASTERISK, string(l.ch), curLine, curCol)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.LTE, literal, curLine, curCol)
		} else {
			tok = token.New(token.LT, string(l.ch), curLine, curCol)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.GTE, literal, curLine, curCol)
		} else {
			tok = token.New(token.GT, string(l.ch), curLine, curCol)
		}
	case ';':
		tok = token.New(token.SEMICOLON, string(l.ch), curLine, curCol)
	case ',':
		tok = token.New(token.COMMA, string(l.ch), curLine, curCol)
	case '(':
		tok = token.New(token.LPAREN, string(l.ch), curLine, curCol)
	case ')':
		tok = token.New(token.RPAREN, string(l.ch), curLine, curCol)
	case '{':
		tok = token.New(token.LBRACE, string(l.ch), curLine, curCol)
	case '}':
		tok = token.New(token.RBRACE, string(l.ch), curLine, curCol)
	case '"':
		str := l.readString()
		tok = token.New(token.STRING, str, curLine, curCol)
		return tok
	case 0:
		tok = token.New(token.EOF, "", curLine, curCol)
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			tokType := token.LookupIdent(ident)
			return token.New(tokType, ident, curLine, curCol)
		} else if isDigit(l.ch) {
			num := l.readNumber()
			return token.New(token.INT, num, curLine, curCol)
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch), curLine, curCol)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
			l.col = 0
		}
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	startPos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func (l *Lexer) readNumber() string {
	startPos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func (l *Lexer) readString() string {
	// Lewati tanda petik pembuka
	l.readChar()
	startPos := l.position
	for l.ch != '"' && l.ch != 0 {
		if l.ch == '\n' {
			l.line++
			l.col = 0
		}
		l.readChar()
	}
	res := l.input[startPos:l.position]
	// Lewati tanda petik penutup jika ada
	if l.ch == '"' {
		l.readChar()
	}
	return res
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// isLetter mendukung alfabet biasa, underscore, dan rentang karakter Aksara Bugis / Lontara
func isLetter(ch rune) bool {
	// Karakter Latin biasa dan garis bawah
	if unicode.IsLetter(ch) || ch == '_' {
		return true
	}
	// Blok Unicode Aksara Bugis / Lontara: U+1A00 sampai U+1A1F
	if ch >= 0x1A00 && ch <= 0x1A1F {
		return true
	}
	return false
}
