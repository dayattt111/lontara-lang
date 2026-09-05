# Belajar Lontara-Lang: 3. Percabangan Kondisi (`rekko` / `sangadinna`)

Percabangan logika menggunakan kata kunci `rekko` (if) dan `sangadinna` (else).

---

## Sintaksis

### Versi Bugis Latin
```lontara
taroi skor = 85;

rekko (skor >= 70) {
    paui("Lulus sitinaja");
} sangadinna {
    paui("Kurang");
}
```

### Versi Aksara Lontara
```lontara
ᨈᨑᨚᨕᨗ angka = 10;

ᨑᨙᨀᨚ (angka > 5) {
    ᨄᨕᨘᨕᨗ("salama");
} ᨔᨂᨉᨗᨊ {
    ᨄᨕᨘᨕᨗ("kurang");
}
```

---

## Operator Pembanding

- `==` (Sama dengan)
- `!=` (Tidak sama dengan)
- `<` (Lebih kecil)
- `>` (Lebih besar)
- `<=` (Lebih kecil atau sama dengan)
- `>=` (Lebih besar atau sama dengan)
