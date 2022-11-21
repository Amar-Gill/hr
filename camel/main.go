package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1 // +1 because first word is not capitalized

	// a string is actually a []byte
	for _, r := range input {
		s := string(r)
		if strings.ToUpper(s) == s {
			answer++
		}
	}

	fmt.Println(answer)
}
