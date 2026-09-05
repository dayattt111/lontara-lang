# Peta Jalan & Alur Kerja (Roadmap & Workflow)

Dokumen ini memuat alur pemrosesan sistem dan peta jalan fitur umum proyek `lontara-lang`.

---

## 📌 Peta Jalan Fitur Publik (Public Roadmap)

```text
[ Core Interpreter ]  -->  [ WebAssembly & Loop ]  -->  [ Tipe Data Koleksi ]  -->  [ Ekosistem & Tooling ]
```

---

## 🛠️ Modul Utama
- **[WORKFLOW.md](WORKFLOW.md)**: Penjelasan rincian alur kerja pemrosesan data (Lexer -> Parser -> AST -> Evaluator).

---

## 🚀 Fitur yang Tersedia saat Ini
- **Dual Syntax Support**: Menulis kode dalam Bugis Latin maupun Aksara Lontara Unicode (`U+1A00`–`U+1A1F`).
- **Pratt Parser Engine**: Parsing ekspresi aritmetika, logika pembanding, variabel, dan percabangan (`rekko` / `sangadinna`).
- **Struktur Perulangan (`siki` / `ᨔᨗᨀᨗ`)**: Looping berbasis kondisi.
- **Pustaka Fungsi Bawaan**: `panjang` / `ᨄᨍ`, `baca` / `ᨅᨌ`, dan `tipe` / `ᨈᨗᨄᨙ`.
- **Pesan Kesalahan Presisi**: Menampilkan lokasi baris dan kolom presisi saat terjadi kesalahan sintaks atau runtime.
- **WebAssembly Engine & Playground**: Interpreter dapat dieksekusi di browser tanpa server backend.
