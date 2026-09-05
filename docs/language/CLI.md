# Panduan Instalasi & Penggunaan CLI (`lontara`)

---

## 1. Cara Instalasi Biner CLI

### Via Go Remote Install
```bash
go install github.com/dayattt111/lontara-lang/cmd/lontara@latest
```

### Dari Repositori Lokal
```bash
go install ./cmd/lontara
```

---

## 2. Pengaturan Variabel PATH

Jika muncul `command not found`, tambahkan `$GOPATH/bin` ke `PATH` shell Anda (`~/.bashrc` atau `~/.zshrc`):

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

---

## 3. Eksekusi Berkas & REPL

### Menjalankan Berkas Kode (.bugis / .lontara)
```bash
lontara examples/halo_dunia.bugis
lontara run examples/halo_dunia.lontara
```

### Mode Interaktif REPL
Jalankan `lontara` tanpa argumen:
```bash
$ lontara
Lontara Programming Language (v0.1.0)
lontara> taroi a = 10;
lontara> paui(a * 2);
20
```
