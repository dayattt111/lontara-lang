package ast

import (
	"bytes"
	"strings"

	"github.com/dayattt111/lontara-lang/pkg/token"
)

// Node adalah antarmuka dasar untuk semua simpul di AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement adalah simpul kode yang tidak menghasilkan nilai (aksi)
type Statement interface {
	Node
	statementNode()
}

// Expression adalah simpul kode yang menghasilkan nilai saat dievaluasi
type Expression interface {
	Node
	expressionNode()
}

// Program adalah root node dari keseluruhan file kode
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// --- STATEMENTS ---

// TaroiStatement: Deklarasi variabel (misal: taroi x = 5;)
type TaroiStatement struct {
	Token token.Token // token.TAROI
	Name  *Identifier
	Value Expression
}

func (ts *TaroiStatement) statementNode()       {}
func (ts *TaroiStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *TaroiStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ts.TokenLiteral() + " ")
	out.WriteString(ts.Name.String())
	out.WriteString(" = ")
	if ts.Value != nil {
		out.WriteString(ts.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// LisuStatement: Pernyataan return (misal: lisu x;)
type LisuStatement struct {
	Token       token.Token // token.LISU
	ReturnValue Expression
}

func (ls *LisuStatement) statementNode()       {}
func (ls *LisuStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LisuStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	if ls.ReturnValue != nil {
		out.WriteString(ls.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement: Ekspresi yang berdiri sendiri sebagai baris kode (misal: 1 + 2;)
type ExpressionStatement struct {
	Token      token.Token // token pertama dari ekspresi
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// BlockStatement: Kumpulan statement di dalam blok kurung kurawal { ... }
type BlockStatement struct {
	Token      token.Token // token.LBRACE
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// --- EXPRESSIONS ---

// Identifier: Nama variabel / fungsi (misal: umur, tambah)
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// IntegerLiteral: Angka bulat (misal: 10, 42)
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// StringLiteral: Teks di dalam tanda petik (misal: "sitinaja")
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

// Boolean: Nilai logika tongeng (true) / banna (false)
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// PrefixExpression: Ekspresi awalan (misal: !banna, -5)
type PrefixExpression struct {
	Token    token.Token // token awalan, misal '!' atau '-'
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression: Ekspresi biner di antara dua operand (misal: 5 + 10, x == y)
type InfixExpression struct {
	Token    token.Token // token operator (+, -, *, /, ==, dst.)
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}

// RekkoExpression: Percabangan if (rekko ... sangadinna ...)
type RekkoExpression struct {
	Token       token.Token // token.REKKO
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (re *RekkoExpression) expressionNode()      {}
func (re *RekkoExpression) TokenLiteral() string { return re.Token.Literal }

func (re *RekkoExpression) String() string {
	var out bytes.Buffer
	out.WriteString("rekko")
	out.WriteString(re.Condition.String())
	out.WriteString(" ")
	out.WriteString(re.Consequence.String())
	if re.Alternative != nil {
		out.WriteString(" sangadinna ")
		out.WriteString(re.Alternative.String())
	}
	return out.String()
}

// JamagauLiteral: Definisi fungsi (misal: jamagau(x, y) { lisu x + y; })
type JamagauLiteral struct {
	Token      token.Token // token.JAMAGAU
	Parameters []*Identifier
	Body       *BlockStatement
}

func (jl *JamagauLiteral) expressionNode()      {}
func (jl *JamagauLiteral) TokenLiteral() string { return jl.Token.Literal }
func (jl *JamagauLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range jl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(jl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(jl.Body.String())
	return out.String()
}

// CallExpression: Pemanggilan fungsi (misal: tambah(2, 3) atau paui("halo"))
type CallExpression struct {
	Token     token.Token // token kurung buka '('
	Function  Expression  // Identifier atau JamagauLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
