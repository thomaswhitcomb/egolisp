package main

import "testing"

var one string = "1"
var two string = "2"
var three string = "3"
var thirtyThreeInt int = 33
var abc string = "abc"
var kw string = ":abc"

func linkSexpr(s ...*cell) *cell {
	var head, curr *cell
	for _, c := range s {
		if head == nil {
			head = c
		} else {
			curr.next = c
		}
		curr = c
	}
	return head
}
func TestAtom(t *testing.T) {
	tests := []struct {
		name     string
		input    *cell
		expected *cell
	}{
		{"1", makeSymbol(&abc), makeTrue()},
		{"2", makeNumber(&thirtyThreeInt), makeTrue()},
		{"3", makeKeyword(&kw), makeTrue()},
		{"4", makeList(makeSymbol(&abc)), makeFalse()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := isatom(nil, tt.input)
			if *x.show() != *tt.expected.show() {
				t.Errorf("TestAtom - Test %s returned value %s and is not equal to %s ", tt.name, *x.show(), *tt.expected.show())
			}
		})
	}
}
func TestCar(t *testing.T) {
	tests := []struct {
		name     string
		input    *cell
		expected *cell
	}{
		{"1", makeList(linkSexpr(makeStr(&one), makeStr(&two))), makeStr(&one)},
		{"2", makeList(makeStr(&one)), makeStr(&one)},
		{"3", makeList(nil), makeFalse()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := car(nil, tt.input)
			if *x.show() != *tt.expected.show() {
				t.Errorf("TestCar - Test %s not equal %s failed ", *x.show(), *tt.expected.show())
			}
		})
	}
}
func TestCdr(t *testing.T) {
	tests := []struct {
		name     string
		input    *cell
		expected *cell
	}{
		{"1", makeList(linkSexpr(makeStr(&one), makeStr(&two))), makeList(makeStr(&two))},
		{"2", makeList(makeStr(&one)), makeList(nil)},
		//{"3", makeList(&(celllist{})), makeFalse()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := cdr(nil, tt.input)
			if *x.show() != *tt.expected.show() {
				t.Errorf("TestCdr - Test %s not equal %s failed ", *x.show(), *tt.expected.show())
			}
		})
	}
}
func TestEQ(t *testing.T) {
	var twelve int = 12
	var twentythree int = 23
	tests := []struct {
		name     string
		input1   *cell
		expected *cell
	}{
		{"1", linkSexpr(makeNumber(&twelve), makeNumber(&twentythree)), makeFalse()},
		{"2", linkSexpr(makeNumber(&twelve), makeNumber(&twelve)), makeTrue()},
		{"3", linkSexpr(makeList(nil), makeList(nil)), makeTrue()},
		{"4", linkSexpr(makeList(makeStr(&one)), makeList(nil)), makeFalse()},
		{"5", linkSexpr(makeList(makeStr(&one)), makeList(makeStr(&one))), makeTrue()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := isequal(nil, tt.input1)
			if *x.show() != *tt.expected.show() {
				t.Errorf("TestEQ - Test %s not equal %s failed ", *x.show(), *tt.expected.show())
			}
		})
	}
}
