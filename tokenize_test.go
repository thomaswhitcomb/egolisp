package main

import "testing"

func TestTokenizer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"1", "1", []string{"1"}},
		{"2", "1 2 3", []string{"1", "2", "3"}},
		{"3", "(  1 (222 )   3)", []string{"(", "1", "(", "222", ")", "3", ")"}},
		{"4", "nil", []string{"nil"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenize(tt.input)
			if len(tokens) != len(tt.expected) {
				t.Errorf("TestTokenize - Started with %s, tokenized and expected %v", tt.input, len(tokens))
			}
			for i := 0; i < len(tokens); i++ {
				if tokens[i] != tt.expected[i] {
					t.Errorf("TestTokenize - Started with %s, tokenized and expected %v", tt.input, tokens)
				}
			}
		})
	}
}
