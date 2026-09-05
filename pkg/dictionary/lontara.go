package dictionary

// Character mendefinisikan informasi satu huruf/karakter Aksara Lontara
type Character struct {
	Rune        rune   // Karakter Unicode rune (misal: 'ᨀ')
	Hex         string // Kode Heksadesimal Unicode (misal: "U+1A00")
	Name        string // Nama karakter Bugis/Lontara (misal: "Ka")
	Latin       string // Transliterasi Latin (misal: "ka")
	Category    string // Kategori: "Inang Sure'" (Huruf Utama), "Ana' Sure'" (Vokalisasi), "Tanda Baca"
	Description string // Keterangan tambahan
}

// LontaraAlphabet menyimpan tabel lengkap karakter Aksara Lontara (U+1A00 - U+1A1F)
var LontaraAlphabet = []Character{
	// --- Inang Sure' (Huruf Utama / Konsonan Suku Kata) ---
	{Rune: 'ᨀ', Hex: "U+1A00", Name: "Ka", Latin: "ka", Category: "Inang Sure'", Description: "Konsonan Ka"},
	{Rune: 'ᨁ', Hex: "U+1A01", Name: "Ga", Latin: "ga", Category: "Inang Sure'", Description: "Konsonan Ga"},
	{Rune: 'ᨂ', Hex: "U+1A02", Name: "Nga", Latin: "nga", Category: "Inang Sure'", Description: "Konsonan Nga"},
	{Rune: 'ᨃ', Hex: "U+1A03", Name: "Ngka", Latin: "ngka", Category: "Inang Sure'", Description: "Konsonan Ngka"},
	{Rune: 'ᨄ', Hex: "U+1A04", Name: "Pa", Latin: "pa", Category: "Inang Sure'", Description: "Konsonan Pa"},
	{Rune: 'ᨅ', Hex: "U+1A05", Name: "Ba", Latin: "ba", Category: "Inang Sure'", Description: "Konsonan Ba"},
	{Rune: 'ᨆ', Hex: "U+1A06", Name: "Ma", Latin: "ma", Category: "Inang Sure'", Description: "Konsonan Ma"},
	{Rune: 'ᨇ', Hex: "U+1A07", Name: "Mpa", Latin: "mpa", Category: "Inang Sure'", Description: "Konsonan Mpa"},
	{Rune: 'ᨈ', Hex: "U+1A08", Name: "Ta", Latin: "ta", Category: "Inang Sure'", Description: "Konsonan Ta"},
	{Rune: 'ᨉ', Hex: "U+1A09", Name: "Da", Latin: "da", Category: "Inang Sure'", Description: "Konsonan Da"},
	{Rune: 'ᨊ', Hex: "U+1A0A", Name: "Na", Latin: "na", Category: "Inang Sure'", Description: "Konsonan Na"},
	{Rune: 'ᨋ', Hex: "U+1A0B", Name: "Nra", Latin: "nra", Category: "Inang Sure'", Description: "Konsonan Nra"},
	{Rune: 'ᨌ', Hex: "U+1A0C", Name: "Ca", Latin: "ca", Category: "Inang Sure'", Description: "Konsonan Ca"},
	{Rune: 'ᨍ', Hex: "U+1A0D", Name: "Ja", Latin: "ja", Category: "Inang Sure'", Description: "Konsonan Ja"},
	{Rune: 'ᨎ', Hex: "U+1A0E", Name: "Nya", Latin: "nya", Category: "Inang Sure'", Description: "Konsonan Nya"},
	{Rune: 'ᨏ', Hex: "U+1A0F", Name: "Nca", Latin: "nca", Category: "Inang Sure'", Description: "Konsonan Nca"},
	{Rune: 'ᨐ', Hex: "U+1A10", Name: "Ya", Latin: "ya", Category: "Inang Sure'", Description: "Konsonan Ya"},
	{Rune: 'ᨑ', Hex: "U+1A11", Name: "Ra", Latin: "ra", Category: "Inang Sure'", Description: "Konsonan Ra"},
	{Rune: 'ᨒ', Hex: "U+1A12", Name: "La", Latin: "la", Category: "Inang Sure'", Description: "Konsonan La"},
	{Rune: 'ᨓ', Hex: "U+1A13", Name: "Wa", Latin: "wa", Category: "Inang Sure'", Description: "Konsonan Wa"},
	{Rune: 'ᨔ', Hex: "U+1A14", Name: "Sa", Latin: "sa", Category: "Inang Sure'", Description: "Konsonan Sa"},
	{Rune: 'ᨕ', Hex: "U+1A15", Name: "A", Latin: "a", Category: "Inang Sure'", Description: "Vokal Dasar A"},
	{Rune: 'ᨖ', Hex: "U+1A16", Name: "Ha", Latin: "ha", Category: "Inang Sure'", Description: "Konsonan Ha"},

	// --- Ana' Sure' (Diakritik Vokal / Vokal Tambahan) ---
	{Rune: 'ᨗ', Hex: "U+1A17", Name: "Tetteng (i)", Latin: "i", Category: "Ana' Sure'", Description: "Diakritik vokal i (atas)"},
	{Rune: 'ᨘ', Hex: "U+1A18", Name: "Pucu' (u)", Latin: "u", Category: "Ana' Sure'", Description: "Diakritik vokal u (bawah)"},
	{Rune: 'ᨙ', Hex: "U+1A19", Name: "Kelling (e)", Latin: "e", Category: "Ana' Sure'", Description: "Diakritik vokal e (kiri)"},
	{Rune: 'ᨚ', Hex: "U+1A1A", Name: "Doping (o)", Latin: "o", Category: "Ana' Sure'", Description: "Diakritik vokal o (kanan)"},
	{Rune: 'ᨛ', Hex: "U+1A1B", Name: "Kecce' (ae)", Latin: "ae", Category: "Ana' Sure'", Description: "Diakritik vokal ae / pepet (atas)"},

	// --- Tanda Baca ---
	{Rune: '᨞', Hex: "U+1A1E", Name: "Pallawa", Latin: ",", Category: "Tanda Baca", Description: "Pemisah kalimat / koma Lontara"},
	{Rune: '᨟', Hex: "U+1A1F", Name: "End of Section", Latin: ".", Category: "Tanda Baca", Description: "Penutup pasal / titik Lontara"},
}

var charMap map[rune]Character

func init() {
	charMap = make(map[rune]Character)
	for _, c := range LontaraAlphabet {
		charMap[c.Rune] = c
	}
}

// IsLontaraRune memeriksa apakah karakter termasuk dalam blok Unicode Aksara Lontara (U+1A00 - U+1A1F)
func IsLontaraRune(r rune) bool {
	return r >= 0x1A00 && r <= 0x1A1F
}

// GetCharacterInfo mengembalikan detail informasi karakter Aksara Lontara berdasarkan rune
func GetCharacterInfo(r rune) (Character, bool) {
	c, ok := charMap[r]
	return c, ok
}
