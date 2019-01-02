package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCharacterFrequency(t *testing.T) {
	emailAddressTestData := []struct {
		address  string
		expected CharacterFrequencies
	}{
		{
			address: "abc@def.ghi",
			expected: CharacterFrequencies{
				"a": 1,
				"b": 1,
				"c": 1,
				"d": 1,
				"e": 1,
				"f": 1,
				"g": 1,
				"h": 1,
				"i": 1,
				"@": 1,
				".": 1,
			},
		},
	}
	for _, td := range emailAddressTestData {
		if result := CharacterFrequencyCount(td.address, nil); !cmp.Equal(result, td.expected) {
			t.Fatalf("The resultant frequency map is not equal to the expected frequency map: \n\taddress: %s, \n\tresult: %#v, \n\texpected: %#v\n", td.address, result, td.expected)
		}
	}
}

func TestSortedCharFrequency(t *testing.T) {
	emailAddressTestData := []struct {
		address  string
		expected SortedCharFreqs
	}{
		{
			address: "aaaabbb@cc.d",
			expected: SortedCharFreqs{
				{"a", 4},
				{"b", 3},
				{"c", 2},
				{"d", 1},
			},
		},
	}
	for _, td := range emailAddressTestData {
		charFreqs := CharacterFrequencyCount(td.address, blackList)
		result := *((&charFreqs).Sorted())
		if !cmp.Equal(result, td.expected) {
			t.Fatalf("The resultant frequency map is not equal to the expected frequency map: \n\taddress: %s, \n\tresult: %#v, \n\texpected: %#v\n", td.address, result, td.expected)
		}
	}
}

func TestCharacterFrequencies(t *testing.T) {
	emailAddressesTestData := []struct {
		addresses []string
		expected  CharacterFrequencies
	}{
		{
			addresses: []string{"abc@def.ghi", "ghi@def.abc"},
			expected: map[string]int{
				"a": 2,
				"b": 2,
				"c": 2,
				"d": 2,
				"e": 2,
				"f": 2,
				"g": 2,
				"h": 2,
				"i": 2,
				"@": 2,
				".": 2,
			},
		},
	}
	for _, td := range emailAddressesTestData {
		if result := CharacterFrequencyCountOfStrings(td.addresses, nil); !cmp.Equal(result, td.expected) {
			t.Fatalf("The resultant frequency map is not equal to the expected frequency map: \n\taddress: %s, \n\tresult: %#v, \n\texpected: %#v\n", td.addresses, result, td.expected)
		}
	}
}
