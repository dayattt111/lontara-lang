# Lontara Lang

```
lontara-lang/
├── .github/
│   └── workflows/
│       └── ci.yml               # Automated test & build saat push/PR
├── cmd/
│   ├── lontara/
│   │   └── main.go              # Entry point CLI (REPL & file runner)
│   └── wasm/
│       └── main.go              # Entry point khusus WebAssembly (Fase 2)
├── pkg/
│   ├── token/
│   │   └── token.go             # Definisi tipe token & tabel kata kunci Bugis/Lontara
│   ├── lexer/
│   │   ├── lexer.go             # Pemecah teks mentah menjadi deretan token (UTF-8/Rune safe)
│   │   └── lexer_test.go        # Unit test lexer
│   ├── ast/
│   │   └── ast.go               # Definisi node-node Abstract Syntax Tree
│   ├── parser/
│   │   ├── parser.go            # Pengubah token menjadi pohon sintaks (AST)
│   │   └── parser_test.go       # Unit test parser
│   ├── object/
│   │   ├── object.go            # Sistem tipe runtime (Integer, Boolean, String, Error)
│   │   └── environment.go       # Penyimpan state variabel & scope (symbol table)
│   ├── evaluator/
│   │   ├── evaluator.go         # Eksekutor AST ke hasil komputasi
│   │   └── evaluator_test.go    # Unit test evaluator
│   └── repl/
│       └── repl.go              # Read-Eval-Print Loop untuk terminal interaktif
├── examples/
│   ├── halo_dunia.bugis         # Contoh program versi latin Bugis
│   └── halo_dunia.lontara       # Contoh program versi aksara Lontara
├── go.mod
├── LICENSE                      # Lisensi open source (misal: MIT atau Apache 2.0)
└── README.md                    # Dokumentasi utama, cara install, sintaks dasar
```
