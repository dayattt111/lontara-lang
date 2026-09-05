# Panduan Sintaks Lontara-Lang

Dokumen ini berisi spesifikasi sintaks dan tata bahasa untuk `lontara-lang`. Bahasa ini mendukung dua mode penulisan: versi Bugis Latin dan versi Aksara Lontara (Unicode U+1A00–U+1A1F).

## 1. Kata Kunci (Keywords)

| Konsep Umum | Bugis Latin | Aksara Lontara | Deskripsi |
| :--- | :--- | :--- | :--- |
| `let / var` | `taroi` | `ᨈᨑᨚᨕᨗ` | Deklarasi variabel |
| `func` | `jamagau` | `ᨍᨆᨁᨕᨘ` | Deklarasi fungsi anonim / ekspresi |
| `if` | `rekko` | `ᨑᨙᨀᨚ` | Evaluasi kondisi percabangan |
| `else` | `sangadinna` | `ᨔᨂᨉᨗᨊ` | Blok alternatif percabangan |
| `true` | `tongeng` | `ᨈᨚᨂᨙ` | Nilai kebenaran true |
| `false` | `banna` | `ᨅᨊ` | Nilai kebenaran false |
| `return` | `lisu` | `ᨒᨗᨔᨘ` | Mengembalikan nilai dari fungsi |
| `print` | `paui` | `ᨄᨕᨘᨕᨗ` | Cetak nilai ke konsol |

## 2. Tipe Data

### Integer
Memuat nilai bilangan bulat positif dan negatif.
```lontara
taroi x = 10;
taroi y = -5;
```

### String
Diapit dengan tanda petik ganda (`"`).
```lontara
taroi nama = "Lontara";
```

### Boolean
Literal boolean menggunakan kata kunci `tongeng` (true) dan `banna` (false).
```lontara
taroi benar = tongeng;
taroi salah = banna;
```

### Null
Dipakai secara internal jika suatu ekspresi tidak menghasilkan nilai.

## 3. Variabel

Pengalokasian variabel menggunakan kata kunci `taroi` diikutsertakan operator penetapan `=`.

```lontara
taroi angka = 42;
taroi pesan = "Salam dari Sulawesi";
```

Dalam Aksara Lontara:
```lontara
ᨈᨑᨚᨕᨗ angka = 42;
```

## 4. Percabangan (`rekko` ... `sangadinna`)

Percabangan mengevaluasi ekspresi boolean dalam tanda kurung `(...)` diikuti oleh blok kode `{...}`.

```lontara
rekko (angka > 20) {
    paui("Angka lebih besar dari 20");
} sangadinna {
    paui("Angka kecil");
}
```

Dalam Aksara Lontara:
```lontara
ᨑᨙᨀᨚ (angka > 20) {
    ᨄᨕᨘᨕᨗ("Angka lebih besar");
}
```

## 5. Fungsi (`jamagau`)

Fungsi dalam `lontara-lang` bersifat first-class value. Fungsi dideklarasikan sebagai ekspresi dengan kata kunci `jamagau`, parameter dalam tanda kurung, dan badan fungsi di dalam kurung kurawal. Untuk mengembalikan nilai, gunakan `lisu`.

```lontara
taroi kali = jamagau(a, b) {
    lisu a * b;
};

taroi hasil = kali(6, 7);
paui(hasil);
```

## 6. Fungsi Bawaan (Built-in)

### `paui` / `ᨄᨕᨘᨕᨗ`
Menampilkan nilai argumen ke stdout.

```lontara
paui("Halo dunia");
paui(100 + 200);
```
