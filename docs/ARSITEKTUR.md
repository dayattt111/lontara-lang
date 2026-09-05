# Arsitektur Lontara-Lang

`lontara-lang` diimplementasikan sebagai *tree-walking interpreter* menggunakan Go (Golang) murni tanpa pustaka eksternal. Interpreter ini memproses kode sumber melalui tahapan tradisional pemrosesan bahasa pemrograman: Lexical Analysis, Parsing, Abstract Syntax Tree (AST), dan Evaluation.

## Diagram Alur Eksekusi

```text
[ Kode Sumber (.lontara / .bugis) ]
               │
               ▼
       ┌───────────────┐
       │     Lexer     │  (Memecah string input menjadi token-token UTF-8)
       └───────┬───────┘
               │
               ▼
       ┌───────────────┐
       │    Parser     │  (Pratt Parser: Membangun struktur pohon AST)
       └───────┬───────┘
               │
               ▼
       ┌───────────────┐
       │      AST      │  (Representasi pohon ekspresi & statement)
       └───────┬───────┘
               │
               ▼
       ┌───────────────┐
       │   Evaluator   │  (Menelusuri AST & mengembalikan Object runtime)
       └───────────────┘
```

## Komponen Komponen Utama

### 1. Lexer (`pkg/lexer`)
Lexer bertugas mengubah stream karakter teks masukan menjadi token numerik dan kontekstual.
- Menggunakan `unicode/utf8` untuk mengodekan dan mendeteksi titik kode Unicode secara aman.
- Mendukung karakter Latin dan rentang aksara Lontara (`U+1A00` sampai `U+1A1F`).
- Menelusuri nomor baris (`line`), kolom (`col`), posisi bita saat ini (`position`), dan posisi bita baca berikutnya (`readPosition`).

### 2. Parser (`pkg/parser`)
Parser mengubah urutan token dari Lexer menjadi pohon sintaks terstruktur (*Abstract Syntax Tree*).
- Mengimplementasikan pendekatan **Pratt Parser** (*Top-Down Operator Precedence*) untuk penanganan prioritas operator matematika dan pembanding.
- Menggunakan *Recursive Descent Parsing* untuk pernyataan tingkat atas seperti deklarasi `taroi` dan pengembalian `lisu`.

### 3. Pohon Sintaks Terstruktur (`pkg/ast`)
Mendefinisikan antarmuka `Node`, `Statement`, dan `Expression`.
- `Program`: Root node dari berkas yang memuat daftar pernyataan.
- `TaroiStatement`: Pernyataan deklarasi variabel.
- `LisuStatement`: Pernyataan nilai kembalian.
- `ExpressionStatement`: Pernyataan berbunga ekspresi.
- `RekkoExpression`, `JamagauLiteral`, `CallExpression`, dll.

### 4. Sistem Objek & Lingkungan (`pkg/object`)
Semua nilai pada runtime dibungkus dalam tipe data yang mengimplementasikan antarmuka `object.Object`.
- `Integer`: Membungkus nilai `int64`.
- `String`: Membungkus nilai `string`.
- `Boolean`: Singleton `TRUE` dan `FALSE`.
- `Null`: Singleton `NULL`.
- `Function`: Menyimpan parameter, badan AST, dan salinan lingkungan leksikal (*closure*).
- `Environment`: Menyimpan pemetaan nama variabel ke objek runtime. Mendukung struktur bertingkat (*enclosed environment*) untuk skop fungsi local.

### 5. Evaluator (`pkg/evaluator`)
Penelusur AST (*tree walker*) yang mengevaluasi setiap node AST secara rekursif berdasarkan status lingkungan (`Environment`).
- Menghasilkan nilai `object.Object`.
- Menangani fungsi bawaan seperti `paui` / `ᨄᨕᨘᨕᨗ`.
- Melakukan evaluasi sirkuit pendek untuk boolean dan penghentian beruntun saat menemukan `ReturnValue` atau `Error`.

### 6. REPL & CLI (`pkg/repl` & `cmd/lontara`)
- `pkg/repl`: Menyediakan antarmuka Read-Eval-Print Loop interaktif melalui terminal.
- `cmd/lontara`: Biner utama yang menangani pembacaan argumen baris perintah (`lontara run <file>` atau mode REPL tanpa argumen).
