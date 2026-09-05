# Tutorial Kontribusi Kamus & Aksara (`pkg/dictionary`)

Dokumen ini memuat langkah demi langkah cara menambah kata kunci baru atau memperbarui tabel huruf Lontara.

---

## 1. Menambah Kata Kunci Baru

1. Buka berkas `pkg/dictionary/keywords.go`.
2. Tambahkan entri struct `Keyword` pada `RegisteredKeywords`:

```go
{
    TokenType: "SIKI",
    Latin:     "siki",
    Lontara:   "ᨔᨗᨀᨗ",
    Meaning:   "Perulangan berbasis kondisi (while/loop)",
},
```

3. Tambahkan konstanta token di `pkg/token/token.go`:
```go
const SIKI = "SIKI"
```

4. Jalankan pengujian:
```bash
go test -v ./pkg/dictionary/... ./pkg/token/...
```

---

## 2. Menambah Karakter Lontara Baru

1. Buka berkas `pkg/dictionary/lontara.go`.
2. Tambahkan entri karakter pada slice `LontaraAlphabet`.
3. Uji perubahan: `go test -v ./pkg/dictionary/...`.
