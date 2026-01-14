package main

import (
	"reflect"
	"testing"
)

func TestTopWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		k        int
		expected []string
	}{
		{
			name:     "K меньше уникальных слов",
			input:    "apple banana apple orange banana apple",
			k:        2,
			expected: []string{"apple", "banana"},
		},
		{
			name:     "K меньше уникальных слов",
			input:    "aa bb cc aa cc cc cc aa ab ac bb",
			k:        3,
			expected: []string{"cc", "aa", "bb"},
		},
		{
			name:     "K меньше уникальных слов",
			input:    "1 1 2 2 3 3 4 4 55",
			k:        3,
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "Пустой список слов",
			input:    "",
			k:        3,
			expected: []string{},
		},
		{
			name:     "Пустой список слов",
			input:    "",
			k:        0,
			expected: []string{},
		},
		{
			name:     "Пустой список слов",
			input:    "   ",
			k:        255,
			expected: []string{},
		},
		{
			name:     "K больше количества уникальных слов",
			input:    "dog cat dog",
			k:        10,
			expected: []string{"dog", "cat"},
		},
		{
			name:     "K больше количества уникальных слов",
			input:    "Борис Василий Максат Максат Василий Борис Василий ",
			k:        10,
			expected: []string{"Василий", "Борис", "Максат"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := topWords(test.input, test.k)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf(
					"expected %v, got %v",
					test.expected,
					result,
				)
			}
		})
	}
}
