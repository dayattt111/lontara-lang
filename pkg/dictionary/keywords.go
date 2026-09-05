package dictionary

// Keyword mendefinisikan pemetaan kata kunci dalam bahasa Latin & Aksara Lontara
type Keyword struct {
	TokenType string // Tipe token (misal: "TAROI", "JAMAGAU")
	Latin     string // Istilah Bugis Latin (misal: "taroi")
	Lontara   string // Istilah Aksara Lontara Unicode (misal: "ᨈᨑᨚᨕᨗ")
	Meaning   string // Padanan umum (misal: "var / let", "func")
}

// RegisteredKeywords memuat seluruh daftar kata kunci resmi lontara-lang.
// Untuk menambah kata kunci baru, cukup tambahkan entri ke dalam slice ini.
var RegisteredKeywords = []Keyword{
	{
		TokenType: "JAMAGAU",
		Latin:     "jamagau",
		Lontara:   "ᨍᨆᨁᨕᨘ",
		Meaning:   "Fungsi (func / function)",
	},
	{
		TokenType: "TAROI",
		Latin:     "taroi",
		Lontara:   "ᨈᨑᨚᨕᨗ",
		Meaning:   "Deklarasi variabel (let / var)",
	},
	{
		TokenType: "REKKO",
		Latin:     "rekko",
		Lontara:   "ᨑᨙᨀᨚ",
		Meaning:   "Percabangan kondisi (if)",
	},
	{
		TokenType: "SANGADINNA",
		Latin:     "sangadinna",
		Lontara:   "ᨔᨂᨉᨗᨊ",
		Meaning:   "Percabangan alternatif (else)",
	},
	{
		TokenType: "TONGENG",
		Latin:     "tongeng",
		Lontara:   "ᨈᨚᨂᨙ",
		Meaning:   "Nilai kebenaran true (true)",
	},
	{
		TokenType: "BANNA",
		Latin:     "banna",
		Lontara:   "ᨅᨊ",
		Meaning:   "Nilai kebenaran false (false)",
	},
	{
		TokenType: "PAUI",
		Latin:     "paui",
		Lontara:   "ᨄᨕᨘᨕᨗ",
		Meaning:   "Cetak keluaran konsol (print / println)",
	},
	{
		TokenType: "LISU",
		Latin:     "lisu",
		Lontara:   "ᨒᨗᨔᨘ",
		Meaning:   "Pengembalian nilai dari fungsi (return)",
	},
	{
		TokenType: "SIKI",
		Latin:     "siki",
		Lontara:   "ᨔᨗᨀᨗ",
		Meaning:   "Perulangan berbasis kondisi (while / loop)",
	},
}

var keywordsMap map[string]string

func init() {
	keywordsMap = make(map[string]string)
	for _, kw := range RegisteredKeywords {
		if kw.Latin != "" {
			keywordsMap[kw.Latin] = kw.TokenType
		}
		if kw.Lontara != "" {
			keywordsMap[kw.Lontara] = kw.TokenType
		}
	}
}

// LookupKeyword mencari tipe token berdasarkan kata (Latin maupun Aksara Lontara)
func LookupKeyword(word string) (string, bool) {
	tokType, ok := keywordsMap[word]
	return tokType, ok
}

// GetKeywordsMap mengembalikan salinan map kata kunci
func GetKeywordsMap() map[string]string {
	copyMap := make(map[string]string, len(keywordsMap))
	for k, v := range keywordsMap {
		copyMap[k] = v
	}
	return copyMap
}
