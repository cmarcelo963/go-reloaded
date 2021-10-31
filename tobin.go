package gr

import (
	"strconv"
	"strings"
)

func ToBin(s string) string {
	sliceOfS := strings.Split(s, " ")
	var newSlice string
	for a := len(sliceOfS) - 1; a >= 0; a-- {
		if sliceOfS[a] != "(bin)" {
			newSlice = sliceOfS[a] + " " + newSlice
		} else if sliceOfS[a] == "(bin)" {
			a -= 1
			var toAdd string
			numberString := strings.Replace(sliceOfS[a], toAdd, "", -1)
			numberString = strings.Replace(numberString, toAdd, "", -1)
			output, _ := strconv.ParseInt(numberString, 2, 64)
			intToS := strconv.FormatInt(output, 10)
			newSlice = intToS + " " + newSlice
		}
	}
	return newSlice
}
