package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var intregexp = regexp.MustCompile("^[0-9]+$")

type parser struct {
	parenStack int
	tokens     []string
	tokenptr   int
}

func newParser() *parser {
	return &parser{
		parenStack: 0,
		tokenptr:   0,
	}
}
func (p *parser) setTokens(t []string) {
	p.tokens = t
}
func (p *parser) balanced() bool {
	var parenStack int = 0
	var ok bool = true
	for _, token := range p.tokens {
		if token == "(" {
			parenStack = parenStack + 1
		}
		if token == ")" {
			parenStack = parenStack - 1
			if parenStack < 0 {
				fmt.Printf("Extra right paren found\n")
				ok = false
			}
		}
	}
	if parenStack > 0 && ok {
		fmt.Printf("Unbalanced parens\n")
		ok = false
	}
	return ok
}
func (p *parser) parse() *cell {
	var n *cell
	var haveQuote bool = false
	var list *cell
	var head, curr *cell
	if !p.balanced() {
		return head
	}
	for p.tokenptr < len(p.tokens) {
		c := p.tokens[p.tokenptr]
		p.tokenptr = p.tokenptr + 1
		switch c {
		case "(":
			list = p.parse()
			n = makeList(list)
		case ")":
			return head
		case "nil":
			n = makeList(makeQuote(makeList(nil)))
		case "#t":
			n = makeTrue()
		case "#f":
			n = makeFalse()
		case "'":
			haveQuote = true
			continue
		default:
			if intregexp.MatchString(c) {
				i, err := strconv.Atoi(c)
				if err != nil {
					fmt.Printf("Invalid number found: %s\n", c)
					n = makeStr(&c)
				} else {
					n = makeNumber(&i)
				}
			} else {
				if c[0] == ':' {
					n = makeKeyword(&c)
				} else {
					n = makeSymbol(&c)
				}
			}
		}
		if haveQuote {
			n = makeList(makeQuote(n))
			haveQuote = false
		}
		if head == nil {
			head = n
		} else {
			curr.next = n
		}
		curr = n
	}
	return head
}

var quoteSymbol string = "quote"

func makeQuote(s *cell) *cell {
	cells := makeSymbol(&quoteSymbol)
	cells.next = s
	return cells
}
