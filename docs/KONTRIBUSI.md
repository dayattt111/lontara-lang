# Panduan Kontribusi Lontara-Lang

Terima kasih telah tertarik untuk berkontribusi pada pengembangan `lontara-lang`! Proyek ini terbuka bagi siapa saja yang ingin membantu memperluas ekosistem, memperbaiki bug, menyempurnakan sintaks, menambah kamus kata kunci, atau melengkapi dokumentasi.

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
   git commit -m "feat: tambah kata kunci perulangan siki"
   git push origin fitur/nama-fitur
   ```

6. **Buka Pull Request (PR)**
   Kirimkan Pull Request ke branch `main` repositori utama dengan deskripsi perubahan yang jelas.

---

## 2. Kontribusi pada Kamus Bahasa & Aksara (`pkg/dictionary`)

Jika Anda ingin berkontribusi menambah kata kunci Bugis/Lontara baru atau melengkapi tabel huruf Aksara Lontara:
- Pelajari panduan terpisah di **[docs/KAMUS.md](KAMUS.md)**.
- Edit berkas **`pkg/dictionary/keywords.go`** untuk menambah kata kunci baru.
- Edit berkas **`pkg/dictionary/lontara.go`** untuk melengkapi tabel huruf Lontara.

---

## 3. Struktur Kode untuk Pengembang

- `pkg/dictionary`: Modul terisolasi registri kamus kata kunci & tabel karakter Lontara.
- `pkg/token`: Tipe token dan pencarian kata kunci dari modul dictionary.
- `pkg/lexer`: Penanganan pemecahan bita/rune UTF-8 menjadi token.
- `pkg/parser`: Implementasi Pratt Parser untuk mengubah token menjadi AST.
- `pkg/ast`: Node definisi pohon sintaks terstruktur.
- `pkg/evaluator`: Penelusur AST dan logika eksekusi runtime.
- `pkg/object`: Tipe data runtime, nilai singleton (`TRUE`, `FALSE`, `NULL`), dan *environment*.
- `pkg/repl`: Antarmuka REPL interaktif terminal.

---

## 4. Standar Pengodean & Pengujian

- Gunakan format standar Go dengan menjalankan `go fmt ./...`.
- Setiap penambahan fitur, kata kunci, atau perbaikan bug wajib menyertakan unit test di berkas `*_test.go` terkait.
- Semua pengujian harus lulus saat menjalankan `go test -count=1 ./...`.


*Pengen deh logo lontara-lang bisa masuk di vscode-material-icon-theme.