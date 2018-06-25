package main

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkFactorial(t *testing.B) {
	program :=
		`
		(defun factorial (n) (if (eq? n 1) 1 (* n (factorial (- n 1))))) 
		(factorial 25)
  	`
	result := run(program)
	fmt.Printf("%s\n", result)

}
func TestBasic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"01", "1", "1"},
		{"01.1", "nil", "()"},
		{"02", "abc", "abc"},
		{"03", "(list 1 22 333)", "(1 22 333)"},
		{"04", "(quote 1)", "1"},
		{"05", "(quote (a))", "(a)"},
		{"06", "(car (list (quote (a b c)) d e f))", "(a b c)"},
		{"07", "(+ 3 4 5)", "12"},
		{"08", "(car (list (quote (a b c)) def ghi))", "(a b c)"},
		{"08.1", "(car (quote (((hotdogs)) (and) (pickle) relish)))", "((hotdogs))"},
		{"08.2", "(car (car (quote (((hotdogs)) (and) (pickle) relish))))", "(hotdogs)"},
		{"09", "(list (+ 5 6 1 8) (quote (a b c)) def ghi)", "(20 (a b c) def ghi)"},
		{"10", "(list \"the big dog\" 1 abc)", "(the big dog 1 abc)"},
		{"11", "(car (cdr (list 1 nil 2)))", "()"},
		{"11.1", "(cdr (cdr (list 1 nil 2)))", "(2)"},
		{"11.2", "(cdr (list (a b c) x y z))", "(x y z)"},
		{"12", "(defun plus () +) ((plus) 2 3) ", "5"},
		{"13", "(defun five () 5) (defun square (x) (* x x)) (square (five)) ", "25"},
		{"14", "(defun plus () +) ((lambda (x y) ((plus) x y)) 23 46) ", "69"},
		{"15", "((lambda () 2))", "2"},
		{"16", "((lambda () (quote (1 3))))", "(1 3)"},
		{"16.1", "(define addr (lambda (x y) (+ x y))) (addr 73 7)", "80"},
		{"17", "(list 1 ((lambda () 2)))", "(1 2)"},
		{"18", "(list 1 ((lambda () (+ 2 8))) 25)", "(1 10 25)"},
		{"19", "(list 1 ((lambda (x y) (+ x y)) 4 5) 27)", "(1 9 27)"},
		{"20", "(if nil true false)", "true"},
		{"21", "(if 1 true false)", "true"},
		{"22", "(if '() true false)", "true"},
		{"23", "(if #f true (if 1 wow blah))", "wow"},
		{"23.1", "(if #f 23 (if #f wow blah))", "blah"},
		{"23.2", "(if #f 345 (if #t wow blah))", "wow"},
		{"23.3", "(if #f 213 (if '() wow blah))", "wow"},
		{"24", "(defun factorial (n) (if (eq? n 1) 1 (* n (factorial (- n 1))))) (factorial 8)", "40320"},
		{"24.1", "(defun length (l) (if (eq? l '()) 0 (+ (length (cdr l)) 1))) (length (list 1 2 3))", "3"},
		{"25", "((lambda (y) ((lambda () y)) ) (list 4 8))", "(4 8)"},
		{"26", "((lambda (y) (+ ((lambda (x) (+ x y)) 4) 5)) 7 )", "16"},
		{"26.1", "(define five 5) (defun times5 (x) (* x five)) (times5 7)", "35"},
		{"26.2", "(define x 5) (defun square (x) (* x x)) (square 9)", "81"},
		{"27", "(- 8 1)", "7"},
		{"28", "(* 8 3)", "24"},
		{"29", "(atom? 3)", "#t"},
		{"30", "(atom? abc)", "#t"},
		{"31", "(atom? :joe)", "#t"},
		{"32", "(atom? (quote (1,2)))", "#f"},
		{"33", "'(a b (c d e) f (g))", "(a b (c d e) f (g))"},
		{"34", "(cons peanut '(butter and jelly))", "(peanut butter and jelly)"},
		{"35", "(cons '((help) this) '(is very ((hard) to learn)))", "(((help) this) is very ((hard) to learn))"},
		{"36", "(cons '(a b (c)) '())", "((a b (c)))"},
		{"37", "(cons a (car '((b) c d)))", "(a b)"},
		{"37.1", "(cons c d)", "(c . d)"},
		{"37.2", "(cons (cons a b) (cons c d))", "((a . b) . (c . d))"},
		{"37.2", "(cons 42 (cons 69 (cons 613 nil)))", "(42 69 613)"},
		{"37.3", "(cons 1 nil)", "(1)"},
		{"37.4", "(cons 2 (cons 1 nil))", "(2 1)"},
		{"37.5", "(cons 2 (cons 1 '()))", "(2 1)"},
		{"38", "(null? (quote ()))", "#t"},
		{"38", "(eq? 1 1)", "#t"},
		{"39", "(eq? 1 2)", "#f"},
		{"40", "(eq? '() '())", "#t"},
		{"41", "(eq? '(1) '())", "#f"},
		{"42", "(eq? :abc :def)", "#f"},
		{"43", "(eq? :abc :abc)", "#t"},
		{"44", "(defun map (f l) (if (eq? l '()) '() (cons (f (car l)) (map f (cdr l))))) (define add5 (lambda (x) (+ x 5))) (map add5 (list 1 2 3))", "(6 7 8)"},
		{"45", "(defun length (l) (if (eq? l '()) 0 (+ (length (cdr l)) 1))) (length (list 1 2 3))", "3"},
		{"46", "(defun length (l) (if (eq? l '()) 0 (+ (length (cdr l)) 1))) (length nil)", "0"},
		{"47", "(pair? (cons a b))", "#t"},
		{"47.1", "(pair? nil)", "#f"},
		{"47.2", "(pair? 876)", "#f"},
		{"48", "(* 3 6)", "18"},
		{"48.1", "(* 3 6 2 2)", "72"},
		{"49.0", "(not #t)", "#f"},
		{"49.1", "(not #f)", "#t"},
		{"49.2", "(not nil)", "#t"},
		{"49.3", "(not 23)", "#t"},
		{"50.0", "(defun fib (n) (if (eq? n 0) 1 (if (eq? n 1) 1 (+ (fib (- n 1)) (fib (- n 2)))))) (fib 13) ", "377"},
		{"50.1", "(defun fib (n) (if (eq? n 0) 1 (if (eq? n 1) 1 (+ (fib (- n 1)) (fib (- n 2)))))) (fib 20) ", "10946"},
		{"50.2", "(defun fib (n) (if (eq? n 0) 1 (if (eq? n 1) 1 (+ (fib (- n 1)) (fib (- n 2)))))) (fib 0) ", "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := run(tt.input)
			if ans[len(ans)-1] != tt.expected {
				t.Errorf("Test %s failed. Expected %s and got %s", tt.name, tt.expected, ans)
			}
		})
	}
}
func TestCLI(t *testing.T) {
	input := `
     (+ 
		 2   3 
		 7 )
  `
	inputReader := bytes.NewReader([]byte(input))
	cli(inputReader)
}
