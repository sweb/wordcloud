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

func TestCreateNGrams(t *testing.T) {
	words := []string{"Dies", "ist", "ein", "Testsatz", "Noch", "ein", "Testsatz"}
	n := 2
	actual, _ := CreateNGrams(words, n)
	expected := []NGram{NGram{"Dies", "ist"}, NGram{"ist", "ein"}, NGram{"ein", "Testsatz"},
		NGram{"Testsatz", "Noch"}, NGram{"Noch", "ein"}, NGram{"ein", "Testsatz"}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v was expected to be %v", actual, expected)
	}

	n = 3
	actual, _ = CreateNGrams(words, n)
	expected = []NGram{NGram{"Dies", "ist", "ein"}, NGram{"ist", "ein", "Testsatz"},
		NGram{"ein", "Testsatz", "Noch"}, NGram{"Testsatz", "Noch", "ein"},
		NGram{"Noch", "ein", "Testsatz"}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v was expected to be %v", actual, expected)
	}

	n = -1
	_, actualErr := CreateNGrams(words, n)
	expectedErr := "n of n-gram needs to be > 0"

	if actualErr.Error() != expectedErr {
		t.Errorf("%v was expected to be %v", actualErr.Error(), expectedErr)
	}
}
