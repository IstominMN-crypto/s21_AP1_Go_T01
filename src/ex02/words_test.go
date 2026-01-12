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
			name:     "a. обычное поведение, K меньше уникальных слов",
			input:    "apple banana apple orange banana apple",
			k:        2,
			expected: []string{"apple", "banana"},
		},
		{
			name:     "b. пустой список слов",
			input:    "",
			k:        3,
			expected: []string{},
		},
		{
			name:     "c. K больше количества уникальных слов",
			input:    "dog cat dog",
			k:        10,
			expected: []string{"dog", "cat"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TopWords(tt.input, tt.k)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf(
					"expected %v, got %v",
					tt.expected,
					result,
				)
			}
		})
	}
}
