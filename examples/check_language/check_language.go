package main

import (
	"fmt"

	"github.com/lemon-mint/keyloc"
)

func main() {
	// Example: Check if a specific language is supported
	lang := "en"
	supported, err := keyloc.CheckLanguage(lang)
	if err != nil {
		fmt.Printf("Error checking language: %v\n", err)
		return
	}
	if supported {
		fmt.Printf("Language '%s' is supported as a keyboard input.\n", lang)
	} else {
		fmt.Printf("Language '%s' is not supported as a keyboard input.\n", lang)
	}
}
