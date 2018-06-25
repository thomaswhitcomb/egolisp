package main

import "strings"

func tokenize(s string) []string {
	var tokens []string = make([]string, 0)
	var b bool = true
	splits := strings.Split(s, "\"")
	for _, split := range splits {
		if b {
			tokens = append(tokens, tokenize1(split)...)
			b = false
		} else {
			tokens = append(tokens, split)
			b = true
		}
	}
	return tokens
}
func tokenize1(s string) []string {
	replacer := strings.NewReplacer("(", " ( ", ")", " ) ", "'", " ' ")
	s = replacer.Replace(s)
	return strings.Fields(s)
}
