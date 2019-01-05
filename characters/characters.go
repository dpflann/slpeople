package characters

import "sort"

type (
	CharacterFrequencies               map[string]int
	SortedCharacterFrequenciesResponse struct {
		*SortedCharFreqs `json:"frequencies"`
	}
	KeyVal struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	}
	SortedCharFreqs []KeyVal
)

func (c *CharacterFrequencies) Sorted() *SortedCharFreqs {
	var cfs SortedCharFreqs
	for char, count := range *c {
		cfs = append(cfs, KeyVal{char, count})
	}

	sort.Slice(cfs, func(i, j int) bool {
		return cfs[i].Value > cfs[j].Value
	})
	return &cfs
}

func CharacterFrequencyCount(str string, blackList map[string]bool) CharacterFrequencies {
	frequencies := CharacterFrequencies{}
	for _, c := range str {
		cStr := string(c)
		if _, ok := blackList[cStr]; ok {
			continue
		}
		if _, ok := frequencies[cStr]; ok {
			frequencies[cStr] += 1
		} else {
			frequencies[cStr] = 1
		}
	}
	return frequencies
}

func CharacterFrequencyCountOfStrings(strs []string, blackList map[string]bool) CharacterFrequencies {
	// Naive handling
	frequencies := CharacterFrequencies{}
	for _, s := range strs {
		res := CharacterFrequencyCount(s, blackList)
		for c, v := range res {
			if _, ok := frequencies[c]; ok {
				frequencies[c] += v
			} else {
				frequencies[c] = v
			}
		}
	}
	return frequencies
}
