package main

import (
	"fmt"
	"github.com/sweb/wordcloud/words"
)

// Defines lower bound of word count to be displayed
var wordThreshold = 2

func main() {
	fmt.Println("###############################################################")
	importedWords, err := words.PutFileContentInSlice("data/words.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	topWords, err := words.PutFileContentInSlice("data/top10000de.txt")
	if err != nil {
		fmt.Println("Unable to load frequency list: " + err.Error())
	}
	fmt.Println("###############################################################")
	wordMap := words.CountOfWords(importedWords, topWords)

	sortedWordList := words.SortMapByValue(wordMap)

	fmt.Printf("Number of words: %d\n", +len(importedWords))
	fmt.Printf("Number of filtered distinct words: %d\n", +len(sortedWordList))
	for _, wordStat := range sortedWordList {
		if wordStat.Value >= wordThreshold {
			fmt.Printf("%v - %d\n", wordStat.Key, wordStat.Value)
		}

	}
}
