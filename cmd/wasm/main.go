package main

import (
	"bytes"
	"fmt"
	"syscall/js"

	"github.com/dayattt111/lontara-lang/pkg/evaluator"
	"github.com/dayattt111/lontara-lang/pkg/lexer"
	"github.com/dayattt111/lontara-lang/pkg/object"
	"github.com/dayattt111/lontara-lang/pkg/parser"
)

func evaluateLontaraWrapper(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return "Kesalahan: Tidak ada kode masukan yang diberikan."
	}
	code := args[0].String()

	var outBuf bytes.Buffer
	evaluator.SetOutput(&outBuf)

	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		var errOut bytes.Buffer
		errOut.WriteString("Kesalahan Sintaks (Parser Error):\n")
		for _, msg := range p.Errors() {
			errOut.WriteString(fmt.Sprintf("  - %s\n", msg))
		}
		return errOut.String()
	}

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)

	printedOutput := outBuf.String()

	if evaluated != nil && evaluated.Type() == object.ERROR_OBJ {
		return printedOutput + evaluated.Inspect() + "\n"
	}

	if printedOutput == "" && evaluated != nil && evaluated.Type() != object.NULL_OBJ {
		return evaluated.Inspect()
	}

	return printedOutput
}

func main() {
	c := make(chan struct{})
	fmt.Println("Lontara-Lang WebAssembly Engine Ter-muat Sempurna.")
	js.Global().Set("evaluateLontara", js.FuncOf(evaluateLontaraWrapper))
	<-c
}
