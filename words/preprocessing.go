// This file stores functions used to pre-process raw content of input-data.
// An example is the split of input-strings into a slice of words.

package words

import (
	"io/ioutil"
	"regexp"
	"strings"
)

var whiteSpaces = regexp.MustCompile(`[\s|.|,|"|;|:|\(|\)|\-|\?|!]`)
var validWord = regexp.MustCompile(`\w`)

// Reads data from a file and splits its content in single words
func PutFileContentInSlice(filename string) ([]string, error) {
	// Read file
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// Cast to string and split by defined whitespaces
	words := splitByWhitespaces(string(body))

	return words, nil
}

// Transforms the list of words into a map that counts the number of appearences
// of each listed word. Afterwards, words that are defined as common are removed
// from the resulting map.
func CountOfWords(words []string, commonWords []string) map[string]int {
	var wordMap = make(map[string]int)
	for _, word := range words {
		// Listing and comparison is done in lower case
		lowerCaseWord := strings.ToLower(word)
		// If the word is not listed yet, initialize it with a counter of 1,
		// else increase the counter by 1.
		wordMap[lowerCaseWord]++
	}
	// Remove common words
	for _, word := range commonWords {
		delete(wordMap, strings.ToLower(word))
	}
	return wordMap
}

// Splits the words of the input string by detecting whitespaces and other
// non-word characters.
func splitByWhitespaces(words string) []string {
	var wordSlice = make([]string, 0)
	// Currently processed word
	var currWord = ""
	// Loop through all characters and decide if this is still the current word
	// or if a whitespace appeared. In case of a whitespace, append the current
	// word to the slice and reset clear the current word to continue.
	for _, char := range words {
		if whiteSpaces.MatchString(string(char)) {
			// Only add the current word if it contains a valid word
			if validWord.MatchString(currWord) {
				currWord = strings.TrimSpace(currWord)
				wordSlice = append(wordSlice, currWord)
				currWord = ""
			}
		} else {
			currWord = currWord + string(char)
		}
	}
	// Handling of the last word after the string processed.
	if validWord.MatchString(currWord) {
		currWord = strings.TrimSpace(currWord)
		wordSlice = append(wordSlice, currWord)
	}
	return wordSlice
}
