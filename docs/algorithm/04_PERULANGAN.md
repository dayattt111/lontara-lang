# Belajar Lontara-Lang: 4. Perulangan (`siki` / `ᨔᨗᨀᨗ`)

Struktur kontrol perulangan (*while-like loop*) menggunakan kata kunci `siki` (Bugis Latin) atau `ᨔᨗᨀᨗ` (Aksara Lontara).

---

## Sintaksis

### Versi Bugis Latin
```lontara
taroi i = 1;
taroi total = 0;

siki (i <= 5) {
    paui("Iterasi baris ke-", i);
    taroi total = total + i;
    taroi i = i + 1;
}

paui("Total akumulasi:", total);
```

### Versi Aksara Lontara
```lontara
ᨈᨑᨚᨕᨗ i = 1;
ᨈᨑᨚᨕᨗ total = 0;

ᨔᨗᨀᨗ (i <= 5) {
    ᨄᨕᨘᨕᨗ("Iterasi baris ke-", i);
    ᨈᨑᨚᨕᨗ total = total + i;
    ᨈᨑᨚᨕᨗ i = i + 1;
}

ᨄᨕᨘᨕᨗ("Total akumulasi:", total);
```

---

## Alur Kerja

Perulangan `siki` mengevaluasi ekspresi kondisi sebelum setiap iterasi. Selama kondisi bernilai `tongeng` (true), blok kode di dalam kurung kurawal `{...}` akan terus dijalankan.
