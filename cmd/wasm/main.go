package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
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

	// Tangkap luaran stdout (paui) sementara
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		w.Close()
		os.Stdout = oldStdout
		<-outC
		var errOut bytes.Buffer
		errOut.WriteString("Kesalahan Sintaks (Parser Error):\n")
		for _, msg := range p.Errors() {
			errOut.WriteString(fmt.Sprintf("  - %s\n", msg))
		}
		return errOut.String()
	}

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)

	w.Close()
	os.Stdout = oldStdout
	printedOutput := <-outC

	if evaluated != nil && evaluated.Type() == object.ERROR_OBJ {
		return printedOutput + evaluated.Inspect() + "\n"
	}

	return printedOutput
}

func main() {
	c := make(chan struct{})
	fmt.Println("Lontara-Lang WebAssembly Engine Ter-muat Sempurna.")
	js.Global().Set("evaluateLontara", js.FuncOf(evaluateLontaraWrapper))
	<-c
}
