package main

import (
	"strconv"
	"strings"
)

func ApplyCasing(text string) string {
	words := strings.Fields(text)
	result := []string{}

	for i := 0; i < len(words); i++ {
		if words[i][0] != '(' {
			result = append(result, words[i])
			continue
		}

		cmd := words[i]

		for i+1 < len(words) && cmd[len(cmd)-1] != ')' {
			i++
			cmd += " " + words[i]
		}

		cmd = cmd[1 : len(cmd)-1]
		count := 1

		if strings.Contains(cmd, ",") {
			parts := strings.Split(cmd, ",")
			cmd = strings.TrimSpace(parts[0])
			n, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			if n > 0 {
				count = n
			}
		}

		for j := 0; j < count && j < len(result); j++ {
			idx := len(result) - 1 - j

			switch cmd {
			case "up":
				result[idx] = strings.ToUpper(result[idx])
			case "low":
				result[idx] = strings.ToLower(result[idx])
			case "cap":
				w := strings.ToLower(result[idx])
				if len(w) > 0 {
					result[idx] = strings.ToUpper(string(w[0])) + w[1:]
				}
			}
		}
	}

	return strings.Join(result, " ")
}
