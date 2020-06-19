// https://www.codewars.com/kata/5508249a98b3234f420000fb/train/go

package main

import (
	"fmt"
	"math"
	"strings"
)

var smallLetters = []rune("abcdefghijklmnopqrstuvwxyz")
var capLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func MovingShift(s string, shift int) []string {
	var runes = []rune(s)

	for idx, chr := range runes {
		if chr >= 'a' && chr <= 'z' {
			// Find the location of the letter in the array
			arrLoc := int(chr - 'a')

			// Increase the location position
			arrLoc += shift + idx

			// New location cannot exceed the array size
			arrLoc %= len(smallLetters)

			// Replace with the new character
			runes[idx] = smallLetters[arrLoc]
		} else if chr >= 'A' && chr <= 'Z' {
			// Find the location of the letter in the array
			arrLoc := int(chr - 'A')

			// Increase the location position
			arrLoc += shift + idx

			// New location cannot exceed the array size
			arrLoc %= len(capLetters)

			// Replace with the new character
			runes[idx] = capLetters[arrLoc]
		} else {
			// If not an alphabet, just leave it as it is
			runes[idx] = chr
		}
	}

	// Now that encryption is completed, let's divide it to 5 people in a best effort basis
	// Count the maximum number of people who can carry the secret
	const totalRunners = 5
	var maxRunners int
	for i := totalRunners; i > 0; i-- {
		if len(s)/i > 0 {
			maxRunners = i
			break
		}
	}

	// Now, let's calculate the maximum possible length of each of the divided message
	var maxPossibleLength = int(math.Ceil(float64(len(s)) / float64(maxRunners)))

	// Let's divide the cipher text to each runner
	var separatedTexts = []string{"", "", "", "", ""}
	for i := 0; i < totalRunners; i++ {
		start := maxPossibleLength * i

		// Boundary check. If the start location exceeds the length of the string, just exit the loop.
		// There is no point in keep doing the loop as there is nothing more to divide.
		if start >= len(s) {
			break
		}

		end := start + maxPossibleLength

		// Boundary check
		if end > len(s) {
			end = len(s)
		}

		separatedTexts[i] = string(runes[start:end])
	}

	return separatedTexts
}

func DemovingShift(arr []string, shift int) string {
	// Join all the strings and convert to array of runes
	var runes = []rune(strings.Join(arr, ""))

	for idx, chr := range runes {
		if chr >= 'a' && chr <= 'z' {
			// Find the location of the letter in the array
			arrLoc := int(chr - 'a')

			// Decrease the location position
			arrLoc -= shift + idx
			arrLoc += ((idx / len(smallLetters)) + 2) * len(smallLetters)

			// New location cannot exceed the array size
			arrLoc %= len(smallLetters)

			// Replace with the new character
			runes[idx] = smallLetters[arrLoc]
		} else if chr >= 'A' && chr <= 'Z' {
			// Find the location of the letter in the array
			arrLoc := int(chr - 'A')

			// Decrease the location position
			arrLoc -= shift + idx
			arrLoc += ((idx / len(capLetters)) + 2) * len(capLetters)

			// New location cannot exceed the array size
			arrLoc %= len(capLetters)

			// Replace with the new character
			runes[idx] = capLetters[arrLoc]
		} else {
			// If not an alphabet, just leave it as it is
			runes[idx] = chr
		}
	}

	return string(runes)
}

func main() {
	var cipherTexts = MovingShift("I should have known that you would have a perfect answer for me!!!", 1)
	fmt.Println(cipherTexts)

	var plainText = DemovingShift(cipherTexts, 1)
	fmt.Println(plainText)
}
