package utils

import "strings"

func CleanInput(text string) []string {
	arr := strings.Split(strings.TrimSpace(text), " ")
	var cleaned []string

	for _, c := range arr {
		if c != "" {
			cleaned = append(cleaned, strings.ToLower(c))
		}
	}

	return cleaned
}
