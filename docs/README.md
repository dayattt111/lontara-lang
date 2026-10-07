<p align="center">
  <img src="assets/lontara.svg" onerror="this.src='../assets/lontara.svg'" alt="Logo Lontara-Lang" width="220">
</p>

# Dokumentasi Resmi Lontara-Lang

Selamat datang di portal dokumentasi resmi **`lontara-lang`**, bahasa pemrograman edukatif berbasis logika bahasa Bugis dan Aksara Lontara Unicode (`U+1A00`–`U+1A1F`).

<p align="center" style="margin: 20px 0;">
  <a href="../" style="display:inline-block; padding: 10px 20px; background-color: #d97706; color: #000; font-weight: bold; text-decoration: none; border: 1px solid #f59e0b; margin-right: 10px;">
    Buka Web Playground Online ↗
  </a>
  <a href="https://github.com/dayattt111/lontara-lang" target="_blank" style="display:inline-block; padding: 10px 20px; background-color: #141417; color: #fbbf24; font-weight: bold; text-decoration: none; border: 1px solid #27272a;">
    GitHub Repository
  </a>
</p>

---

## Panduan Cepat Belajar

Pilih topik pembelajaran di bawah ini atau gunakan navigasi pencarian di sidebar sebelah kiri:

### 1. Modul Belajar Algoritma
Panduan bertahap memahami logika pemrograman dasar menggunakan istilah bahasa Bugis:
* [1. Cetak Konsol (`paui` / `ᨄᨕᨘᨕᨗ`)](algorithm/01_PRINT.md) — Mengeluarkan teks atau nilai ke layar.
* [2. Deklarasi Variabel (`taroi` / `ᨈᨑᨚᨕᨗ`)](algorithm/02_VARIABEL.md) — Menyimpan nilai dan variabel.
* [3. Percabangan Kondisi (`rekko` & `sangadinna`)](algorithm/03_PERCABANGAN.md) — Logika `if-else`.
* [4. Perulangan (`siki` / `ᨔᨗᨀᨗ`)](algorithm/04_PERULANGAN.md) — Loop berbasis kondisi.
* [5. Fungsi & Prosedur (`jamagau` & `lisu`)](algorithm/05_FUNGSI.md) — Deklarasi fungsi dan return value.
* [6. Fungsi Bawaan (`panjang`, `tipe`)](algorithm/06_FUNGSI_BAWAAN.md) — Built-in functions bawaan runtime.

---

### 2. Bahasa & Aksara Lontara
* [Spesifikasi Bahasa](language/SPESIFIKASI.md) — Ringkasan fitur mesin, pengetikan dinamis, dan closures.
* [Panduan CLI & REPL](language/CLI.md) — Cara menjalankan biner `lontara` di terminal lokal.
* [Keyboard & Cara Mengetik Aksara](language/KEYBOARD.md) — Konfigurasi layout keyboard Lontara di Linux, Windows, & Mac.
* [Kamus & Registri Kata Kunci](dictionary/KATA_KUNCI.md) — Pemetaan lengkap istilah Bugis, Lontara, dan maknanya.
* [Tabel Karakter Aksara Lontara](dictionary/TABEL_LONTARA.md) — Inang Sure', Ana' Sure', dan Tanda Baca.

---

### 3. Arsitektur & Performa Mesin
* [Arsitektur Interpreter](architecture/README.md) — Alur Tree-Walking Interpreter (Lexer $\to$ Parser $\to$ AST $\to$ Evaluator).
* [Pratt Parser & Pohon AST](architecture/PARSER.md) — Cara parsing ekspresi dengan operator presedens.
* [Evaluator & Runtime Environment](architecture/EVALUATOR.md) — Mesin evaluasi dan lexical scoping.
* [Analisis Performa (Benchmark)](architecture/BENCHMARK.md) — Hasil profiling CPU/Memori dengan Flame Graph & Call Graph.
* [WebAssembly Engine (WASM)](website/WASM.md) — Kompilasi Go ke WASM untuk browser.

---

### 4. Kontribusi & Roadmap
* [Peta Jalan (Roadmap)](roadmap/README.md) — Rencana pengembangan fitur berikutnya.
* [Alur Kerja Pipeline](roadmap/WORKFLOW.md) — Standarisasi workflow pengembangan.
* [Panduan Berkontribusi](kontribusi/PANDUAN.md) — Konvensi Git, commit, dan Pull Request.
* [Tutorial Menambah Kata Kunci](kontribusi/KAMUS_TUTORIAL.md) — Langkah-langkah memperkaya kosakata Bugis/Lontara.
* [Automasi CI/CD](kontribusi/AUTOMATION.md) — Rilis biner otomatis & pemindaian keamanan.
