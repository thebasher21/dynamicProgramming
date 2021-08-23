package wordConstruct

import (
	"strings"
)

// CanConstruct() returns a boolean whether the given target word can be constructed using the words in the word bank
func CanConstructRec(targetWord string, wordBank []string, memo map[string]bool) bool {
	returnObj, isPresent := memo[targetWord]
	if isPresent {
		return returnObj
	}
	if targetWord == "" {
		return true
	}
	for _, word := range wordBank {
		if strings.Index(targetWord, word) == 0 {
			suffix := targetWord[len(word):]
			if CanConstructRec(suffix, wordBank, memo) {
				memo[targetWord] = true
				return true
			}
		}
	}
	memo[targetWord] = false
	return false
}

// AllConstruct() returns all the ways the given target word can be constructed using the words in the word bank
func AllConstructRec(targetWord string, wordBank []string, memo map[string][][]string) [][]string {
	returnObj, isPresent := memo[targetWord]
	if isPresent {
		return returnObj
	}
	if targetWord == "" {
		return [][]string{{}}
	}
	result := [][]string{}
	for _, word := range wordBank {
		if strings.Index(targetWord, word) == 0 {
			suffix := targetWord[len(word):]
			//fmt.Println(word, suffix)
			suffixWays := AllConstructRec(suffix, wordBank, memo)
			for index := range suffixWays {
				result = append(result, append([]string{word}, suffixWays[index]...))
			}
		}
		//fmt.Println(result)
	}
	memo[targetWord] = result
	return result
}

// CountConstruct() returns a count of how many ways the given target word can be constructed using the words in the word bank
func CountConstructRec(targetWord string, wordBank []string, memo map[string]int) int {
	returnObj, isPresent := memo[targetWord]
	if isPresent {
		return returnObj
	}
	if targetWord == "" {
		return 1
	}
	totalCount := 0
	for _, word := range wordBank {
		if strings.Index(targetWord, word) == 0 {
			suffix := targetWord[len(word):]
			numWays := CountConstructRec(suffix, wordBank, memo)
			totalCount += numWays
		}
	}
	memo[targetWord] = totalCount
	return totalCount
}

/* func CanConstructIter(targetWord string, wordBank []string) bool {
	if len(targetWord) == 0 {
		return true
	}
	for _, word := range wordBank {
		if strings.Index(targetWord, word) == 0 {
			suffix := targetWord[len(word):]
		}
	}
	return false
} */
