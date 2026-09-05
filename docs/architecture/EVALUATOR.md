# Evaluator & Runtime System

Modul `pkg/evaluator` menelusuri simpul-simpul AST dan mengevaluasinya menjadi objek runtime yang mengimplementasikan antarmuka `object.Object` pada `pkg/object`.

---

## Tipe Objek Runtime (`pkg/object`)

- `Integer`: Membungkus nilai `int64`.
- `String`: Membungkus nilai `string`.
- `Boolean`: Membungkus nilai singleton `TRUE` dan `FALSE`.
- `Null`: Membungkus nilai singleton `NULL`.
- `Function`: Menyimpan parameter, badan AST, dan skop leksikal (*Environment*).
- `Error`: Membungkus pesan kesalahan runtime.

---

## Error Reporting Presisi (Fase 2)

Evaluator mengekstrak posisi `Line` dan `Column` dari token AST node untuk setiap kesalahan runtime yang terjadi, sehingga pesan error ditampilkan dengan format presisi:

```text
Kasalahan: [Baris 5, Kolom 17] pambagean nol: tekkulle nibage nol
```
