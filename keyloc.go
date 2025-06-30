package keyloc

import "strings"

// normalizeLangCode converts a language tag into a consistent, basic format.
// e.g., "en-US", "en_GB", "EN" all become "en".
func normalizeLangCode(lang string) string {
	lower := strings.ToLower(lang)
	normalized := strings.ReplaceAll(lower, "_", "-")
	parts := strings.Split(normalized, "-")
	return parts[0]
}

func CheckLanguage(lang string) (bool, error) {
	langs, err := getLanguages()
	if err != nil {
		return false, err
	}

	normalizedInput := normalizeLangCode(lang)

	for _, l := range langs {
		if normalizeLangCode(l) == normalizedInput {
			return true, nil
		}
	}
	return false, nil
}