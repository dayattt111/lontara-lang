# Panduan Kontribusi Lontara-Lang

Terima kasih telah tertarik untuk berkontribusi pada pengembangan `lontara-lang`! Proyek ini terbuka bagi siapa saja yang ingin membantu memperluas ekosistem, memperbaiki bug, menyempurnakan sintaks, atau menambah dokumentasi.

## 1. Alur Kerja Kontribusi (Workflow)

1. **Fork Repositori**
   Buat fork dari repositori [lontara-lang](https://github.com/dayattt111/lontara-lang) ke akun GitHub Anda.

2. **Clone Repositori Lokal**
   ```bash
   git clone https://github.com/<username-anda>/lontara-lang.git
   cd lontara-lang
   ```

3. **Buat Branch Baru**
   Gunakan nama branch yang deskriptif:
   ```bash
   git checkout -b fitur/nama-fitur
   # atau
   git checkout -b perbaikan/deskripsi-bug
   ```

4. **Lakukan Perubahan Kode & Pengujian**
   Pastikan kode yang diubah telah lolos seluruh unit test tanpa error:
   ```bash
   go test -v ./...
   ```

5. **Commit & Push**
   ```bash
   git commit -m "feat: tambah dukungan operator X"
   git push origin fitur/nama-fitur
   ```

6. **Buka Pull Request (PR)**
   Kirimkan Pull Request ke branch `main` repositori utama dengan deskripsi perubahan yang jelas.

---

## 2. Struktur Kode untuk Pengembang

- `pkg/lexer`: Penanganan pemecahan bita/rune UTF-8 menjadi token.
- `pkg/parser`: Implementasi Pratt Parser untuk mengubah token menjadi AST.
- `pkg/ast`: Node definisi pohon sintaks terstruktur.
- `pkg/evaluator`: Penelusur AST dan logika eksekusi runtime.
- `pkg/object`: Tipe data runtime, nilai singleton (`TRUE`, `FALSE`, `NULL`), dan *environment*.
- `pkg/repl`: Antarmuka REPL interaktif terminal.

---

## 3. Standar Pengodean & Pengujian

- Gunakan format standar Go dengan menjalankan `go fmt ./...`.
- Setiap penambahan fitur atau perbaikan bug pada lexer, parser, atau evaluator wajib menyertakan unit test di berkas `*_test.go` terkait.
- Semua pengujian harus lulus saat menjalankan `go test -count=1 ./...`.
