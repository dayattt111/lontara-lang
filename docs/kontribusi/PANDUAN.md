# Panduan Kontribusi Lontara-Lang

Terima kasih telah berpartisipasi dalam pengembangan `lontara-lang`!

---

## Alur Kerja Git & Pull Request

1. **Fork Repositori**: Buat fork ke akun GitHub Anda.
2. **Clone Lokal**:
   ```bash
   git clone https://github.com/<username>/lontara-lang.git
   cd lontara-lang
   ```
3. **Buat Branch Fitur**:
   ```bash
   git checkout -b fitur/nama-fitur
   ```
4. **Jalankan Pengujian**:
   ```bash
   go test -v ./pkg/...
   ```
5. **Commit & Push**:
   ```bash
   git commit -m "feat: deskripsi perubahan"
   git push origin fitur/nama-fitur
   ```
6. **Kirim Pull Request (PR)** ke branch `main`.
