package main

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// randomInt returns a random integer between min and max
func randomInt(min, max int) int {
	if min == max {
		return min
	}

	return min + rand.Intn(max-min)
}

// this function checks if the token is a quantifier
func isQuantifier(token string) bool {
	return token == "*" || token == "+" || token == "?" || (strings.HasPrefix(token, "{") && strings.HasSuffix(token, "}"))
}

// getRepeat returns the repeat value of a quantifier
func getRepeat(token string) int {
	if token == "*" {
		return randomInt(0, 10)
	} else if token == "+" {
		return randomInt(1, 10)
	} else if token == "?" {
		return randomInt(0, 1)
	} else {
		// remove the curly braces
		token = token[1 : len(token)-1]
		if !strings.Contains(token, ",") {
			num, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}

			return num
		}

		// split the string by the comma
		split := strings.Split(token, ",")
		if len(split) != 2 {
			panic("invalid quantifier")
		}

		// get the min and max values
		min, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		max, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		return randomInt(min, max)
	}
}

// getRange returns the range of a character class
func getRange(pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(allChars, -1)
}

// getNegatedRange returns the negated range of a character class
func getNegatedRange(pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(allChars, -1)
}

// getRandomInRange returns a random character from the range
func getRandomInRange(r []string) string {
	if len(r) == 0 {
		return ""
	}
	return r[randomInt(0, len(r))]
}

// repeatGenerate is a fibonacci
// r range of characters
// n the number of times to repeat
func repeatGenerate(r []string, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += getRandomInRange(r)
	}

	return s
}

// parseToken
func parseToken(token string, repeat int) string {
	r := getRange(token)
	val := repeatGenerate(r, repeat)
	return val
}
