# Panduan WebAssembly (WASM) & Web Playground

Dokumen ini menjelaskan rancangan, kompilasi, dan integrasi modul WebAssembly (`cmd/wasm`) pada `lontara-lang`.

---

## 1. Latar Belakang

Modul WebAssembly memungkinkan interpreter `lontara-lang` dijalankan secara langsung di dalam peramban web (browser) pengguna tanpa memerlukan server backend. Hal ini dipakai untuk menyediakan **Web Playground** interaktif.

---

## 2. Kompilasi biner WebAssembly

```bash
GOOS=js GOARCH=wasm go build -o examples/lontara.wasm ./cmd/wasm
```

---

## 3. Integrasi JavaScript (`syscall/js`)

Biner WASM mengekspor fungsi JavaScript global `evaluateLontara(codeString)` yang mengembalikan luaran teks cetakan `paui` atau laporan kesalahan runtime.

---

## 4. Web Playground UI

Halaman interaktif tersedia di [examples/playground.html](../../examples/playground.html).
