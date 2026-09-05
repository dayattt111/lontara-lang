package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dayattt111/lontara-lang/pkg/evaluator"
	"github.com/dayattt111/lontara-lang/pkg/lexer"
	"github.com/dayattt111/lontara-lang/pkg/object"
	"github.com/dayattt111/lontara-lang/pkg/parser"
	"github.com/dayattt111/lontara-lang/pkg/repl"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Lontara Programming Language (v0.1.0)")
		fmt.Println("Ketik perintah atau ekspresi untuk mengevaluasi.")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	var filePath string
	if args[0] == "run" {
		if len(args) < 2 {
			fmt.Println("Penggunaan: lontara run <berkas.lontara|berkas.bugis>")
			os.Exit(1)
		}
		filePath = args[1]
	} else {
		filePath = args[0]
	}

	ext := filepath.Ext(filePath)
	if ext != ".lontara" && ext != ".bugis" {
		fmt.Printf("Peringatan: ekstensi berkas '%s' sebaiknya menggunakan .lontara atau .bugis\n", ext)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Gagal membaca berkas %s: %v\n", filePath, err)
		os.Exit(1)
	}

	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println("Kesalahan Sintaks (Parser Error):")
		for _, msg := range p.Errors() {
			fmt.Printf("  - %s\n", msg)
		}
		os.Exit(1)
	}

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil && evaluated.Type() == object.ERROR_OBJ {
		fmt.Printf("Kesalahan Eksekusi: %s\n", evaluated.Inspect())
		os.Exit(1)
	}
}
