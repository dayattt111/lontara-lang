# Panduan Instalasi dan Penggunaan Lontara-Lang

Dokumen ini menjelaskan cara memasang, mengonfigurasi `PATH`, dan mengoperasikan `lontara-lang` di berbagai lingkungan sistem operasi.

## 1. Cara Instalasi Biner CLI (`lontara`)

### A. Untuk Pengguna Umum (Via Go Remote Install)
Pengguna luar yang ingin mencoba `lontara` dapat langsung memasangnya dari repositori GitHub menggunakan perintah:

```bash
go install github.com/dayattt111/lontara-lang/cmd/lontara@latest
```

### B. Untuk Pengembang / Lokal Repository
Jika Anda berada dalam folder repositori proyek lokal `lontara-lang`:

```bash
go install ./cmd/lontara
```

Atau kompilasi langsung biner lokal:

```bash
go build -o lontara ./cmd/lontara
```

---

## 2. Solusi Error `lontara: command not found`

Jika setelah menjalankan `go install` perintah `lontara` masih menghasilkan error `command not found`, hal ini disebabkan folder biner Go (`$GOPATH/bin` atau `~/go/bin`) belum terdaftar pada variabel lingkungan `PATH` di shell Anda.

### Solusi Sementara (Session Saat Ini):
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Solusi Permanen (Linux / macOS):
Tambahkan baris berikut ke akhir berkas konfigurasi shell Anda (`~/.bashrc`, `~/.zshrc`, atau `~/.profile`):

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Kemudian muat ulang konfigurasi:
```bash
source ~/.bashrc   # atau source ~/.zshrc
```

---

## 3. Cara Penggunaan Perintah CLI

Setelah `lontara` terpasang dan `PATH` terkonfigurasi, Anda dapat mengeksekusi perintah berikut di terminal:

### A. Menjalankan Berkas Sumber (`.bugis` atau `.lontara`)

Anda dapat mengeksekusi berkas secara langsung dengan menyertakan nama berkas:

```bash
lontara test.bugis
lontara test.lontara
```

Atau menggunakan sub-perintah `run`:

```bash
lontara run test.bugis
lontara run test.lontara
```

### B. Mode Interaktif (REPL)

Jalankan perintah `lontara` tanpa argumen apa pun untuk masuk ke mode REPL interaktif:

```bash
$ lontara
Lontara Programming Language (v0.1.0)
Ketik perintah atau ekspresi untuk mengevaluasi.
lontara> taroi x = 5;
lontara> taroi y = 10;
lontara> x + y;
15
```

---

## 4. Menjalankan Tanpa Instalasi (`go run`)

Jika Anda belum ingin memasang biner `lontara` secara global dan hanya ingin mengeksekusi berkas sumber saat pengembangan repositori, gunakan `go run` yang mengarah ke paket `cmd/lontara`:

```bash
go run ./cmd/lontara examples/halo_dunia.bugis
go run ./cmd/lontara examples/halo_dunia.lontara
```

*Catatan: Menjalankan `go run berkas.bugis` secara langsung akan gagal karena `go run` mengekspektasikan berkas sumber bahasa Go (`.go`).*
