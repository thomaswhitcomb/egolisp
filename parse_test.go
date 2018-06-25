package main

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"1", "nil", "(quote ())"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenize(tt.input)
			parser := newParser()
			parser.setTokens(tokens)
			cells := parser.parse()
			if *cells.show() != tt.expected {
				t.Errorf("TestParse - Started with %s, tokenized and expected %v", tt.input, *cells.show())
			}
		})
	}
}
