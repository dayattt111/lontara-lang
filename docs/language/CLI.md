# Panduan Instalasi & Penggunaan CLI (`lontara`)

Dokumen ini menjelaskan petunjuk lengkap bagi pengguna umum (luar) maupun pengembang lokal yang ingin menguji dan mengeksekusi program `lontara-lang`.

---

## 🚀 1. Cara Mencoba Tanpa Clone Repositori

### Opsi A: Tanpa Instalasi (Web Playground di Browser)
Jika Anda hanya ingin mencoba menulis dan mengeksekusi kode secara langsung tanpa memasang apapun:
- Buka file **`examples/playground.html`** di browser Anda.
- Anda dapat langsung mengetik kode Bugis Latin atau Aksara Lontara dan mengeklik tombol **Jalankan Kode ▶**.

### Opsi B: Pengguna Umum via `go install` (Global CLI)
Jika di komputer Anda sudah terpasang Go (Golang 1.21+), Anda dapat memasang biner CLI `lontara` secara global tanpa perlu melakukan git clone:

```bash
go install github.com/dayattt111/lontara-lang/cmd/lontara@latest
```

---

## ⚙️ 2. Pengaturan Variabel PATH

Jika setelah menjalankan `go install` perintah `lontara` menghasilkan error `command not found`, tambahkan direktori biner Go ke variabel `PATH` shell Anda (`~/.bashrc`, `~/.zshrc`, atau `~/.profile`):

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Muat ulang konfigurasi shell:
```bash
source ~/.bashrc   # atau source ~/.zshrc
```

---

## 📝 3. Langkah Menulis & Eksekusi Kode Pertama

Setelah `lontara` terpasang, Anda dapat membuat berkas kode pertama dengan editor teks apapun (VS Code, Notepad, Vim, dll.):

1. Buat berkas baru bernama `halo.bugis`:
   ```lontara
   taroi pesang = "Salama dari Lontara-Lang!";
   paui(pesang);
   ```

2. Jalankan berkas melalui terminal:
   ```bash
   lontara halo.bugis
   # atau
   lontara run halo.bugis
   ```

---

## 💻 4. Mode Interaktif (REPL)

Jalankan perintah `lontara` tanpa argumen untuk masuk ke mode REPL terminal interaktif:

```bash
$ lontara
Lontara Programming Language (v0.1.0)
Ketik perintah atau ekspresi untuk mengevaluasi.
lontara> taroi a = 10;
lontara> paui(a * 2);
20
```
