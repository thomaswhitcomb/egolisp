package main

import "testing"

func TestSexpr(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected bool
	}{
		{"1", makeTrue().isTrue(), true},
		{"2", makeList(nil).isList(), true},
		{"3", makeSymbol(new(string)).isScalar(), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input != tt.expected {
				t.Errorf("TestSexpr - Test %s failed ", tt.name)
			}
		})
	}
}

var byebye string = "byebye"

func TestSexprShow(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"1", *makeTrue().show(), "#t"},
		{"2", *makeFalse().show(), "#f"},
		{"3", *makeSymbol(&byebye).show(), "byebye"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input != tt.expected {
				t.Errorf("TestSexprShow - Test %s failed.  Expected %s ", tt.name, tt.expected)
			}
		})
	}
}

var v string = "v"
var n int = 123
var k string = ":dog"

func TestSexprTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected bool
	}{
		{"1", makeSymbol(&v).isSymbol(), true},
		{"2", makeSymbol(&v).isNumber(), false},
		{"3", makeNumber(&n).isSymbol(), false},
		{"4", makeSymbol(&v).isSymbol(), true},
		{"5", makeKeyword(&k).isSymbol(), false},
		{"6", makeKeyword(&k).isNumber(), false},
		{"7", makeKeyword(&k).isKeyword(), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input != tt.expected {
				t.Errorf("TestSexprShow - Test %s failed.  Expected %t ", tt.name, tt.expected)
			}
		})
	}
}
