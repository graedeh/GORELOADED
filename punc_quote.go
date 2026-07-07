package main

import (
	"strings"
	"regexp"
)
func Punctuation(word string) string {
	grace := regexp.MustCompile(`\s([,.?;:'!]+)`)
	word = grace.ReplaceAllString(word,"$1")

	grace2 := regexp.MustCompile(`([,.?;:'!]+)([\s+,.?;:'!])`)
	word = grace2.ReplaceAllString(word,"$1 $2")
	return  word
}
func Quote(text string) string {
	part := strings.Split(text,"'")
	for i := 0; i < len(part); i++ {
		if i % 2 == i {
			part[i] = strings.Join(part, " ")
		}
	}
	book := strings.Join(part,"'")
	cat := strings.Split(book,`"`)
	for i := 0; i < len(cat);i++ {
		if i % 2 == 1 {
			cat[i] = strings.TrimSpace(cat[i])
		}
	}
	return  strings.TrimSpace(strings.Join(strings.Fields(strings.Join(cat,`"`))," "))
}