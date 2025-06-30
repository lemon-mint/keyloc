# Keyloc

**Keyloc** is a Go library designed to detect and check available keyboard languages or input sources on a system. It supports multiple platforms, including Windows, macOS, and Linux, using platform-specific methods to retrieve language information. The library normalizes language codes for consistent comparison, making it easy to determine if a specific language is supported as a keyboard input.

## Features

- **Cross-Platform Support**: Works on Windows, macOS, and Linux with tailored implementations for each operating system.
- **Language Normalization**: Converts various language code formats (e.g., "en-US", "en_GB") into a consistent base format (e.g., "en").
- **Simple API**: Provides easy-to-use functions to check if a language is available and to retrieve the full list of supported keyboard languages or input sources.

## Installation

To use Keyloc in your Go project, you can add it as a dependency:

```bash
go get github.com/lemon-mint/keyloc
```

## Usage

Keyloc provides two main functions to interact with keyboard language information:

### Checking a Specific Language

Here's a basic example of how to use Keyloc to check if a specific language is available:

```go
package main

import (
	"fmt"
	"github.com/lemon-mint/keyloc"
)

func main() {
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
```

### Listing All Supported Languages

You can also retrieve the full list of supported keyboard languages or input sources on the system:

```go
package main

import (
	"fmt"
	"github.com/lemon-mint/keyloc"
)

func main() {
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
```

### Running Examples

Example files are provided in the `examples` directory. To run an example, navigate to the specific file and execute it individually:

```bash
go run examples/check_language.go
```

or

```bash
go run examples/list_languages.go
```

## How It Works

- **macOS**: Queries system preferences for enabled input sources, preferred languages, and voice services to build a list of language codes.
- **Windows**: Uses system calls to retrieve keyboard layout information and maps Windows language IDs (LCIDs) to standard language codes.
- **Linux**: (Implementation details available in the source code for Linux-specific handling.)

## Requirements

- **Go Version**: 1.24.4 or compatible versions.

## Contributing

Contributions are welcome! If you have suggestions, bug reports, or want to add support for additional platforms or features, please open an issue or submit a pull request on the [GitHub repository](https://github.com/lemon-mint/keyloc).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Developed by [lemon-mint](https://github.com/lemon-mint).
