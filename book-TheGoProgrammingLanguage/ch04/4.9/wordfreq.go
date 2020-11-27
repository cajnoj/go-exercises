package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type wordToFrequencyMap map[string]int
type frequencyToWordsMap map[int]string

func groupByFrequency(wf wordToFrequencyMap) frequencyToWordsMap {

	freqWord := make(frequencyToWordsMap)

	for w, f := range wf {
		freqWord[f] = fmt.Sprintf("%s, %q", freqWord[f], w)
	}

	return freqWord

}

func getKeysSorted(m frequencyToWordsMap) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	return keys
}

func countWords(r io.Reader) wordToFrequencyMap {
	wordFreq := make(wordToFrequencyMap)

	input := bufio.NewScanner(r)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordFreq[input.Text()]++
	}

	return wordFreq
}

func main() {
	// Scan input and keep word counts in map
	wordFreq := countWords(os.Stdin)

	// Group word-count words by frequencies
	freqWord := groupByFrequency(wordFreq)

	// Order frequencies
	frequencies := getKeysSorted(freqWord)

	// Print frequency-to-word-list map contents ordered by frequencies
	for _, f := range frequencies {
		fmt.Printf("%10d: %s\n", f, freqWord[f])
	}
}
