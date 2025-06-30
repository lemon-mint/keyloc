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
	case strings.Contains(lowerIdentifier, "russian"):
		return "ru"
	case strings.Contains(lowerIdentifier, "japanese"):
		return "ja"
	case strings.Contains(lowerIdentifier, "french"):
		return "fr"
	case strings.Contains(lowerIdentifier, "german"):
		return "de"
	case strings.Contains(lowerIdentifier, "spanish"):
		return "es"
	case strings.Contains(lowerIdentifier, "chinese"):
		return "zh"
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