package main

import (
	"fmt"
	"os"
)

var TRUE string = "#t"
var FALSE string = "#f"
var nilconst string = "nil"
var lambdastr string = "(lambda)"

const (
	list         = iota
	fn           = iota
	number       = iota
	keyword      = iota
	string_      = iota
	symbol       = iota
	nil_         = iota
	true_        = iota
	false_       = iota
	dotted_pair_ = iota
)

type cell struct {
	type_ int8
	value interface{}
	next  *cell
}
type dottedPair struct {
	car *cell
	cdr *cell
}

func makeDottedPair(c1 *cell, c2 *cell) *cell {
	return &cell{type_: dotted_pair_, value: &dottedPair{car: c1, cdr: c2}, next: nil}
}
func (p *cell) isDottedPair() bool {
	return p.type_ == dotted_pair_
}
func (p *cell) dottedPair() (*cell, *cell) {
	dp := (p.value).(*dottedPair)
	return dp.car, dp.cdr
}
func makeList(ss *cell) *cell {
	return &cell{type_: list, value: ss, next: nil}
}

func (p *cell) size() int {
	if p == nil {
		return 0
	}
	var n int
	for n = 1; p.next != nil; n = n + 1 {
		p = p.next
	}
	return n
}
func (p *cell) clone() *cell {
	return &cell{
		type_: p.type_,
		value: p.value,
		next:  nil}
}
func (p *cell) list() *cell {
	return (p.value).(*cell)
}
func (p *cell) isEmptyList() bool {
	return p.isList() && p.list() == nil
}
func (p *cell) isList() bool {
	return p.type_ == list
}

func makeLambda(l func(env *envirs, args *cell) *cell) *cell {
	return &cell{type_: fn, value: l, next: nil}
}
func (p *cell) lambda() func(*envirs, *cell) *cell {
	return (p.value).(func(*envirs, *cell) *cell)
}
func (p *cell) isLambda() bool {
	return p.type_ == fn
}
func makeStr(s *string) *cell {
	return &cell{type_: string_, value: s, next: nil}
}
func (p *cell) str() *string {
	s := (p.value).(*string)
	return s
}
func (p *cell) isStr() bool {
	return p.type_ == string_
}
func makeTrue() *cell {
	return &cell{type_: true_, value: nil, next: nil}
}
func (p *cell) isTrue() bool {
	return p.type_ == true_
}
func makeFalse() *cell {
	return &cell{type_: false_, value: nil, next: nil}
}
func (p *cell) isFalse() bool {
	return p.type_ == false_
}

func makeNumber(i *int) *cell {
	return &cell{type_: number, value: i, next: nil}
}
func (p *cell) number() *int {
	return (p.value).(*int)
}
func (p *cell) isNumber() bool {
	return p.type_ == number
}

func makeKeyword(s *string) *cell {
	return &cell{type_: keyword, value: s, next: nil}
}
func (p *cell) isKeyword() bool {
	return p.type_ == keyword
}

func makeSymbol(s *string) *cell {
	return &cell{type_: symbol, value: s, next: nil}
}
func (p *cell) isSymbol() bool {
	return p.type_ == symbol
}
func (p *cell) symbol() *string {
	return (p.value).(*string)
}

func (p *cell) isScalar() bool {
	return p.type_ != list
}

func (p *cell) show() *string {
	switch true {
	case p.isStr():
		return p.str()
	case p.isDottedPair():
		n1, n2 := p.dottedPair()
		s := fmt.Sprintf("(%s . %s)", *n1.show(), *n2.show())
		return &s
	case p.isTrue():
		return &TRUE
	case p.isFalse():
		return &FALSE
	case p.isNumber():
		s := fmt.Sprintf("%d", *p.number())
		return &s
	case p.isKeyword():
		return p.str()
	case p.isSymbol():
		return p.symbol()
	case p.isLambda():
		return &lambdastr
	case p.isList():

		var s string = "("
		for cell := p.list(); cell != nil; cell = cell.next {
			if s != "(" {
				s = s + " "
			}
			s = s + *cell.show()
		}
		s = s + ")"
		return &s
	default:
		fmt.Printf("Invalid type in show(): %v\n", p)
		os.Exit(1)
	}
	return nil
}
