package main

import (
	"fmt"
	"github.com/sweb/wordcloud/words"
)

// Defines lower bound of word count to be displayed
const wordThreshold = 2

// Defines maximum n of n-grams
const maxNGramSize = 4

// Filenames that are going to be imported
var files = []string{"data/words.txt", "data/words2.txt"}

func main() {
	fmt.Println("###############################################################")
	// Contains common words
	frequencyList, err := words.PutFileContentInSlice("data/top10000de.txt")
	if err != nil {
		fmt.Println("Unable to load frequency list: " + err.Error())
	}
	for _, filename := range files {
		fmt.Printf("Opening file %s...\n", filename)
		doc, err := words.NewDocument(filename, frequencyList, maxNGramSize,
			wordThreshold)

		if err != nil {
			fmt.Println("Error: " + err.Error())
		} else {
			fmt.Println("File processed!")
			fmt.Printf("Number of words: %d\n", +len(doc.ImportedWords))
			fmt.Printf("Number of filtered distinct words: %d\n",
				len(doc.DistinctiveWords))
			fmt.Printf("n-grams: (%d)\n", len(doc.NGramMap))
			for ngram, _ := range doc.NGramMap {
				fmt.Println("  " + ngram)
			}

			fmt.Printf("Distinctive words: (%d)\n", len(doc.SortedWordList))
			for _, wordStat := range doc.SortedWordList {
				if wordStat.Value >= wordThreshold {
					fmt.Printf("  %s - %d\n", wordStat.Key, wordStat.Value)
				}
			}
			fmt.Println("###############################################################")
		}
	}

	fmt.Println("End of application")

}
