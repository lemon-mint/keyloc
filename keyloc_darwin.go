//go:build darwin
package keyloc

import (
	"os/exec"
	"regexp"
	"strings"
)

// mapIdentifierToLangCode maps common macOS input source identifiers to language codes.
func mapIdentifierToLangCode(identifier string) string {
	lowerIdentifier := strings.ToLower(identifier)
	switch {
	case strings.Contains(lowerIdentifier, "korean"), strings.Contains(lowerIdentifier, "hangul"):
		return "ko"
	case strings.Contains(lowerIdentifier, "u.s."), strings.Contains(lowerIdentifier, "abc"), strings.Contains(lowerIdentifier, "english"):
		return "en"
	case strings.Contains(lowerIdentifier, "russian"), strings.Contains(lowerIdentifier, "cyrillic"):
		return "ru"
	case strings.Contains(lowerIdentifier, "japanese"), strings.Contains(lowerIdentifier, "kana"), strings.Contains(lowerIdentifier, "romaji"):
		return "ja"
	case strings.Contains(lowerIdentifier, "french"):
		return "fr"
	case strings.Contains(lowerIdentifier, "german"):
		return "de"
	case strings.Contains(lowerIdentifier, "spanish"):
		return "es"
	case strings.Contains(lowerIdentifier, "chinese"), strings.Contains(lowerIdentifier, "pinyin"), strings.Contains(lowerIdentifier, "zhuyin"), strings.Contains(lowerIdentifier, "cangjie"):
		return "zh"
	case strings.Contains(lowerIdentifier, "italian"):
		return "it"
	case strings.Contains(lowerIdentifier, "portuguese"):
		return "pt"
	case strings.Contains(lowerIdentifier, "dutch"):
		return "nl"
	case strings.Contains(lowerIdentifier, "swedish"):
		return "sv"
	case strings.Contains(lowerIdentifier, "danish"):
		return "da"
	case strings.Contains(lowerIdentifier, "norwegian"):
		return "no"
	case strings.Contains(lowerIdentifier, "finnish"):
		return "fi"
	case strings.Contains(lowerIdentifier, "polish"):
		return "pl"
	case strings.Contains(lowerIdentifier, "turkish"):
		return "tr"
	case strings.Contains(lowerIdentifier, "arabic"):
		return "ar"
	case strings.Contains(lowerIdentifier, "hebrew"):
		return "he"
	case strings.Contains(lowerIdentifier, "greek"):
		return "el"
	case strings.Contains(lowerIdentifier, "thai"):
		return "th"
	case strings.Contains(lowerIdentifier, "vietnamese"):
		return "vi"
	case strings.Contains(lowerIdentifier, "hindi"):
		return "hi"
	case strings.Contains(lowerIdentifier, "bengali"):
		return "bn"
	case strings.Contains(lowerIdentifier, "punjabi"):
		return "pa"
	case strings.Contains(lowerIdentifier, "gujarati"):
		return "gu"
	case strings.Contains(lowerIdentifier, "tamil"):
		return "ta"
	case strings.Contains(lowerIdentifier, "telugu"):
		return "te"
	case strings.Contains(lowerIdentifier, "kannada"):
		return "kn"
	case strings.Contains(lowerIdentifier, "malayalam"):
		return "ml"
	case strings.Contains(lowerIdentifier, "indonesian"):
		return "id"
	case strings.Contains(lowerIdentifier, "malay"):
		return "ms"
	case strings.Contains(lowerIdentifier, "filipino"):
		return "fil"
	case strings.Contains(lowerIdentifier, "ukrainian"):
		return "uk"
	case strings.Contains(lowerIdentifier, "czech"):
		return "cs"
	case strings.Contains(lowerIdentifier, "slovak"):
		return "sk"
	case strings.Contains(lowerIdentifier, "hungarian"):
		return "hu"
	case strings.Contains(lowerIdentifier, "romanian"):
		return "ro"
	case strings.Contains(lowerIdentifier, "bulgarian"):
		return "bg"
	case strings.Contains(lowerIdentifier, "croatian"):
		return "hr"
	case strings.Contains(lowerIdentifier, "serbian"):
		return "sr"
	case strings.Contains(lowerIdentifier, "slovenian"):
		return "sl"
	case strings.Contains(lowerIdentifier, "estonian"):
		return "et"
	case strings.Contains(lowerIdentifier, "latvian"):
		return "lv"
	case strings.Contains(lowerIdentifier, "lithuanian"):
		return "lt"
	// Add more general mappings here
	default:
		return "" // Return empty if no mapping found
	}
}

func getLanguages() ([]string, error) {
	langSet := make(map[string]bool)

	// Try to get languages from AppleEnabledInputSources (keyboard layouts)
	cmd := exec.Command("defaults", "read", "com.apple.HIToolbox", "AppleEnabledInputSources")
	output, err := cmd.Output()
	if err == nil {
		re := regexp.MustCompile(`(KeyboardLayout Name|Bundle ID) = "?([\w\.]+)"?;`)
		matches := re.FindAllStringSubmatch(string(output), -1)

		for _, match := range matches {
			if len(match) > 2 {
				identifier := match[2]
				if lang := mapIdentifierToLangCode(identifier); lang != "" {
					langSet[lang] = true
				}
			}
		}
	}

	// Try to get languages from AppleLanguages (system preferred languages)
	fallbackLangs, err := getAppleLanguagesFallback()
	if err == nil {
		for _, lang := range fallbackLangs {
			langSet[lang] = true
		}
	}

	// Try to get languages from Voice Services (installed voices)
	voiceLangs, err := getVoiceServicesLanguages()
	if err == nil {
		for _, lang := range voiceLangs {
			langSet[lang] = true
		}
	}

	langs := make([]string, 0, len(langSet))
	for lang := range langSet {
		langs = append(langs, lang)
	}

	return langs, nil
}

func getAppleLanguagesFallback() ([]string, error) {
	cmd := exec.Command("defaults", "read", "-g", "AppleLanguages")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	langSet := make(map[string]bool)
	// Regex to find all language codes like "en-US", "ko"
	re := regexp.MustCompile(`"([a-zA-Z\-]+)"`)
	matches := re.FindAllStringSubmatch(string(output), -1)

	for _, match := range matches {
		if len(match) > 1 {
			// Normalize the extracted language tag before adding to the set
			// This will convert "en-US" to "en", "ko-KR" to "ko"
			langSet[normalizeLangCode(match[1])] = true
		}
	}

	langs := make([]string, 0, len(langSet))
	for lang := range langSet {
		langs = append(langs, lang)
	}

	return langs, nil
}

func getVoiceServicesLanguages() ([]string, error) {
	cmd := exec.Command("defaults", "read", "com.apple.voiceservices")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	langSet := make(map[string]bool)
	// Regex to find Languages = ( "lang-CODE" ); blocks
	// More flexible regex to handle varying whitespace and quotes
	re := regexp.MustCompile(`Languages\s*=\s*\(\s*"([a-zA-Z\-]+)"\s*\);`)
	matches := re.FindAllStringSubmatch(string(output), -1)

	for _, match := range matches {
		if len(match) > 1 {
			langSet[normalizeLangCode(match[1])] = true
		}
	}

	langs := make([]string, 0, len(langSet))
	for lang := range langSet {
		langs = append(langs, lang)
	}

	return langs, nil
}
