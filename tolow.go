package gr

import (
	"strings"
)

func ToLow(s string) string {
	sliceOfS := strings.Split(s, " ")
	var newSlice string

	var pos int
	for a := len(sliceOfS) - 1; a >= 0; a-- {
		char1 := []rune(sliceOfS[a])
		if len(char1) > 1 {
			if (char1[0] >= '0' && char1[1] == ')') ||
				(char1[0] <= '9' && char1[1] == ')') {
				pos = int(char1[0]) - 48
				if sliceOfS[a-1] != "(low," {
					newSlice = sliceOfS[a] + " " + newSlice
				}
			}
		}
		if sliceOfS[a] != "(low)" &&
			sliceOfS[a] != "(low," {
			if len(char1) > 1 {
				if (char1[0] >= '0' && char1[0] <= '9') && char1[1] == ')' {
					continue
				} else {
					newSlice = sliceOfS[a] + " " + newSlice
				}
			} else {
				newSlice = sliceOfS[a] + " " + newSlice
			}
		} else if sliceOfS[a] == "(low)" {
			a--
			var stringToAdd []rune
			for _, char := range sliceOfS[a] {

				stringToAdd = append(stringToAdd, char+32)

			}
			newSlice = string(stringToAdd) + " " + newSlice
		} else if sliceOfS[a] == "(low," {
			if pos <= a {
				for i := pos; i > 0; i-- {
					pos--
					var stringToAdd []rune
					for _, char := range sliceOfS[a-1] {

						stringToAdd = append(stringToAdd, char+32)

					}
					newSlice = string(stringToAdd) + " " + newSlice
					a--
				}
			} else {
				return "A part of the test file is not valid."
			}
		}
	}
	return newSlice
}
