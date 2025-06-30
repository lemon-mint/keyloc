package main

import (
	"fmt"

	"github.com/lemon-mint/keyloc"
)

func main() {
	// Example: Get the list of supported languages
	langs, err := keyloc.GetLanguages()
	if err != nil {
		fmt.Printf("Error getting languages: %v\n", err)
		return
	}

	fmt.Println("Supported keyboard languages or input sources:")
	for i, lang := range langs {
		fmt.Printf("%d. %s\n", i+1, lang)
	}
}
