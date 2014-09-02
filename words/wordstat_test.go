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
