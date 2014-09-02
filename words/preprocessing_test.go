package words

import (
	"reflect"
	"testing"
)

func TestCountOfWords(t *testing.T) {
	words := []string{"Dies", "ist", "ein", "Testsatz", "Noch", "ein", "Testsatz"}
	commonWords := []string{"ist", "ein"}

	actual := CountOfWords(words, commonWords)
	expected := make(map[string]int)
	expected["dies"] = 1
	expected["testsatz"] = 2
	expected["noch"] = 1

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v was expected to be %v", actual, expected)
	}
}

func TestSplitByWhitespaces(t *testing.T) {
	words := "Dies ist ein Testsatz. Noch ein Testsatz."
	actual := splitByWhitespaces(words)
	expected := []string{"Dies", "ist", "ein", "Testsatz", "Noch", "ein", "Testsatz"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v was expected to be %v", actual, expected)
	}
}
