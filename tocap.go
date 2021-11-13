package gr

import (
	"strings"
)

func ToCap(s string) string {
	sliceOfS := strings.Split(s, " ")
	var newSlice string
	var pos int

	for a := len(sliceOfS) - 1; a >= 0; a-- {
		char1 := []rune(sliceOfS[a])
		if len(char1) > 1 {
			if (char1[0] >= '0' && char1[1] == ')') ||
				(char1[0] <= '9' && char1[1] == ')') {
				pos = int(char1[0]) - 48
				if sliceOfS[a-1] != "(cap," {
					newSlice = sliceOfS[a] + " " + newSlice
				}
			}
		}
		if sliceOfS[a] != "(cap)" &&
			sliceOfS[a] != "(cap," {
			if len(char1) > 1 {
				if (char1[0] >= '0' && char1[0] <= '9') && char1[1] == ')' {
					continue
				} else {
					newSlice = sliceOfS[a] + " " + newSlice
				}
			} else {
				newSlice = sliceOfS[a] + " " + newSlice
			}
		} else if sliceOfS[a] == "(cap)" {
			a--
			var stringToAdd []rune
			for index, char := range sliceOfS[a] {

				if index == 0 {
					stringToAdd = append(stringToAdd, char-32)
				} else {
					stringToAdd = append(stringToAdd, char)
				}
			}
			newSlice = string(stringToAdd) + " " + newSlice
		} else if sliceOfS[a] == "(cap," {
			// if pos < a {
			for i := pos; i > 0; i-- {
				pos--
				var stringToAdd []rune
				for index, char := range sliceOfS[a-1] {
					if index == 0 {
						stringToAdd = append(stringToAdd, char-32)
					} else {
						stringToAdd = append(stringToAdd, char)
					}
				}
				newSlice = string(stringToAdd) + " " + newSlice
				a--
			}
			//}
		}
	}
	return newSlice
}
