package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type WordStat struct {
	Word  string
	Count int
}

type TextStats struct {
	charCount  int
	wordCount  int
	lineCount  int
	spaceCount int
	wordFreq   map[string]int
}

func infoFile(text string) TextStats {
	var stats TextStats
	stats.wordFreq = make(map[string]int)

	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		line := scanner.Text()
		stats.lineCount++
		stats.charCount += len(line)
		stats.spaceCount += strings.Count(line, " ")

		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word)
			if word == "" {
				continue
			}
			stats.wordCount++
			stats.wordFreq[word]++
		}
	}
	return stats
}

func seeStats(stats TextStats) {
	fmt.Printf("Символов: %d\n", stats.charCount)
	fmt.Printf("Слов:     %d\n", stats.wordCount)
	fmt.Printf("Строк:    %d\n", stats.lineCount)
	fmt.Printf("Пробелов: %d\n", stats.spaceCount)
}

func getTopWords(wordFreq map[string]int, topN int) []WordStat {
	stats := make([]WordStat, 0, len(wordFreq))

	for word, count := range wordFreq {
		stats = append(stats, WordStat{Word: word, Count: count})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[j].Count
	})

	if topN > len(stats) {
		topN = len(stats)
	}
	return stats[:topN]
}

func main() {
	/*wordsCount := make(map[string]int)
	var stats []WordStat
	for word, count := range wordsCount {
		stats = append(stats, WordStat{Word: word, Count: count})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[i].Count
	})*/

	text := "Я не знаю, что сказать тебе при встрече Не могу найти хотя бы пары слов А недолгий вечер, а недолгий вечер Скоро станет ночью тёмною без снов А недолгий вечер, а недолгий вечер Станет ночью тёмною без снов"

	stats := infoFile(text)
	seeStats(stats)

	top := getTopWords(stats.wordFreq, 5)

	fmt.Println("\nТоп слов")
	for i, ws := range top {
		fmt.Printf("%d. %s: %d\n", i+1, ws.Word, ws.Count)
	}
}
