# Alur Kerja & Pemrosesan Data (Execution Workflow)

Dokumen ini menjelaskan alur pemrosesan data dan tahap-tahap eksekusi kode pada `lontara-lang` dari teks kode sumber hingga luaran eksekusi.

---

## 🔄 Bagan Alur Kerja Sistem (Pipeline)

```text
[ Berkas Kode (.bugis / .lontara) ]
                 │
                 ▼
       ┌───────────────────┐
       │ 1. Lexical Analysis│  (Analisis UTF-8 & Pemecahan Token Leksikal)
       └─────────┬─────────┘
                 │
                 ▼
       ┌───────────────────┐
       │ 2. Pratt Parsing  │  (Penguraian Sintaksis ke Pohon AST)
       └─────────┬─────────┘
                 │
                 ▼
       ┌───────────────────┐
       │ 3. AST Tree Walk  │  (Penelusuran Pohon Sintaks Terstruktur)
       └─────────┬─────────┘
                 │
                 ▼
       ┌───────────────────┐
       │ 4. Environment    │  (Evaluasi Dinamis & Manajemen Skop Variabel)
       └─────────┬─────────┘
                 │
                 ▼
       [ Hasil Eksekusi / Output Konsol ]
```

---

## 📌 Rincian Tahap Pemrosesan

### 1. Lexical Analysis (`pkg/lexer`)
- Membaca teks masukan UTF-8 bita per bita/rune.
- Mendeteksi token pemisah, operator, angka, string, serta kata kunci Bugis Latin maupun Aksara Lontara Unicode (`U+1A00`–`U+1A1F`).
- Mencatat informasi baris (*line*) dan kolom (*column*) pada setiap token.

### 2. Syntactic Parsing (`pkg/parser`)
- Menggunakan metode **Pratt Parser** (*Top-Down Operator Precedence*) untuk mengevaluasi hirarki prioritas operator.
- Mengubah urutan token menjadi simpul-simpul pohon sintaks terstruktur (*Abstract Syntax Tree* / AST pada `pkg/ast`).

### 3. Tree-Walking Evaluation (`pkg/evaluator`)
- Menelusuri simpul AST secara rekursif.
- Mengalokasikan variabel dan closure fungsi dalam *Lexical Environment* (`pkg/object`).
- Mengembalikan objek nilai runtime atau menangkap pesan kesalahan presisi dengan lokasi baris/kolom.

### 4. Standard Output & CLI / WASM Interface (`cmd/lontara` / `cmd/wasm`)
- Menampilkan hasil cetakan `paui` ke terminal stdout atau meneruskannya ke peramban web melalui modul WebAssembly.
