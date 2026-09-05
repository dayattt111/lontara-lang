# Panduan Penggunaan & Arsitektur WebAssembly (WASM)

Dokumen ini menjelaskan rancangan, kompilasi, dan integrasi modul WebAssembly (`cmd/wasm`) pada `lontara-lang`.

---

## 1. Latar Belakang

Modul WebAssembly memungkinkan interpreter `lontara-lang` dijalankan secara langsung di dalam peramban web (browser) pengguna tanpa memerlukan server backend untuk mengeksekusi kode. Hal ini berguna untuk membuat **Web Playground** interaktif di mana pengguna dapat mencoba menulis dan menjalankan kode sintaks Bugis/Lontara langsung di web.

---

## 2. Struktur Direktori WASM

```text
cmd/wasm/
└── main.go       # Titik masuk utama eksekutor WebAssembly
```

---

## 3. Kompilasi ke Biner WebAssembly

Untuk mengompilasi biner Go menjadi file `.wasm`, gunakan variabel lingkungan `GOOS=js` dan `GOARCH=wasm`:

```bash
GOOS=js GOARCH=wasm go build -o lontara.wasm ./cmd/wasm
```

---

## 4. Integrasi dengan Halaman Web (HTML & JavaScript)

### A. Menyalin Jembatan JavaScript Go (`wasm_exec.js`)

Go menyediakan file jembatan runtime `wasm_exec.js` yang harus diikutsertakan di dalam peramban web. Anda dapat menyalin berkas tersebut dari instalasi Go lokal Anda:

```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

### B. Contoh Implementasi HTML/JS

```html
<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <title>Lontara-Lang Web Playground</title>
    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("lontara.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        });
    </script>
</head>
<body>
    <h1>Lontara-Lang Web Playground</h1>
    <p>Buka konsol pengembang (F12) untuk melihat luaran dari mesin WASM Lontara.</p>
</body>
</html>
</html>
```

---

## 5. Status Pengembangan WASM

- [x] Pengaturan struktur direktori biner `cmd/wasm`.
- [x] Implementasi ekspor API JavaScript (`js.Global().Set("evaluateLontara", ...)`) untuk menghubungkan masukan teks dari textarea HTML langsung ke `lexer`, `parser`, dan `evaluator`.
- [x] Pengalihan keluaran `paui` (`stdout`) dan laporan kesalahan ke elemen konsol Web Playground ([examples/playground.html](../examples/playground.html)).
