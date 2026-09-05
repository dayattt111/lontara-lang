# Dokumentasi Kamus & Aksara Lontara (`pkg/dictionary`)

Proyek `lontara-lang` menyediakan modul terisolasi khusus di direktori **[`pkg/dictionary`](../pkg/dictionary)** untuk mengelola tabel karakter Aksara Lontara Unicode dan kamus kata kunci sintaksis pemrograman.

---

## 1. Struktur Pustaka Kamus (`pkg/dictionary`)

```text
pkg/dictionary/
├── lontara.go        # Tabel huruf utama (Inang Sure'), diakritik vokal (Ana' Sure'), & tanda baca
├── keywords.go       # Registri pemetaan kata kunci Bugis Latin & Aksara Lontara
└── dictionary_test.go# Unit test verifikasi pencarian kata kunci & huruf Lontara
```

---

## 2. Tabel Aksara Lontara (`lontara.go`)

Seluruh huruf dan karakter Aksara Lontara pada rentang Unicode `U+1A00` sampai `U+1A1F` terdaftar di slice `LontaraAlphabet` pada `pkg/dictionary/lontara.go`:

### A. Inang Sure' (Huruf Utama / Suku Kata Konsonan)
| Rune | Hex | Nama | Latin | Kategori |
| :---: | :---: | :--- | :--- | :--- |
| `ᨀ` | `U+1A00` | Ka | ka | Inang Sure' |
| `ᨁ` | `U+1A01` | Ga | ga | Inang Sure' |
| `ᨂ` | `U+1A02` | Nga | nga | Inang Sure' |
| `ᨃ` | `U+1A03` | Ngka | ngka | Inang Sure' |
| `ᨄ` | `U+1A04` | Pa | pa | Inang Sure' |
| `ᨅ` | `U+1A05` | Ba | ba | Inang Sure' |
| `ᨆ` | `U+1A06` | Ma | ma | Inang Sure' |
| `ᨇ` | `U+1A07` | Mpa | mpa | Inang Sure' |
| `ᨈ` | `U+1A08` | Ta | ta | Inang Sure' |
| `ᨉ` | `U+1A09` | Da | da | Inang Sure' |
| `ᨊ` | `U+1A0A` | Na | na | Inang Sure' |
| `ᨋ` | `U+1A0B` | Nra | nra | Inang Sure' |
| `ᨌ` | `U+1A0C` | Ca | ca | Inang Sure' |
| `ᨍ` | `U+1A0D` | Ja | ja | Inang Sure' |
| `ᨎ` | `U+1A0E` | Nya | nya | Inang Sure' |
| `ᨏ` | `U+1A0F` | Nca | nca | Inang Sure' |
| `ᨐ` | `U+1A10` | Ya | ya | Inang Sure' |
| `ᨑ` | `U+1A11` | Ra | ra | Inang Sure' |
| `ᨒ` | `U+1A12` | La | la | Inang Sure' |
| `ᨓ` | `U+1A13` | Wa | wa | Inang Sure' |
| `ᨔ` | `U+1A14` | Sa | sa | Inang Sure' |
| `ᨕ` | `U+1A15` | A | a | Inang Sure' |
| `ᨖ` | `U+1A16` | Ha | ha | Inang Sure' |

### B. Ana' Sure' (Diakritik Vokal Tambahan)
| Rune | Hex | Nama | Vokal | Posisi |
| :---: | :---: | :--- | :--- | :--- |
| `ᨗ` | `U+1A17` | Tetteng | i | Atas huruf |
| `ᨘ` | `U+1A18` | Pucu' | u | Bawah huruf |
| `ᨙ` | `U+1A19` | Kelling | e | Kiri huruf |
| `ᨚ` | `U+1A1A` | Doping | o | Kanan huruf |
| `ᨛ` | `U+1A1B` | Kecce' | ae / pepet | Atas huruf |

### C. Tanda Baca
| Rune | Hex | Nama | Fungsi |
| :---: | :---: | :--- | :--- |
| `᨞` | `U+1A1E` | Pallawa | Pemisah kalimat (koma) |
| `᨟` | `U+1A1F` | End of Section | Penutup pasal/paragraf (titik) |

---

## 3. Panduan Kontribusi Kamus (`pkg/dictionary`)

Bagi kontributor yang ingin membantu memperluas kamus kata kunci atau karakter Aksara Lontara:

### Scenario A: Menambah Kata Kunci Pemrograman Baru
1. Buka berkas **`pkg/dictionary/keywords.go`**.
2. Tambahkan struct `Keyword` baru pada slice `RegisteredKeywords`:

```go
var RegisteredKeywords = []Keyword{
    // ...
    {
        TokenType: "SIKI",        // Tipe token baru
        Latin:     "siki",        // Kata kunci dalam Bugis Latin
        Lontara:   "ᨔᨗᨀᨗ",        // Kata kunci dalam Aksara Lontara
        Meaning:   "Perulangan (for / while)",
    },
}
```

3. Daftarkan konstanta token baru di `pkg/token/token.go`:
```go
const SIKI = "SIKI"
```

4. Uji perubahan Anda dengan pengujian otomatis:
```bash
go test -v ./pkg/dictionary/... ./pkg/token/...
```

### Scenario B: Menambah Karakter atau Metadata Huruf Lontara
1. Buka berkas **`pkg/dictionary/lontara.go`**.
2. Tambahkan entri karakter baru pada slice `LontaraAlphabet`:

```go
{
    Rune:        'ᨀ',
    Hex:         "U+1A00",
    Name:        "Ka",
    Latin:       "ka",
    Category:    "Inang Sure'",
    Description: "Konsonan Ka",
},
```

3. Jalankan pengujian:
```bash
go test -v ./pkg/dictionary/...
```

Modul `pkg/token` secara otomatis menggunakan pencarian dari `pkg/dictionary.LookupKeyword()` sehingga Anda **tidak perlu** mengedit logika pencarian token secara manual di file lexer/token lain.
