//go:build linux

package keyloc

import (
	"os/exec"
	"strings"
)

// mapLayoutToLangCode maps common X11 keyboard layouts to ISO 639-1 language codes.
func mapLayoutToLangCode(layout string) string {
	switch layout {
	case "us", "gb", "ca", "au":
		return "en"
	case "kr":
		return "ko"
	case "ru":
		return "ru"
	case "jp":
		return "ja"
	case "cn":
		return "zh"
	case "de":
		return "de"
	case "fr":
		return "fr"
	case "es":
		return "es"
	// Add more mappings as needed
	default:
		return layout // Return the original layout if no mapping is found
	}
}

func getLanguages() ([]string, error) {
	// localectl often provides more reliable layout info than environment variables
	cmd := exec.Command("localectl", "status")
	output, err := cmd.Output()
	if err != nil {
		// Fallback for systems without systemd/localectl
		cmd = exec.Command("setxkbmap", "-query")
		output, err = cmd.Output()
		if err != nil {
			return nil, err
		}
	}

	langSet := make(map[string]bool)
	outputStr := string(output)

	// Parse localectl or setxkbmap output
	lines := strings.Split(outputStr, "\n")
	for _, line := range lines {
		// For localectl: "X11 Layout: us,kr"
		// For setxkbmap: "layout:     us,kr"
		if strings.Contains(line, "Layout:") || strings.Contains(line, "layout:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				layouts := strings.Split(strings.TrimSpace(parts[1]), ",")
				for _, layout := range layouts {
					langSet[mapLayoutToLangCode(layout)] = true
				}
				break // Assume the first layout line is the most relevant
			}
		}
	}

	langs := make([]string, 0, len(langSet))
	for lang := range langSet {
		langs = append(langs, lang)
	}

	return langs, nil
}
