# Automasi GitHub Workflows, Release & Keamanan

Repositori `lontara-lang` telah dilengkapi dengan sistem automasi CI/CD, rilis biner otomatis, pemindaian keamanan CVE, serta pembaruan kontributor otomatis melalui GitHub Actions.

---

## 1. 🚀 Automasi Rilis Biner ([GitHub Releases](https://github.com/dayattt111/lontara-lang/releases))
- **File Workflow**: `.github/workflows/release.yml`
- **Cara Kerja**: Setiap kali Anda membuat tag versi rilis baru (contoh: `v0.2.0` atau `v1.0.0`) dan melakukan push tag ke GitHub:
  ```bash
  git tag v0.2.0
  git push origin v0.2.0
  ```
  GitHub Actions akan secara otomatis mengompilasi biner siap pakai untuk **Linux (amd64 & arm64)**, **macOS (amd64 & Apple Silicon arm64)**, **Windows (.exe)**, serta **WebAssembly (`lontara.wasm`)** dan mempublikasikannya langsung ke menu [Releases](https://github.com/dayattt111/lontara-lang/releases).

---

## 2. 🛡️ Pemindaian Keamanan & CVE Otomatis
- **File Workflow**: `.github/workflows/security.yml` & `.github/dependabot.yml`
- **Cara Kerja**:
  - **CodeQL Security Analysis**: Memindai kode dari potensi cela keamanan, *memory leak*, atau *vulnerability* secara otomatis pada setiap PR dan rilis.
  - **Dependabot**: Secara otomatis memantau kerentanan CVE dan membuka *Pull Request* atau *Security Advisory* jika ditemukan dependensi yang rentan.

---

## 3. 👥 Daftar Kontributor Otomatis di README
- **Widget**: `<img src="https://contrib.rocks/image?repo=dayattt111/lontara-lang" />`
- **Cara Kerja**: Foto profil GitHub Anda (dan setiap kontributor yang PR-nya disetujui) akan tampil secara otomatis di halaman utama `README.md` tanpa perlu mengedit file secara manual.
