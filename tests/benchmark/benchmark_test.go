package benchmark

import (
	"testing"

	"github.com/dayattt111/lontara-lang/pkg/evaluator"
	"github.com/dayattt111/lontara-lang/pkg/lexer"
	"github.com/dayattt111/lontara-lang/pkg/object"
	"github.com/dayattt111/lontara-lang/pkg/parser"
)

// BenchmarkLontaraLoop menguji kecepatan eksekusi 1000 iterasi perulangan siki
func BenchmarkLontaraLoop(b *testing.B) {
	code := `
taroi i = 1;
taroi total = 0;
siki (i <= 1000) {
    taroi total = total + i;
    taroi i = i + 1;
}
`
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l := lexer.New(code)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		evaluator.Eval(program, env)
	}
}

// BenchmarkLontaraFunction menguji kecepatan pemanggilan fungsi jamagau
func BenchmarkLontaraFunction(b *testing.B) {
	code := `
taroi tambah = jamagau(x, y) {
    lisu x + y;
};
taroi i = 0;
siki (i < 500) {
    taroi res = tambah(i, 10);
    taroi i = i + 1;
}
`
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l := lexer.New(code)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		evaluator.Eval(program, env)
	}
}
