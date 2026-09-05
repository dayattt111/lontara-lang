# Peta Jalan Pengembangan (Roadmap) Lontara-Lang

Dokumen ini memuat rencana pengembangan jangka pendek, menengah, dan panjang untuk `lontara-lang`.

---

## 📌 Status Rilis & Tahapan (Milestones)

```text
[v0.1.0 - Core Interpreter]  -->  [v0.2.0 - WASM & Loop]  -->  [v0.3.0 - Koleksi Data]  -->  [v1.0.0 - Ekosistem & Tooling]
         (Selesai)                      (Selesai)                    (Dalam Proses)                 (Tujuan Akhir)
```

---

## ✅ Fase 1: Fondasi Interpreter & Dual Sintaks (v0.1.0 - Selesai)

- [x] **Lexer UTF-8**: Penanganan karakter Latin dan Aksara Lontara Unicode (`U+1A00`–`U+1A1F`).
- [x] **Parser (Pratt Parser)**: Parsing ekspresi aritmetika, pembanding, dan kondisi.
- [x] **Registri Kamus Kata Kunci (`pkg/dictionary`)**: Dukungan simultan sintaksis Bugis Latin dan Aksara Lontara.
- [x] **Evaluator & Objek Runtime**: Pengalokasian variabel (`taroi`), fungsi (*first-class closures* `jamagau`), percabangan (`rekko` / `sangadinna`), dan pengembalian nilai (`lisu`).
- [x] **CLI & REPL (`cmd/lontara`)**: Eksekutor berkas `.bugis` / `.lontara` dan mode REPL interaktif terminal.
- [x] **Dokumentasi Dasar**: Spesifikasi sintaksis, arsitektur, panduan penggunaan, dan panduan kontribusi.

---

## ✅ Fase 2: Mesin WebAssembly & Fitur Sintaksis Tambahan (v0.2.0 - Selesai)

- [x] **Mesin WebAssembly (WASM)**:
  - Kompilasi biner `cmd/wasm` ke `lontara.wasm` agar interpreter dapat berjalan langsung di peramban web (*zero backend server execution*).
  - Pembuatan Web Playground interaktif di browser ([examples/playground.html](../examples/playground.html)).
- [x] **Struktur Perulangan (Looping)**:
  - Penambahan kata kunci perulangan Bugis Latin & Aksara Lontara (`siki` / `ᨔᨗᨀᨗ`).
- [x] **Fungsi Bawaan (Built-in) Tambahan**:
  - Penambahan fungsi pembacaan input konsol (`baca` / `ᨅᨌ`), pengukur panjang string (`panjang` / `ᨄᨍ`), dan cek tipe data (`tipe` / `ᨈᨗᨄᨙ`).
- [x] **Penyempurnaan Pesan Kesalahan (Error Reporting)**:
  - Menampilkan posisi baris dan kolom yang presisi saat terjadi kesalahan sintaks (*parser error*) maupun kesalahan eksekusi (*runtime error*).

---

## 💡 Fase 3: Tipe Data Koleksi & Modul standar (v0.3.0 - Rencana)

- [ ] **Tipe Data Array / List**:
  - Sintaksis array literal (contoh: `[1, 2, 3]`).
  - Operasi elemen array dan iterasi.
- [ ] **Tipe Data Peta (Hash Map / Dictionary)**:
  - Pasangan kunci-nilai (contoh: `{"kunci": "nilai"}`).
- [ ] **Penyatuan Sistem Modul / Impor Berkas**:
  - Memungkinkan pengimporan berkas `.bugis` atau `.lontara` lain dalam satu proyek.

---

## 🎨 Fase 4: Ekosistem & Alat Pengembang (v1.0.0 Target)

- [ ] **Ekstensi VS Code / Editor Integration**:
  - *Syntax Highlighting* dan *Auto-completion* untuk berkas `.bugis` dan `.lontara` di VS Code.
  - Ikon khusus untuk berkas `.bugis` / `.lontara`.
- [ ] **Pemberitahuan & Usulan Komunitas Icon Theme**:
  - Mengajukan integrasi ikon resmi Lontara-Lang pada ekosistem editor (*vscode-material-icon-theme*).
- [ ] **Paket Biner Distribusi**:
  - Rilis biner CLI mandiri untuk Linux, macOS, dan Windows via GitHub Releases dan Homebrew/Go install.

---

## 🤝 Masukan & Usulan Fitur

Jika Anda memiliki ide kata kunci Bugis baru, usulan sintaksis, atau ingin berkontribusi pada salah satu item di roadmap di atas, silakan buka issue atau kirimkan Pull Request sesuai dengan **[docs/KONTRIBUSI.md](KONTRIBUSI.md)**.
