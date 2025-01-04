package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello  World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello world goodbye world",
			expected: []string{"hello", "world", "goodbye", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected: %d, got: %d", len(c.expected), len(actual))
			continue
		}

		for i, word := range actual {
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %s, got: %s", expectedWord, word)
			}
		}
	}
}
