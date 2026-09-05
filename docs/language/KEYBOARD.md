# Panduan Keyboard & Cara Mengetik Aksara Lontara

`lontara-lang` dirancang agar mudah digunakan oleh siapa saja. Dokumentasi ini menjelaskan opsi penulisan dan cara mengetik Aksara Lontara di berbagai sistem operasi.

---

## 💡 2 Opsi Penulisan Kode

### 1. Bugis Latin (Sangat Direkomendasikan untuk Pemula)
Anda **tidak memerlukan keyboard khusus**. Cukup gunakan keyboard QWERTY standar dengan kata kunci istilah Bugis Latin:

```lontara
taroi i = 1;
siki (i <= 5) {
    paui("Angka:", i);
    taroi i = i + 1;
}
```

---

### 2. Aksara Lontara Unicode Native (`U+1A00`–`U+1A1F`)
Jika Anda ingin menulis langsung menggunakan karakter asli Aksara Lontara:

```lontara
ᨈᨑᨚᨕᨗ i = 1;
ᨔᨗᨀᨗ (i <= 5) {
    ᨄᨕᨘᨕᨗ("Angka:", i);
    ᨈᨑᨚᨕᨗ i = i + 1;
}
```

---

## ⌨️ Cara Mengetik Aksara Lontara di Berbagai Platform

### A. Di Web Browser (Web Playground)
Cara paling mudah tanpa install apapun:
Buka **[examples/playground.html](../../examples/playground.html)** pada browser Anda. Terdapat menu dropdown sampel otomatis yang dapat langsung diklik dan dijalankan.

---

### B. Linux (Ubuntu / Debian / Fedora / Arch)
1. **Menggunakan Unicode Hex Input**:
   - Tekan `Ctrl + Shift + U` bersamaan.
   - Ketik kode heksadesimal Unicode Lontara (contoh: `1a08` untuk `ᨈ`).
   - Tekan `Enter` atau `Spasi`.
2. **Menggunakan Keyman / IBus**:
   - Pasang IBus / Keyman: `sudo apt install ibus-table` atau `sudo apt install keyman`.
   - Tambahkan tata letak (*layout*) **Bugis (Lontara)** pada pengaturan Input Source OS Anda.

---

### C. Windows (10 / 11)
1. **Aplikasi Keyman Desktop**:
   - Unduh gratis **Keyman Desktop for Windows** dari `keyman.com`.
   - Pasang paket keyboard **Bugis / Lontara Keyboard**.
   - Aktifkan keyboard melalui taskbar bahasa (ALT + SHIFT).
2. **Font Support**: Windows 10 & 11 sudah menyertakan font bawaan *Leelawadee UI* yang mendukung rendering Aksara Lontara Unicode secara rapi.

---

### D. macOS
1. Unduh aplikasi **Keyman for macOS** dari `keyman.com`.
2. Pasang tata letak keyboard Bugis Lontara.
3. Atau gunakan metode Unicode Hex Input macOS via *System Preferences -> Keyboard -> Input Sources*.

---

## 📋 Tabel Pemetaan Ringkas Karakter & Kode Unicode

| Tipe Karakter | Contoh Huruf | Kode Hex | Cara Baca |
| :--- | :---: | :---: | :--- |
| **Inang Sure' (Konsonan)** | `ᨀ` | `U+1A00` | Ka |
| | `ᨁ` | `U+1A01` | Ga |
| | `ᨈ` | `U+1A08` | Ta |
| | `ᨑ` | `U+1A11` | Ra |
| | `ᨔ` | `U+1A14` | Sa |
| **Ana' Sure' (Vokal)** | `ᨗ` | `U+1A17` | Vokal i (Tetteng) |
| | `ᨘ` | `U+1A18` | Vokal u (Pucu') |
| | `ᨙ` | `U+1A19` | Vokal e (Kelling) |
| | `ᨚ` | `U+1A1A` | Vokal o (Doping) |
| **Tanda Baca** | `᨞` | `U+1A1E` | Pemisah Kalimat (Koma / Pallawa) |
| | `᨟` | `U+1A1F` | Penutup Paragraf (Titik) |

Detail tabel selengkapnya dapat dibaca pada **[docs/dictionary/TABEL_LONTARA.md](../dictionary/TABEL_LONTARA.md)**.
