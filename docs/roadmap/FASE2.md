# Dokumentasi Terstruktur Rilis Fase 2 (v0.2.0)

Dokumen ini menyajikan laporan rincian teknis pencapaian pengembangan **Fase 2** pada `lontara-lang`.

---

## 🚀 Ringkasan Rilis Fitur Fase 2

Fase 2 melengkapi kemampuan interpreter dengan fitur perulangan, penambahan pustaka fungsi bawaan dasar, penyempurnaan laporan kesalahan, serta portabilitas WebAssembly untuk peramban web.

---

## 🛠️ Rincian Fitur & Komponen

### 1. Struktur Perulangan (`siki` / `ᨔᨗᨀᨗ`)
- **Deskripsi**: Penambahan struktur kontrol perulangan berbasis kondisi (*while-like loop*).
- **Kata Kunci**: `siki` (Bugis Latin) dan `ᨔᨗᨀᨗ` (Aksara Lontara Unicode).
- **Modul Terkait**:
  - `pkg/dictionary/keywords.go` & `pkg/token/token.go`
  - `pkg/ast/ast.go` (`SikiExpression`)
  - `pkg/parser/parser.go` (`parseSikiExpression`)
  - `pkg/evaluator/evaluator.go` (`evalSikiExpression`)
- **Dokumentasi Algoritma**: [04_PERULANGAN.md](../algorithm/04_PERULANGAN.md)

---

### 2. Pustaka Fungsi Bawaan Baru
- **`panjang` / `ᨄᨍ`**: Mengembalikan jumlah karakter UTF-8 rune dari string.
- **`baca` / `ᨅᨌ`**: Membaca satu baris masukan teks dari stdin.
- **`tipe` / `ᨈᨗᨄᨙ`**: Mengembalikan jenis tipe data objek dalam bentuk String.
- **Modul Terkait**: `pkg/evaluator/evaluator.go` (`builtins` map).
- **Dokumentasi Algoritma**: [06_FUNGSI_BAWAAN.md](../algorithm/06_FUNGSI_BAWAAN.md)

---

### 3. Penyempurnaan Error Reporting Presisi
- **Deskripsi**: Menampilkan nomor baris (*line*) dan kolom (*column*) dari token AST secara presisi saat terjadi kesalahan sintaks atau runtime.
- **Format Pesan**:
  `Kasalahan: [Baris X, Kolom Y] <pesan kesalahan>`
- **Modul Terkait**: `pkg/evaluator/evaluator.go` (`newError`) & `cmd/lontara/main.go`.

---

### 4. Mesin WebAssembly (WASM) & Web Playground
- **Deskripsi**: Ekspor biner WASM `lontara.wasm` dari `cmd/wasm` via `syscall/js` agar interpreter dapat dieksekusi di peramban web tanpa server backend.
- **Berkas Terkait**:
  - `cmd/wasm/main.go`
  - `examples/playground.html` (Interactive Web Playground UI)
- **Dokumentasi Web**: [WASM.md](../website/WASM.md)

---

## 🧪 Verifikasi & Pengujian Rilis

- **Unit Tests**: Pass 100% pada `go test ./pkg/...`.
- **WASM Build**: Pass 100% pada `GOOS=js GOARCH=wasm go build ./cmd/wasm`.
