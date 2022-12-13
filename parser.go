package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Pattern is a regular expression.
type Pattern string

// NewPattern returns a new Pattern.
func NewPattern(pattern string) (Pattern, error) {
	if !strings.HasPrefix(pattern, "/") || !strings.HasSuffix(pattern, "/") {
		return "", fmt.Errorf("invalid pattern")
	}

	// remove the slashes
	pattern = pattern[1 : len(pattern)-1]

	return Pattern(pattern), nil
}

// String returns the string representation of the pattern.
func (p Pattern) String() string {
	return string(p)
}

// Generate generates a regular expression from the pattern.
func (p Pattern) Generate() string {
	newString := ""

	tokens := regexp.MustCompile(`(\[[^\]]+\]|\{[^\}]+\})|.`).FindAllString(p.String(), -1)

	for len(tokens) > 0 {
		i := len(tokens) - 1
		repeat := 1
		token := tokens[i]
		if isQuantifier(token) {
			repeat = getRepeat(token)
			token = tokens[i-1]
			tokens = tokens[:i-1]
		} else {
			tokens = tokens[:i]
		}

		newString = parseToken(token, repeat) + newString
	}

	return newString
}
