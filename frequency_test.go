package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCharacterFrequency(t *testing.T) {
	emailAddressTestData := []struct {
		address  string
		expected map[string]int
	}{
		{
			address: "abc@def.ghi",
			expected: map[string]int{
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
		if result := CharacterFrequencyCount(td.address); !cmp.Equal(result, td.expected) {
			t.Fatalf("The resultant frequency map is not equal to the expected frequency map: \n\taddress: %s, \n\tresult: %#v, \n\texpected: %#v\n", td.address, result, td.expected)
		}
	}
}

func TestCharacterFrequencies(t *testing.T) {
	emailAddressesTestData := []struct {
		addresses []string
		expected  map[string]int
	}{
		{
			addresses: []string{"abc@def.ghi"},
			expected: map[string]int{
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
	for _, td := range emailAddressesTestData {
		if result := CharacterFrequencyCountOfStrings(td.addresses); !cmp.Equal(result, td.expected) {
			t.Fatalf("The resultant frequency map is not equal to the expected frequency map: \n\taddress: %s, \n\tresult: %#v, \n\texpected: %#v\n", td.addresses, result, td.expected)
		}
	}
}
