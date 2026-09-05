package dictionary

import (
	"testing"
)

func TestLookupKeyword(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		found    bool
	}{
		{"taroi", "TAROI", true},
		{"ᨈᨑᨚᨕᨗ", "TAROI", true},
		{"jamagau", "JAMAGAU", true},
		{"ᨍᨆᨁᨕᨘ", "JAMAGAU", true},
		{"rekko", "REKKO", true},
		{"sangadinna", "SANGADINNA", true},
		{"tongeng", "TONGENG", true},
		{"banna", "BANNA", true},
		{"paui", "PAUI", true},
		{"lisu", "LISU", true},
		{"sembarang", "", false},
	}

	for _, tt := range tests {
		tokType, found := LookupKeyword(tt.input)
		if found != tt.found {
			t.Errorf("input %q: ditemukan=%v, diharapkan=%v", tt.input, found, tt.found)
		}
		if tokType != tt.expected {
			t.Errorf("input %q: tipe=%q, diharapkan=%q", tt.input, tokType, tt.expected)
		}
	}
}

func TestLontaraAlphabet(t *testing.T) {
	if len(LontaraAlphabet) == 0 {
		t.Fatal("LontaraAlphabet tidak boleh kosong")
	}

	for _, char := range LontaraAlphabet {
		if !IsLontaraRune(char.Rune) {
			t.Errorf("rune %q (%U) harus berada di rentang Lontara U+1A00..U+1A1F", char.Rune, char.Rune)
		}

		info, ok := GetCharacterInfo(char.Rune)
		if !ok {
			t.Errorf("GetCharacterInfo gagal untuk rune %q", char.Rune)
		}
		if info.Name != char.Name {
			t.Errorf("Nama tidak cocok: %q != %q", info.Name, char.Name)
		}
	}
}
