package keyloc

import (
	"testing"
)

func TestGetLanguages(t *testing.T) {
	langs, err := getLanguages()
	if err != nil {
		t.Fatalf("getLanguages() returned an error: %v", err)
	}

	expectedLangs := map[string]bool{
		"en": true,
		"ko": true,
	}

	foundLangs := make(map[string]bool)
	for _, lang := range langs {
		normalized := normalizeLangCode(lang)
		if expectedLangs[normalized] {
			foundLangs[normalized] = true
		}
	}

	if len(foundLangs) != len(expectedLangs) {
		t.Errorf("Did not find all expected languages. Got: %v, Want: %v", langs, expectedLangs)
	}
}

func TestCheckLanguage(t *testing.T) {
	tests := []struct {
		lang     string
		expected bool
		name     string
	}{
		{"en", true, "Lowercase"},
		{"ko", true, "Lowercase Korean"},
		{"ru", false, "Russian not present"},
		{"EN", true, "Uppercase"},
		{"en-US", true, "Locale format en-US"},
		{"en_GB", true, "Locale format en_GB"},
		{"ko-KR", true, "Locale format ko-KR"},
		{"ja", false, "Japanese not present"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := CheckLanguage(test.lang)
			if err != nil {
				t.Errorf("CheckLanguage(%q) returned an error: %v", test.lang, err)
			}
			if got != test.expected {
				t.Errorf("CheckLanguage(%q) = %v, want %v", test.lang, got, test.expected)
			}
		})
	}
}

func TestNormalizeLangCode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"en-US", "en"},
		{"fr_CA", "fr"},
		{"ZH-Hant", "zh"},
		{"es", "es"},
		{"DE", "de"},
	}

	for _, test := range tests {
		if normalizeLangCode(test.input) != test.expected {
			t.Errorf("normalizeLangCode(%q) = %q, want %q", test.input, normalizeLangCode(test.input), test.expected)
		}
	}
}
