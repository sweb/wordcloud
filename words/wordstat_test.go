package words

import "testing"

func TestSortMapByValue(t *testing.T) {
	testmap := make(map[string]int)
	testmap["t1"] = 5
	testmap["t2"] = 3
	testmap["t3"] = 1
	testmap["t4"] = 10

	expectedFirst := WordStat{Key: "t4", Value: 10}
	expectedLast := WordStat{Key: "t3", Value: 1}
	actual := SortMapByValue(testmap)

	if actual[0] != expectedFirst || actual[3] != expectedLast {
		t.Fail()
	}
}

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
		actual := input.ng.IsDistinctive(input.distinctWords)
		if actual != input.expected {
			t.Errorf("Case %v: %v was expected to be %v", input.ng, actual, input.expected)
		}
	}

}
