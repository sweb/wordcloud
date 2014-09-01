// This file stores functions used to pre-process raw content of input-data.
// An example is the split of input-strings into a slice of words.

package words

import (
	"errors"
	"fmt"
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
	fmt.Printf("Success: File %v loaded!\n", filename)
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
		currWordCount := wordMap[lowerCaseWord]
		if currWordCount == 0 {
			wordMap[lowerCaseWord] = 1
		} else {
			wordMap[lowerCaseWord] = currWordCount + 1
		}
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

// This function creates n-grams based on a given set of words and the size of
// n. n needs to be greater than 0.
func CreateNGrams(words []string, n int) ([]NGram, error) {
	if n < 0 {
		return nil, errors.New("n of n-gram needs to be > 0")
	}
	// n reduces total number of n-grams but even if n equals the number of
	// words, one n-gram is created.
	numberOfNGrams := len(words) - (n - 1)
	nGramList := make([]NGram, numberOfNGrams)
	for i := 0; i < numberOfNGrams; i++ {
		nGram := make(NGram, n)
		for j := 0; j < n; j++ {
			nGram[j] = words[i+j]
		}
		nGramList[i] = nGram
	}
	return nGramList, nil
}
