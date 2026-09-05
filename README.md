# lontara-lang

<p align="center">
    <img src="assets/lontara.svg" alt="Lontara-Lang Logo By Muhammad Amin Hidayat(Gen AI)" width="400">
</p>

![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-orange?style=flat-square)

**Bahasa Pemrograman Edukatif Berbasis Logika Bahasa Bugis dan Aksara Lontara.**

`lontara-lang` adalah bahasa pemrograman yang dirancang untuk mengenalkan konsep-konsep dasar ilmu komputer dan logika algoritma melalui kosa kata bahasa Bugis dan dukungan native Aksara Lontara. Bahasa ini diimplementasikan menggunakan Go (Golang) murni tanpa ketergantungan pustaka eksternal (*zero external dependencies*).

---

## 1. Latar Belakang dan Filosofi

Pengembangan `lontara-lang` didorong oleh tiga pilar utama:

1. **Edukasi Algoritma dan Pemrograman:** Memudahkan pengenalan pemrograman bagi pelajar di daerah Sulawesi Selatan dan masyarakat luas dengan menggunakan istilah lokal yang akrab dan mudah dipahami.
2. **Preservasi Aksara dan Budaya Digital:** Mendukung penulisan kode menggunakan karakter Unicode Aksara Lontara (`U+1A00` sampai `U+1A1F`). Langkah ini bertujuan menjadikan teknologi digital sebagai media aktif pelestarian warisan budaya Nusantara.
3. **Portabilitas dan Kemudahan Akses:** Dibangun di atas bahasa Go untuk menghasilkan static binary tunggal yang ringan, cepat, dan dapat dijalankan di berbagai sistem operasi tanpa konfigurasi lingkungan runtime yang rumit.

---

## 2. Inspirasi dan Ucapan Terima Kasih

Proyek `lontara-lang` terinspirasi oleh karya-karya hebat dalam komunitas bahasa pemrograman esoteris (*esolang*) dan edukatif Nusantara pendahulu, seperti:

- **JawaScript (`.jawa`)** – Bahasa pemrograman berbasis sintaks bahasa Jawa.
- **SundaScript / Sundalang** – Bahasa pemrograman berbasis sintaks bahasa Sunda.

Apresiasi tinggi disampaikan kepada para kreator proyek-proyek tersebut yang telah membuka jalan bagi integrasi bahasa daerah ke dalam dunia teknologi perangkat lunak.

---

## 3. Tabel Pemetaan Sintaks

Berikut adalah tabel perbandingan sintaks kata kunci antara bahasa umum (JavaScript/Go/Python), versi Bugis Latin, dan versi Aksara Lontara:

| Konsep Umum | Bugis Latin | Aksara Lontara | Keterangan |
| :--- | :--- | :--- | :--- |
| `let / var` | `taroi` | `ᨈᨑᨚᨕᨗ` | Deklarasi variabel |
| `func` | `jamagau` | `ᨍᨆᨁᨕᨘ` | Deklarasi fungsi |
| `if` | `rekko` | `ᨑᨙᨀᨚ` | Percabangan kondisi |
| `else` | `sangadinna` | `ᨔᨂᨉᨗᨊ` | Kondisi alternatif |
| `true` | `tongeng` | `ᨈᨚᨂᨙ` | Nilai kebenaran true |
| `false` | `banna` | `ᨅᨊ` | Nilai kebenaran false |
| `return` | `lisu` | `ᨒᨗᨔᨘ` | Pengembalian nilai dari fungsi |
| `print` | `paui` | `ᨄᨕᨘᨕᨗ` | Mencetak ke konsol/stdout |

---

## 4. Contoh Kode

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

## 5. Panduan Instalasi dan Penggunaan

### Prasyarat
- Go versi 1.21 atau yang lebih baru.

### Instalasi CLI

Untuk pengguna lain yang ingin mencoba `lontara` dari repositori GitHub:

```bash
go install github.com/dayattt111/lontara-lang/cmd/lontara@latest
```

Untuk pengembanan lokal dalam repositori ini:

```bash
go install ./cmd/lontara
```

*Catatan: Jika muncul error `lontara: command not found`, pastikan `$GOPATH/bin` atau `~/go/bin` terdaftar dalam variabel `PATH` shell Anda:*
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Menjalankan Berkas Kode (`.bugis` / `.lontara`)

Anda dapat mengeksekusi berkas sumber secara langsung atau melalui sub-perintah `run`:

```bash
# Eksekusi langsung
lontara test.bugis
lontara test.lontara

# Eksekusi dengan sub-perintah run
lontara run test.bugis
lontara run test.lontara
```

Jika ingin menjalankan tanpa melakukan `go install` (saat pengembangan repositori):

```bash
go run ./cmd/lontara examples/halo_dunia.bugis
```

### Mode Interaktif (REPL)

Jalankan perintah `lontara` tanpa argumen untuk masuk ke sesi REPL (*Read-Eval-Print Loop*):

```bash
$ lontara
Lontara Programming Language (v0.1.0)
Ketik perintah atau ekspresi untuk mengevaluasi.
lontara> taroi a = 10;
lontara> taroi b = 20;
lontara> a + b;
30
lontara> paui("Salam!");
Salam!
```

Dokumentasi lengkap penanganan masalah dan instalasi dapat dibaca di **[docs/PENGGUNAAN.md](docs/PENGGUNAAN.md)**.

---

## 6. Struktur Proyek

Struktur repositori `lontara-lang` disusun modular mengikuti konvensi proyek Go:

```text
.
├── LICENSE                  # Lisensi proyek (MIT License)
├── README.md                # Dokumentasi utama repositori
├── cmd/
│   ├── lontara/             # Entry point untuk CLI runner dan REPL (main.go)
│   └── wasm/                # Entry point untuk kompilasi WebAssembly (main.go)
├── docs/
│   ├── ARSITEKTUR.md        # Dokumentasi arsitektur internal interpreter
│   ├── PENGGUNAAN.md        # Panduan detail instalasi, PATH, dan CLI runner
│   └── SINTAKS.md           # Spesifikasi sintaksis dan manual tata bahasa
├── examples/
│   ├── halo_dunia.bugis     # Contoh program dalam versi Bugis Latin
│   └── halo_dunia.lontara   # Contoh program dalam versi Aksara Lontara Unicode
├── go.mod                   # Berkas modul Go
└── pkg/
    ├── ast/                 # Node definisi Abstract Syntax Tree (AST)
    ├── evaluator/           # Mesin penelusur AST dan evaluasi nilai runtime
    ├── lexer/               # Analisis leksikal dan pembacaan token UTF-8
    ├── object/              # Sistem tipe data objek runtime dan lingkungan variabel
    ├── parser/              # Parser sintaksis mengimplementasikan Pratt Parser
    ├── repl/                # Antarmuka REPL terminal interaktif
    └── token/               # Definisi tipe token dan pemetaan kata kunci
```

---

## 7. Roadmap Milestone

- **Fase 1 (Tahap Saat Ini): CLI Core Engine**
  - Implementasi interpreter murni (*tree-walking interpreter*) di terminal.
  - Sesi REPL interaktif dan pemrosesan berkas (`.lontara` / `.bugis`).
  - Dukungan tipe data dasar: Integer, String, Boolean, Null.
  - Operator aritmetika dan pembanding (`+`, `-`, `*`, `/`, `<`, `>`, `<=`, `>=`, `==`, `!=`).
  - Deklarasi variabel (`taroi`), fungsi anonim & rekursif (`jamagau`), pengembalian nilai (`lisu`).
  - Percabangan kondisi (`rekko` ... `sangadinna`).
  - Cetak keluaran bawaan (`paui` / `ᨄᨕᨘᨕᨗ`).

- **Fase 2: WebAssembly Playground (Wasm Engine)**
  - Kompilasi modul Wasm untuk mendukung *playground* interaktif berbasis web tanpa ketergantungan pada backend server.

- **Fase 3: Ekosistem & Standard Library Lontara**
  - Integrasi penuh kamus kata kunci berbasis Unicode Aksara Lontara.
  - Pengembangan pustaka standar untuk manipulasi string, array, dan fungsi I/O tambahan.

---

## 8. Kontribusi dan Lisensi

### Kontribusi
Kontribusi dari komunitas sangat terbuka! Anda dapat berkontribusi melalui:
- Pelaporan *bug* atau usulan fitur via **GitHub Issues**.
- Pengiriman perbaikan atau fitur baru via **Pull Request**.
- Penambahan contoh program atau penyempurnaan dokumentasi.

### Lisensi
Proyek ini didistribusikan di bawah lisensi [MIT License](LICENSE).
