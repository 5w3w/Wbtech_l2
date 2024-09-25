package main

import (
	"fmt"
	"sort"
	"strings"
)

// Функция для сортировки букв в слове
func sortedWord(word string) string {
	letters := strings.Split(word, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

// Функция поиска множеств анаграмм
func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)
	wordSeen := make(map[string]bool)

	// Приведение всех слов к нижнему регистру и сортировка букв
	for _, word := range words {
		wordLower := strings.ToLower(word)
		sorted := sortedWord(wordLower)

		anagramMap[sorted] = append(anagramMap[sorted], wordLower)
	}

	result := make(map[string][]string)

	// Фильтрация множеств, которые состоят из более чем одного слова
	for _, anagramSet := range anagramMap {
		if len(anagramSet) > 1 {
			sort.Strings(anagramSet)

			// Находим первое встретившееся слово, которого ещё не было в результирующем множестве
			for _, word := range anagramSet {
				if !wordSeen[word] {
					result[word] = anagramSet
					for _, w := range anagramSet {
						wordSeen[w] = true
					}
					break
				}
			}
		}
	}

	return result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток", "кто"}
	anagrams := findAnagrams(words)

	for key, anagramSet := range anagrams {
		fmt.Printf("%s: %v\n", key, anagramSet)
	}
}
