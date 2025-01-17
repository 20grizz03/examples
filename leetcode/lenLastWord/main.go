package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" && words[i] != " " {
			return len(words[i])
		}
	}
	return 0

}

func main() {
	len := lengthOfLastWord(" fly me to the moon ")
	fmt.Print(len)
}
