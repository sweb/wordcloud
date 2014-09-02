package words

type Document struct {
	Filename         string
	ImportedWords    []string
	DistinctiveWords map[string]int
	NGramMap         NGramMap
	SortedWordList   WordStatList
}

func NewDocument(filename string, frequencyList []string, maxNGramSize int,
	wordThreshold int) (*Document, error) {
	importedWords, err := PutFileContentInSlice(filename)
	if err != nil {
		return nil, err
	}
	distinctiveWords := CountOfWords(importedWords, frequencyList)
	ngramMap := make(NGramMap)
	err = ngramMap.CombineDistinctiveNGrams(importedWords, maxNGramSize,
		distinctiveWords, wordThreshold)
	if err != nil {
		return nil, err
	}
	sortedWordList := SortMapByValue(distinctiveWords)
	doc := Document{filename, importedWords, distinctiveWords, ngramMap, sortedWordList}
	return &doc, nil
}
