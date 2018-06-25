package main

import "testing"

func TestFind1(t *testing.T) {
	e := newEnvirs(nil)
	e.put("a", makeTrue())
	s := e.find("a")
	if *s.show() != *makeTrue().show() {
		t.Errorf("TestFind1 - Expected %s, got %s", *makeTrue().show(), *s.show())
	}
}
func TestFind2(t *testing.T) {
	e := newEnvirs(nil)
	e.put("a", makeTrue())
	e.put("b", makeList(nil))
	e = newEnvirs(e)
	e.put("a", makeList(nil))
	s := e.find("a")
	if *s.show() != *makeList(nil).show() {
		t.Errorf("TestFind2.1 - Expected %s, got %s", *makeList(nil).show(), *s.show())
	}
	s = e.find("b")
	if *s.show() != *makeList(nil).show() {
		t.Errorf("TestFind2.2 - Expected %s, got %s", *makeList(nil).show(), *s.show())
	}
}
