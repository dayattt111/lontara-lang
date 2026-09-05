# Pratt Parser & Abstract Syntax Tree (AST)

Modul parser pada `pkg/parser` bertugas mengonversi urutan token leksikal menjadi Abstract Syntax Tree (AST) terstruktur pada `pkg/ast`.

---

## Metodologi Pratt Parsing

Parser mengimplementasikan **Pratt Parser** (*Top-Down Operator Precedence*) untuk menangani hierarki prioritas operator:

| Precedence Constant | Level | Token / Operator |
| :--- | :---: | :--- |
| `LOWEST` | 1 | Token umum |
| `EQUALS` | 2 | `==`, `!=` |
| `LESSGREATER` | 3 | `<`, `>`, `<=`, `>=` |
| `SUM` | 4 | `+`, `-` |
| `PRODUCT` | 5 | `*`, `/` |
| `PREFIX` | 6 | `-X`, `!X` |
| `CALL` | 7 | `fungsi(X)` |

---

## Simpul AST Utama (`pkg/ast`)

- `TaroiStatement`: Pernyataan deklarasi `taroi`.
- `LisuStatement`: Pernyataan return `lisu`.
- `RekkoExpression`: Ekspresi percabangan `rekko` ... `sangadinna`.
- `SikiExpression`: Ekspresi perulangan `siki` (Fase 2).
- `JamagauLiteral`: Definisi fungsi `jamagau`.
- `CallExpression`: Pemanggilan fungsi.
