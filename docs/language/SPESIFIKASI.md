# Spesifikasi Bahasa Lontara-Lang

`lontara-lang` adalah bahasa pemrograman berjenis *tree-walking interpreter* berfitur *dual-syntax* (Bugis Latin dan Aksara Lontara Unicode `U+1A00`–`U+1A1F`).

---

## Fitur Utama Language Engine

1. **Dual Syntax Native**: Kode sumber dapat ditulis dalam format `.bugis` (Bugis Latin) atau `.lontara` (Aksara Lontara Unicode).
2. **First-Class Functions**: Fungsi dideklarasikan sebagai ekspresi `jamagau` dan mendukung *lexical closures*.
3. **Dynamic Typing**: Tipe data ditentukan secara dinamis pada saat runtime (`INTEGER`, `STRING`, `BOOLEAN`, `NULL`, `FUNCTION`, `BUILTIN`).
4. **Presisi Error Reporting**: Menyampaikan informasi lokasi baris dan kolom yang presisi saat terjadi kesalahan sintaksis atau kesalahan runtime.
