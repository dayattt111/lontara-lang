package evaluator

import (
	"testing"

	"github.com/dayattt111/lontara-lang/pkg/lexer"
	"github.com/dayattt111/lontara-lang/pkg/object"
	"github.com/dayattt111/lontara-lang/pkg/parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	// Jika parser menghasilkan error, hentikan dan tampilkan pesan error
	if len(p.Errors()) > 0 {
		return &object.Error{Message: fmt.Sprintf("parser error: %v", p.Errors())}
	}

	env := object.NewEnvironment()
	return Eval(program, env)
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"50 / 2 * 2 + 10", 60},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"tongeng", true},
		{"banna", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 == 1", true},
		{"1 != 2", true},
		{"tongeng == tongeng", true},
		{"banna == tongeng", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestRekkoSangadinnaExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"rekko (tongeng) { 10 }", 10},
		{"rekko (banna) { 10 }", nil},
		{"rekko (1 < 2) { 10 } sangadinna { 20 }", 10},
		{"rekko (1 > 2) { 10 } sangadinna { 20 }", 20},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestTaroiStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"taroi a = 5; a;", 5},
		{"taroi a = 5 * 5; a;", 25},
		{"taroi a = 5; taroi b = a; b;", 5},
		{"taroi a = 5; taroi b = a; taroi c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"taroi identitas = jamagau(x) { x; }; identitas(5);", 5},
		{"taroi dobel = jamagau(x) { x * 2; }; dobel(5);", 10},
		{"taroi tambah = jamagau(x, y) { x + y; }; tambah(5, 5);", 10},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

// Helper assertions
func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("obj bukan *object.Integer. didapat=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("nilai salah. diharapkan=%d, didapat=%d", expected, result.Value)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("obj bukan *object.Boolean. didapat=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("nilai salah. diharapkan=%t, didapat=%t", expected, result.Value)
		return false
	}
	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("obj bukan NULL. didapat=%T (%+v)", obj, obj)
		return false
	}
	return true
}
