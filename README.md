# lontara-lang

<p align="center">
    <img src="assets/lontara.svg" alt="Lontara-Lang Logo By Muhammad Amin Hidayat(Gen AI)" width="400">
</p>

<p align="center">
    <img src="https://img.shields.io/badge/Go-1.21%2B-blue?style=flat-square" alt="Go Version">
    <img src="https://img.shields.io/badge/License-MIT-green?style=flat-square" alt="License">
    <img src="https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-orange?style=flat-square" alt="Platform">
</p>

**Bahasa Pemrograman Edukatif Berbasis Logika Bahasa Bugis dan Aksara Lontara.**

`lontara-lang` adalah bahasa pemrograman berjenis *tree-walking interpreter* yang dirancang untuk mengenalkan konsep ilmu komputer dan algoritma pemrograman melalui istilah lokal bahasa Bugis serta dukungan native Aksara Lontara Unicode (`U+1A00`–`U+1A1F`). Bahasa ini dikembangkan menggunakan Go (Golang) murni tanpa pustaka eksternal (*zero external dependencies*).

---

## Inspirasi Proyek

Pengembangan `lontara-lang` terinspirasi dari proyek-proyek esolang edukatif Nusantara pendahulu yang mengenalkan bahasa pemrograman berbasis budaya daerah:

- **[sunda-lang](https://github.com/randspace0/sunda-lang)** – Bahasa pemrograman berbasis sintaks bahasa Sunda.
- **JawaScript** – Bahasa pemrograman berbasis sintaks bahasa Jawa.

Apresiasi setinggi-tingginya untuk para pembuat proyek tersebut atas inspirasi pelestarian budaya digital di Nusantara.

---

## Uji Coba Pertama (Hello World)

Anda dapat menulis program `lontara-lang` dalam format `.bugis` (versi Latin) atau `.lontara` (versi Aksara Lontara).

### Versi Bugis Latin (`examples/halo_dunia.bugis`)

```lontara
taroi pesang = "Halo, Dunia!";
paui(pesang);

taroi lima = 5;
taroi sepulo = 10;

taroi tambah = jamagau(x, y) {
    lisu x + y;
};

taroi hasil = tambah(lima, sepulo);
rekko (hasil >= 15) {
    paui("Hasil pencumlahan sitinaja (>= 15)");
} sangadinna {
    paui("Hasil pencumlahan kurang");
}
```

### Versi Aksara Lontara (`examples/halo_dunia.lontara`)

```lontara
ᨈᨑᨚᨕᨗ pesang = "Halo, Dunia!";
ᨄᨕᨘᨕᨗ(pesang);

ᨈᨑᨚᨕᨗ angka = 10;
ᨑᨙᨀᨚ (angka > 5) {
    ᨄᨕᨘᨕᨗ("salama");
}
```

---

## Sekilas Kata Kunci

| Konsep Umum | Bugis Latin | Aksara Lontara | Deskripsi |
| :--- | :--- | :--- | :--- |
| `let / var` | `taroi` | `ᨈᨑᨚᨕᨗ` | Deklarasi variabel |
| `func` | `jamagau` | `ᨍᨆᨁᨕᨘ` | Deklarasi fungsi |
| `if` | `rekko` | `ᨑᨙᨀᨚ` | Percabangan kondisi |
| `else` | `sangadinna` | `ᨔᨂᨉᨗᨊ` | Blok alternatif percabangan |
| `true` / `false` | `tongeng` / `banna` | `ᨈᨚᨂᨙ` / `ᨅᨊ` | Nilai boolean |
| `return` | `lisu` | `ᨒᨗᨔᨘ` | Pengembalian nilai dari fungsi |
| `print` | `paui` | `ᨄᨕᨘᨕᨗ` | Cetak nilai ke konsol |

Panduan sintaksis dan tata bahasa lengkap dapat dibaca di **[docs/SINTAKS.md](docs/SINTAKS.md)**.

---

## Instalasi dan Setup

### Memasang Biner CLI (`lontara`)

- **Pengguna Umum (via Go Install):**
  ```bash
  go install github.com/dayattt111/lontara-lang/cmd/lontara@latest
  ```

- **Pengembang Lokal:**
  ```bash
  go install ./cmd/lontara
  ```

*Catatan: Jika muncul error `lontara: command not found`, pastikan `$GOPATH/bin` terdaftar dalam variabel `PATH` Anda (`export PATH=$PATH:$(go env GOPATH)/bin`).*

### Eksekusi Berkas Kode

```bash
# Eksekusi langsung berkas
lontara examples/halo_dunia.bugis
lontara examples/halo_dunia.lontara

# Atau menggunakan sub-perintah run
lontara run examples/halo_dunia.bugis

# Atau menggunakan go run (tanpa install)
go run ./cmd/lontara examples/halo_dunia.bugis
```

### Mode Interaktif (REPL)

```bash
$ lontara
Lontara Programming Language (v0.1.0)
Ketik perintah atau ekspresi untuk mengevaluasi.
lontara> taroi a = 10;
lontara> paui(a * 2);
20
```

Panduan lengkap mengenai setup dan penanganan masalah dapat dibaca di **[docs/PENGGUNAAN.md](docs/PENGGUNAAN.md)**.

---

## Struktur Proyek

```text
.
├── LICENSE                  # Lisensi proyek (MIT License)
├── README.md                # Dokumentasi utama repositori
├── assets/                  # Berkas aset visual dan logo proyek
├── cmd/
│   ├── lontara/             # Biner CLI runner dan REPL utama (main.go)
│   └── wasm/                # Biner kompilasi WebAssembly (main.go)
├── docs/
│   ├── ARSITEKTUR.md        # Dokumentasi arsitektur mesin interpreter
│   ├── KAMUS.md             # Dokumentasi tabel huruf Lontara & registri kata kunci
│   ├── KONTRIBUSI.md        # Panduan alur kerja untuk kontributor
│   ├── PENGGUNAAN.md        # Panduan detail instalasi, setup PATH, dan CLI
│   ├── ROADMAP.md           # Peta jalan pengembangan jangka pendek & panjang
│   ├── SINTAKS.md           # Spesifikasi sintaksis dan manual kata kunci
│   └── WASM.md              # Panduan arsitektur & kompilasi WebAssembly
├── examples/
│   ├── halo_dunia.bugis     # Contoh berkas sintaks Bugis Latin
│   └── halo_dunia.lontara   # Contoh berkas sintaks Aksara Lontara Unicode
├── go.mod                   # Berkas modul Go
└── pkg/
    ├── ast/                 # Node Abstract Syntax Tree (AST)
    ├── dictionary/          # Modul terisolasi kamus Aksara Lontara & pemetaan kata kunci
    ├── evaluator/           # Mesin penelusur AST dan evaluasi nilai
    ├── lexer/               # Analisis leksikal dan pembacaan token UTF-8
    ├── object/              # Tipe data runtime dan lingkup variabel (Environment)
    ├── parser/              # Parser sintaksis berbasis Pratt Parser
    ├── repl/                # Antarmuka REPL terminal interaktif
    └── token/               # Definisi token dan pencarian kata kunci
```

---

## Arsitektur Interpreter Singkat

`lontara-lang` menggunakan alur pemrosesan interpreter tradisional tanpa dependensi luar:

```text
Teks Sumber (.bugis/.lontara) ──> Lexer (Token) ──> Parser (AST) ──> Evaluator (Object Value)
```

1. **Lexer:** Memecah teks masukan UTF-8 (termasuk Aksara Lontara) menjadi urutan token.
2. **Parser:** Mengubah token menjadi struktur pohon AST menggunakan metode Pratt Parser.
3. **Evaluator:** Menelusuri pohon AST secara rekursif dan mengembalikan objek nilai runtime.

Penjelasan teknis mendalam mengenai arsitektur interpreter dapat dibaca di **[docs/ARSITEKTUR.md](docs/ARSITEKTUR.md)**.

---

## Indeks Dokumentasi

| Dokumen | Deskripsi |
| :--- | :--- |
| **[docs/SINTAKS.md](docs/SINTAKS.md)** | Spesifikasi sintaksis, kata kunci, tipe data, dan kontrol alur |
| **[docs/KAMUS.md](docs/KAMUS.md)** | Dokumentasi tabel huruf Aksara Lontara Unicode dan pendaftaran kata kunci |
| **[docs/PENGGUNAAN.md](docs/PENGGUNAAN.md)** | Panduan instalasi, konfigurasi `PATH`, dan penggunaan perintah CLI |
| **[docs/ARSITEKTUR.md](docs/ARSITEKTUR.md)** | Detail rancangan mesin interpreter (Lexer, Parser, AST, Evaluator) |
| **[docs/ROADMAP.md](docs/ROADMAP.md)** | Peta jalan fitur & tahapan pengembangan `lontara-lang` |
| **[docs/WASM.md](docs/WASM.md)** | Panduan kompilator WebAssembly & integrasi Web Playground |
| **[docs/KONTRIBUSI.md](docs/KONTRIBUSI.md)** | Panduan kontribusi kode, alur Git branch, dan standar pengujian |

---

## Kontribusi

Kontribusi dari komunitas sangat disukai! Silakan baca panduan lengkap alur kerja kontribusi pada **[docs/KONTRIBUSI.md](docs/KONTRIBUSI.md)** sebelum mengajukan Pull Request atau membuat Issue baru.

---

## Lisensi

Proyek ini didistribusikan di bawah lisensi [MIT License](LICENSE).
