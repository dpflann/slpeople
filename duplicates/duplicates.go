package duplicates

import (
	"sort"
	"strings"
	"unicode/utf8"

	chars "github.com/slpeople/characters"
)

type (
	PossibleDuplicates [][]string
	thresholdSettings  struct {
		distanceThreshold int
		lengthThreshold   int
	}
)

/*** Level 3: Duplicate Email Addresses ***/
// References:
//  >> https://stackoverflow.com/questions/577463/finding-how-similar-two-strings-are
//  >> https://github.com/agnivade/levenshtein/blob/master/levenshtein.go
//  >> https://gist.github.com/andrei-m/982927#gistcomment-1931258

// FindPossbileDuplicates will search for possible duplicate strings that were for instance
// generated due to a typo upon input. The provide slice of strings is first grouped
// into slices that contain the same characters.
// Then these groups are then analzyed to see which strings could be duplicates
// of each other. Each string is compared to every other string. During comparison
// if the lengths of the two strings are equal or one is greater than the other
// plus a threshold value, then those strings will then have their Levenshtein
// distance computed. If the Levenshtein distance is less than threshold, then
// the two strings are considered possible duplicates and candidate for review.
// At the point the strings have the same characters, are similar length, and do
// not require many operations to convert one string to the other.
func FindPossibleDuplicates(strs []string, settings thresholdSettings) PossibleDuplicates {
	anagrams := make(map[string][]string)
	for _, s := range strs {
		charFreq := chars.CharacterFrequencyCount(s, nil)
		uniqueChars := make([]string, len(charFreq))
		for ch := range charFreq {
			uniqueChars = append(uniqueChars, ch)
		}
		sort.Strings(uniqueChars)
		uniqueCharsStr := strings.Join(uniqueChars, "")
		if dupeList, ok := anagrams[uniqueCharsStr]; ok {
			anagrams[uniqueCharsStr] = append(dupeList, s)
		} else {
			anagrams[uniqueCharsStr] = []string{s}
		}
	}
	duplicates := PossibleDuplicates{}
	for _, dupeList := range anagrams {
		// If there is only 1 string, then skip it.
		if len(dupeList) < 2 {
			continue
		}
		for i := 0; i < len(strs); i++ {
			dupes := []string{strs[i]}
			for j := i + 1; j < len(strs); j++ {
				if !compareLengths(strs[i], strs[j], settings.lengthThreshold) {
					continue
				}
				if ComputeDistance(strs[i], strs[j]) > settings.distanceThreshold {
					continue
				}
				dupes = append(dupes, strs[j])
			}
			if len(dupes) > 1 {
				duplicates = append(duplicates, dupes)
			}
		}

	}

	return duplicates
}

func compareLengths(str1, str2 string, threshold int) bool {
	if len(str1) == len(str2) {
		return true
	}
	if len(str1) > len(str2) {
		return (len(str1) - len(str2) - threshold) == 0
	}
	return (len(str2) - len(str1) - threshold) == 0
}

/***
    NB: Use the concept of edit distance and the Levenshtein function.
    The following code was pulled from github, it seemed like a reasonable
    thing to do, besides write it or find and import a third party.

    Source: https://github.com/agnivade/levenshtein/blob/master/levenshtein.go

    Note that this source is a go version of the following javascript implementation
    of the Levenshtein algorithm which is here:
        >> https://gist.github.com/andrei-m/982927#gistcomment-1931258
***/

// ComputeDistance computes the levenshtein distance between the two
// strings passed as an argument. The return value is the levenshtein distance
//
// Works on runes (Unicode code points) but does not normalize
// the input strings. See https://blog.golang.org/normalization
// and the golang.org/x/text/unicode/norm pacage.
// Source: https://github.com/agnivade/levenshtein/blob/master/levenshtein.go
func ComputeDistance(a, b string) int {
	if len(a) == 0 {
		return utf8.RuneCountInString(b)
	}

	if len(b) == 0 {
		return utf8.RuneCountInString(a)
	}

	if a == b {
		return 0
	}

	// We need to convert to []rune if the strings are non-ascii.
	// This could be avoided by using utf8.RuneCountInString
	// and then doing some juggling with rune indices.
	// The primary challenge is keeping track of the previous rune.
	// With a range loop, its not that easy. And with a for-loop
	// we need to keep track of the inter-rune width using utf8.DecodeRuneInString
	s1 := []rune(a)
	s2 := []rune(b)

	// swap to save some memory O(min(a,b)) instead of O(a)
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	lenS1 := len(s1)
	lenS2 := len(s2)

	// init the row
	x := make([]int, lenS1+1)
	for i := 0; i <= lenS1; i++ {
		x[i] = i
	}

	// fill in the rest
	for i := 1; i <= lenS2; i++ {
		prev := i
		var current int

		for j := 1; j <= lenS1; j++ {

			if s2[i-1] == s1[j-1] {
				current = x[j-1] // match
			} else {
				current = min(min(x[j-1]+1, prev+1), x[j]+1)
			}
			x[j-1] = prev
			prev = current
		}
		x[lenS1] = prev
	}
	return x[lenS1]
}

// Source: https://github.com/agnivade/levenshtein/blob/master/levenshtein.go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
