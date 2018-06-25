package main

import (
	"fmt"
	"os"
)

type funintf func(envir *envirs, args *cell) *cell

var specialFuncs = map[string]string{
	"quote":  "dynamic",
	"defun":  "static",
	"if":     "static",
	"lambda": "static",
	"define": "static",
}

type fundef struct {
	ep       funintf
	arityMin int
	arityMax int
}

var funcs map[string]*fundef
var vars map[string]*cell

func init() {
	vars = map[string]*cell{}
	funcs = map[string]*fundef{
		"quote":  &fundef{quote, 1, 1},
		"list":   &fundef{list_, 0, 10},
		"+":      &fundef{add, 2, 10},
		"-":      &fundef{subtract, 2, 2},
		"*":      &fundef{multiply, 2, 10},
		"quit":   &fundef{quit, 0, 0},
		"car":    &fundef{car, 1, 1},
		"cdr":    &fundef{cdr, 1, 1},
		"if":     &fundef{ifthenelse, 3, 3},
		"eq?":    &fundef{isequal, 2, 2},
		"lambda": &fundef{lambda_, 2, 2},
		"defun":  &fundef{defun, 3, 3},
		"atom?":  &fundef{isatom, 1, 1},
		"cons":   &fundef{cons, 2, 2},
		"null?":  &fundef{isnull, 1, 1},
		"define": &fundef{define, 2, 2},
		"not":    &fundef{not, 1, 1},
		"pair?":  &fundef{isPair, 1, 1},
	}
}

func define(env *envirs, cells *cell) *cell {
	if !cells.isSymbol() && !cells.isKeyword() {
		fmt.Printf("var '%s' must be symbol or keyword\n", *cells.show())
		return makeFalse()
	}
	vars[*cells.show()] = eval(env, cells.next)
	return makeTrue()
}
func isnull(env *envirs, cells *cell) *cell {
	if !cells.isList() {
		fmt.Printf("null? must have a list s-expressions as the first parameter. received %s\n", *cells.show())
		return makeFalse()
	}
	if cells.list().size() == 0 {
		return makeTrue()
	}
	return makeFalse()

}
func cons(env *envirs, cells *cell) *cell {
	if !cells.next.isList() {
		n := makeDottedPair(cells, cells.next)
		return n
	}
	n := cells.clone()
	n.next = cells.next.list()
	return makeList(n)
}
func isatom(env *envirs, cells *cell) *cell {
	if cells.isList() {
		return makeFalse()
	} else {
		return makeTrue()
	}
}
func isequal(env *envirs, cells *cell) *cell {
	if *cells.show() == *cells.next.show() {
		return makeTrue()
	} else {
		return makeFalse()
	}
}

func ifthenelse(env *envirs, cells *cell) *cell {
	cond := eval(env, cells)
	if cond.isFalse() {
		return eval(env, cells.next.next)
	} else {
		return eval(env, cells.next)
	}
}
func car(env *envirs, cells *cell) *cell {
	if cells.isList() {
		if cells.list().size() == 0 {
			return makeFalse()
		} else {
			return cells.list().clone()
		}
	} else {
		fmt.Printf("car requires a list\n")
		os.Exit(1)
	}
	return makeFalse()
}
func cdr(env *envirs, cells *cell) *cell {
	if cells.isList() {
		x := cells.list().next
		return makeList(x)
	} else {
		fmt.Printf("cdr requires a list. received %s\n", *cells.show())
		os.Exit(1)
	}
	return makeFalse()
}
func quit(env *envirs, cells *cell) *cell {
	os.Exit(0)
	return nil
}
func subtract(env *envirs, cells *cell) *cell {
	i := *cells.number() - *cells.next.number()
	return makeNumber(&i)
}
func add(env *envirs, cells *cell) *cell {
	var number int
	for n := cells; n != nil; n = n.next {
		number = number + *n.number()
	}
	return makeNumber(&number)
}
func multiply(env *envirs, cells *cell) *cell {
	var number int = 1
	for n := cells; n != nil; n = n.next {
		number = number * *n.number()
		//fmt.Printf("multiplying %d times %d\n", number, *n.number())
	}
	return makeNumber(&number)
}
func list_(env *envirs, cells *cell) *cell {
	return makeList(cells)
}

func quote(env *envirs, cells *cell) *cell {
	return cells
}

func newFunc(params *cell, body *cell) func(*envirs, *cell) *cell {
	fn := func(env *envirs, ps *cell) *cell {
		var body *cell = body
		var params *cell = params
		var env1 *envirs = newEnvirs(env)
		for p := params.list(); p != nil; p = p.next {
			if ps == nil {
				fmt.Printf("Choak - env: %v and body: %v\n", *env1, *body.show())
				os.Exit(2)
			}
			env1.put(*p.symbol(), ps.clone())
			ps = ps.next
		}
		return eval(env1, body)
	}
	return fn
}

func defun(env *envirs, cells *cell) *cell {
	if cells.isSymbol() == false {
		fmt.Printf("Missing or incorrect function name\n")
		return makeFalse()
	}
	if cells.next.isList() == false {
		fmt.Printf("Missing or incorrect formal parameters\n")
		return makeFalse()
	}
	funcs[*cells.symbol()] = &fundef{
		ep:       newFunc(cells.next, cells.next.next),
		arityMin: cells.next.list().size(),
		arityMax: cells.next.list().size(),
	}
	return makeTrue()
}
func lambda_(env *envirs, cells *cell) *cell {
	if cells.isList() == false {
		fmt.Printf("Missing or incorrect formal parameters\n")
		return nil
	}
	fnptr := newFunc(cells, cells.next)

	return makeLambda(fnptr)
}
func not(env *envirs, cells *cell) *cell {
	if cells.isTrue() {
		return makeFalse()
	} else {
		return makeTrue()
	}
}
func isPair(env *envirs, cells *cell) *cell {
	if cells.isDottedPair() {
		return makeTrue()
	} else {
		return makeFalse()
	}
}
