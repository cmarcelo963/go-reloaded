package gr

import (
	"strconv"
	"strings"
)

func ToHex(s string) string {
	sliceOfS := strings.Split(s, " ")
	var newSlice string
	for a := len(sliceOfS) - 1; a >= 0; a-- {
		if sliceOfS[a] != "(hex)" {
			newSlice = sliceOfS[a] + " " + newSlice
		} else if sliceOfS[a] == "(hex)" {
			a -= 1
			var toAdd string
			numberString := strings.Replace(sliceOfS[a], toAdd, "", -1)
			numberString = strings.Replace(numberString, toAdd, "", -1)
			output, _ := strconv.ParseInt(numberString, 16, 64)
			intToS := strconv.FormatInt(output, 10)
			newSlice = intToS + " " + newSlice
		}
	}
	return newSlice
}