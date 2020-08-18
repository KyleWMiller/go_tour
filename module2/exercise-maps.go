package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s) 
	sentenceMap := make(map[string]int)

	for _, v := range words {
				
		if _, ok := sentenceMap[v]; ok {
			sentenceMap[v]++
		} else {
			sentenceMap[v] = 1
		}
	}

	return sentenceMap
}

func main() {
	wc.Test(WordCount)
}
