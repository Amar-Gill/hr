package main

import (
	"fmt"
)

func main() {
	var s string
	var k int32

	fmt.Print("String to encode: ")
	fmt.Scanf("%s", &s)
	fmt.Print("Number of letters to rotate: ")
	fmt.Scanf("%d", &k)

	fmt.Println(caeser(s, k))
}

func caeser(s string, k int32) string {

	const minCap = 'A'
	const maxCap = 'Z'

	const minLowerCase = 'a'
	const maxLowerCase = 'z'

	ret := ""

	for _, r := range s {
		isCapitalLetter := (r >= minCap && r <= maxCap)
		isLowerCaseLetter := (r >= minLowerCase && r <= maxLowerCase)

		if !isCapitalLetter && !isLowerCaseLetter {
			ret = ret + string(r)
			continue
		}

		incrementor := k % 26

		newRune := r + incrementor

		if isCapitalLetter && newRune > maxCap {
			newRune = (newRune - maxCap) + minCap - 1
		}

		if isLowerCaseLetter && newRune > maxLowerCase {
			newRune = (newRune - maxLowerCase) + minLowerCase - 1
		}

		ret = ret + string(newRune)
	}
	return ret
}
