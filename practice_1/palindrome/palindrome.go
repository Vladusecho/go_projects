package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var allowed = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func palindrome(n string) bool {
	total := ""
	n = strings.ToLower(n)
	for _, i := range n {
		if unicode.IsLetter(i) {
			total += string(i)
			continue
		}
	}
	return total == reverse(total)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := scanner.Text()
	fmt.Println(palindrome(n))
}
