# Arsitektur Interpreter Lontara-Lang

`lontara-lang` diimplementasikan sebagai *tree-walking interpreter* menggunakan Go murni tanpa pustaka eksternal.

---

## Diagram Pipeline Pemrosesan

```text
[ Teks Sumber (.bugis / .lontara) ]
                │
                ▼
        ┌───────────────┐
        │     Lexer     │  (Memecah bita/rune UTF-8 menjadi urutan Token)
        └───────┬───────┘
                │
                ▼
        ┌───────────────┐
        │    Parser     │  (Pratt Parser: Membangun pohon AST terstruktur)
        └───────┬───────┘
                │
                ▼
        ┌───────────────┐
        │      AST      │  (Simpul pohon Statement & Expression)
        └───────┬───────┘
                │
                ▼
        ┌───────────────┐
        │   Evaluator   │  (Penelusur AST & pengembali Objek Runtime)
        └───────────────┘
```

---

## Modul Terkait
- [PARSER.md](PARSER.md): Penjelasan detail rancangan Pratt Parser & AST.
- [EVALUATOR.md](EVALUATOR.md): Logika evaluasi runtime, closures, & error reporting.
