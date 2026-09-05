# Belajar Lontara-Lang: 6. Fungsi Bawaan (Built-in Functions)

`lontara-lang` menyediakan beberapa fungsi bawaan sistem untuk tugas-tugas umum.

---

## 1. `panjang` / `ᨄᨍ`
Mengembalikan jumlah karakter (UTF-8 rune) dari suatu string.

```lontara
taroi n = panjang("Salama Lontara");
paui(n); // Output: 14
```

---

## 2. `baca` / `ᨅᨌ`
Membaca satu baris masukan teks dari pengguna via terminal (stdin).

```lontara
paui("Masukkan nama Anda:");
taroi nama = baca();
paui("Halo,", nama);
```

---

## 3. `tipe` / `ᨈᨗᨄᨙ`
Mengembalikan nama jenis tipe data suatu objek (`INTEGER`, `STRING`, `BOOLEAN`, dll.).

```lontara
paui(tipe(100));    // Output: INTEGER
paui(tipe("teks"));  // Output: STRING
```
