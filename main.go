package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// normalizeString trims, converts to lowercase, and removes all non-alphanumeric characters from a string
func normalizeString(str string) string {
	var normalized []rune
	for _, r := range str {
		if r >= 'a' && r <= 'z' {
			normalized = append(normalized, r)
		} else if r >= 'A' && r <= 'Z' {
			normalized = append(normalized, r)
		} else if r >= '0' && r <= '9' {
			normalized = append(normalized, r)
		}
	}
	return strings.ToLower(string(normalized))
}

// extractStrings parses a file and returns a map of all unique strings
func extractStrings(filePath string) (map[string]bool, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	uniqueStrings := make(map[string]bool)
	for _, token := range strings.Fields(string(data)) {
		token = normalizeString(token)
		uniqueStrings[token] = true
	}
	return uniqueStrings, nil
}

// collectUniqueStrings walks a directory and returns a map of all unique strings
func collectUniqueStrings(dir string) (map[string]bool, error) {
	uniqueStrings := make(map[string]bool)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if strings.HasSuffix(path, ".md") {
				fmt.Println(path)
				strs, err := extractStrings(path)
				if err != nil {
					return err
				}
				for str := range strs {
					uniqueStrings[str] = true
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return uniqueStrings, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: resonate <directory>")
		return
	}
	dir := os.Args[1]
	uniqueStrings, err := collectUniqueStrings(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Unique strings: %d\n", len(uniqueStrings))
}
