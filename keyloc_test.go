package keyloc

import (
	"testing"
	"runtime"
)

func TestGetLanguages(t *testing.T) {
	langs, err := getLanguages()
	if err != nil {
		t.Fatalf("getLanguages() returned an error: %v", err)
	}

	expectedLangs := map[string]bool{
		"en": true,
	}

	// Only expect 'ko' on macOS and Windows, not necessarily on Linux CI
	if runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
		expectedLangs["ko"] = true
	}

	foundLangs := make(map[string]bool)
	for _, lang := range langs {
		normalized := normalizeLangCode(lang)
		if expectedLangs[normalized] {
			foundLangs[normalized] = true
		}
	}

	for expectedLang := range expectedLangs {
		if !foundLangs[expectedLang] {
			t.Errorf("Did not find expected language: %q. Found: %v", expectedLang, langs)
		}
	}
}

func TestCheckLanguage(t *testing.T) {
	tests := []struct {
		lang     string
		expected bool
		name     string
		skipOnOS string // "linux", "darwin", "windows" or empty
	}{
		{"en", true, "Lowercase", ""},
		{"ko", true, "Lowercase Korean", "linux"},
		{"ru", false, "Russian not present", ""},
		{"EN", true, "Uppercase", ""},
		{"en-US", true, "Locale format en-US", ""},
		{"en_GB", true, "Locale format en_GB", ""},
		{"ko-KR", true, "Locale format ko-KR", "linux"},
		{"ja", false, "Japanese not present", ""},
	}

	for _, test := range tests {
		if test.skipOnOS == runtime.GOOS {
			t.Logf("Skipping test %q on %s", test.name, runtime.GOOS)
			continue
		}
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