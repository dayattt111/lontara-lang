# Belajar Lontara-Lang: 5. Fungsi (`jamagau` / `lisu`)

Fungsi dalam `lontara-lang` bersifat *first-class citizen*. Fungsi dideklarasikan sebagai ekspresi dengan kata kunci `jamagau` (func) dan nilai kembali menggunakan `lisu` (return).

---

## Sintaksis

### Versi Bugis Latin
```lontara
taroi tambah = jamagau(a, b) {
    lisu a + b;
};

taroi hasil = tambah(10, 20);
paui("Hasil penjumlahan:", hasil);
```

### Versi Aksara Lontara
```lontara
ᨈᨑᨚᨕᨗ kali = ᨍᨆᨁᨕᨘ(x, y) {
    ᨒᨗᨔᨘ x * y;
};

taroi res = kali(4, 5);
paui(res);
```

---

## Closure & Lexical Scope

Fungsi dalam `lontara-lang` mendukung *closure*, sehingga fungsi dapat mengingat variabel dari lingkup (*environment*) di mana fungsi tersebut diciptakan.
