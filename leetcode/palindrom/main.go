package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	iToS := strconv.Itoa(x)

	for i := 0; i < len(iToS)/2; i++ {
		if iToS[i] != iToS[len(iToS)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print(isPalindrome(12222))
}
