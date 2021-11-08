package gr

func Finalise(s string) string {
	toPrint := []rune(s)
	newSlice := ""
	for a := 0; a < len(toPrint); a++ {
		if toPrint[a] == ' ' {
			if a == len(toPrint)-1 {
				continue
			} else if (a+1) != len(toPrint) && toPrint[a+1] == ' ' {
				//a++
			} else {
				newSlice += string(toPrint[a])
			}
		} else {
			newSlice += string(toPrint[a])
		}
	}
	return newSlice
}
