package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return
	}
	text := process(string(data))
	text += "\n"
	os.WriteFile(outputFile, []byte(text), 0644)
}
func process(text string) string {
	word := strings.Split(text,"\n")
	for i := range word {
		 word[i] = ApplyCasing(word[i])
		 word[i] = Quote(word[i])
		 word[i] = Punctuation(word[i])
		 word[i] = FixAtoAn(word[i])
		 word[i] = ReplaceNums(word[i])
	}
	return strings.Join(word,"\n")
}
