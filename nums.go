package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ReplaceNums(text string) string {
	w := strings.Fields(text)
	for i := 0; i < len(w); i++ {
		if w[i] == "(hex)" && i > 0 {
			hexEx, err := strconv.ParseInt(w[i-1], 16, 64)
			if err != nil {
				fmt.Println("Conversion Failed: ", err)
			}
			w[i-1] = strconv.Itoa(int(hexEx))
			w = append(w[:i], w[i+1:]...)
			i--
			continue

		} else if w[i] == "(bin)" && i > 0 {
			binEx, err := strconv.ParseInt(w[i-1], 2, 64)
			if err != nil {
				fmt.Println("Conversion Failed: ", err)
			}
			w[i-1] = strconv.Itoa(int(binEx))
			w = append(w[:i], w[i+1:]...)
			i--
			continue
		}
	}
	return strings.Join(w, " ")
}
