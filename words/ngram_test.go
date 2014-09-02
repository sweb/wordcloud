package words

import (
	"reflect"
	"testing"
)

type nGramKeyTest struct {
	ng       NGram
	expected string
}

func TestNGramKey(t *testing.T) {
	testdata := []nGramKeyTest{
		nGramKeyTest{NGram{"dies"}, "dies"},
		nGramKeyTest{NGram{"dies", "ist"}, "dies ist"},
		nGramKeyTest{NGram{"dies", "ist", "ein", "test"}, "dies ist ein test"}}

	for _, input := range testdata {
		actual := input.ng.Key()
		if actual != input.expected {
			t.Errorf("%v was expected to be %v", actual, input.expected)
		}
	}

}

type nGramIsDistinctiveTest struct {
	ng            NGram
	distinctWords map[string]int
	expected      bool
}

func TestNGramIsDistinctive(t *testing.T) {
	distinctiveWords := make(map[string]int)
	distinctiveWords["test"] = 1
	distinctiveWords["fall"] = 1

	testdata := []nGramIsDistinctiveTest{
		nGramIsDistinctiveTest{NGram{"dies"}, distinctiveWords, false},
		nGramIsDistinctiveTest{NGram{"dies", "ist"}, distinctiveWords, false},
		nGramIsDistinctiveTest{NGram{"dies", "ist", "ein", "test"}, distinctiveWords, false},
		nGramIsDistinctiveTest{NGram{"test", "fall"}, distinctiveWords, true},
		nGramIsDistinctiveTest{NGram{"test", "fall", "ist"}, distinctiveWords, false}}

	for _, input := range testdata {
		actual := input.ng.IsDistinctive(input.distinctWords, 1)
		if actual != input.expected {
			t.Errorf("Case %v: %v was expected to be %v", input.ng, actual, input.expected)
		}
	}

}

func TestCreateNGrams(t *testing.T) {
	words := []string{"Dies", "ist", "ein", "Testsatz", "Noch", "ein", "Testsatz"}
	n := 2
	actual, _ := createNGrams(words, n)
	expected := []NGram{NGram{"Dies", "ist"}, NGram{"ist", "ein"}, NGram{"ein", "Testsatz"},
		NGram{"Testsatz", "Noch"}, NGram{"Noch", "ein"}, NGram{"ein", "Testsatz"}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v was expected to be %v", actual, expected)
	}

	n = 3
	actual, _ = createNGrams(words, n)
	expected = []NGram{NGram{"Dies", "ist", "ein"}, NGram{"ist", "ein", "Testsatz"},
		NGram{"ein", "Testsatz", "Noch"}, NGram{"Testsatz", "Noch", "ein"},
		NGram{"Noch", "ein", "Testsatz"}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v was expected to be %v", actual, expected)
	}

	n = -1
	_, actualErr := createNGrams(words, n)
	expectedErr := "n of n-gram needs to be > 0"

	if actualErr.Error() != expectedErr {
		t.Errorf("%v was expected to be %v", actualErr.Error(), expectedErr)
	}
}

type nGramCountOfNGramsTest struct {
	nGramMap      NGramMap
	words         []string
	n             int
	distinctWords map[string]int
	expected      NGramMap
}

func TestAddDistinctiveNGrams(t *testing.T) {
	expected1 := make(NGramMap)
	expected1["Test Fall"] = 1

	distinctiveWords := make(map[string]int)
	distinctiveWords["test"] = 1
	distinctiveWords["fall"] = 1

	testdata := []nGramCountOfNGramsTest{
		nGramCountOfNGramsTest{make(NGramMap), []string{"Dies", "ist", "ein", "Test", "Fall"},
			2, distinctiveWords, expected1},
		nGramCountOfNGramsTest{make(NGramMap), []string{"Dies", "ist", "ein", "Test", "ein", "Fall"},
			2, distinctiveWords, make(NGramMap)}}

	for _, input := range testdata {
		input.nGramMap.addDistinctiveNGrams(input.words, input.n, input.distinctWords, 1)
		actual := input.nGramMap
		if !reflect.DeepEqual(actual, input.expected) {
			t.Errorf("Case %v: %v was expected to be %v", input.words, actual, input.expected)
		}
	}
}
