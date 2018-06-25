[![Maintainability](https://api.codeclimate.com/v1/badges/20726fae0d4326316bbc/maintainability)](https://codeclimate.com/github/thomaswhitcomb/mini-scheme/maintainability)

# mini-scheme

A simple,embed-able lisp interpreter implemented in Golang.

## Why

I am learning Golang and love writing mini-languages so I thought this would be fun.

## Status

Very early implementation.  Solid data structure for s-expressions with core support for named functions and anonymous functions.  Almost no error checking, get it right or it might core.

## How to run it

1. Clone the repo
2. cd egolisp
3. make
4. ./bin/darwin/egolisp [optional file name]

The optional file name contains a set of s-expressions that get evaluated before the REPL starts.  It is an easy way to load your predefined functions.

## Current Capabilities

### Data types
  - Positive integer
  - Symbols 
  - Keywords
  - Boolean (T, nil)
  - List 

### Built-in functions
  - Math functions

		(+ integer integer integer...)  
		(* integer integer integer...)  
		(- integer integer)
		
  - Boolean functions 

		(= number s-expression numbers-expression)  
		(eq? non-number s-expression non-number s-expression)  
		(atom s-expression)  
		
  - List functions

		(list s-expression...)  
		(cons s-expression list s-expression)  
		(cdr list)  
		(car list)  

  - Conditional flow functions
 
  		(if s-expression s-expression s-expression)

  - Miscellaneous  
		
		(quote s-expression)
		(atom? s-expression)
		
  - Meta functions  
		
		(defun symbol (formal-parameters) s-expression) 
		(lambda (formal-parameters) s-expression)

  - REPL functions

		(quit)
		

## Working Examples

```lisp 
(cons a b) -> (a . b)
(cons 1 nil) -> (1)
(quote (a)) -> (a)
(list 1 2 3) -> (1 2 3)
(car (list (quote (a b c)) d e f) -> (a b c)
(+ 33 14 50) -> 97
(defun plus () +) ((plus) 2 3) -> 5
(defun plus () +) ((lambda (x y) ((plus) x y)) 23 46) -> 69
(list 1 ((lambda (x y) (+ x y)) 4 5) 27) -> (1 9 27)
(defun five () 5) (defun square (x) (* x x)) (square (five)) -> 25
(if 1 true false) -> true
(if nil true false) -> false
(defun factorial (n) (if (eq? n 1) 1 (* n (factorial (- n 1))))) 
  (factorial 8) -> 40320
((lambda (y) (+ ((lambda (x) (+ x y)) 4) 5)) 7 ) -> 16
(defun map (f l) (if (eq? l '()) '() (cons (f (car l)) (map f (cdr l))))) 
  (define add5 (lambda (x) (+ x 5))) 
  (map add5 (list 1 2 3)) -> (6 7 8)
```



