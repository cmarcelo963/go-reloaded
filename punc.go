package gr

func Punc(s string) string {
	runeOfS := []rune(s)
	newSlice := ""
	for a := 0; a < len(runeOfS); a++ {
		if a == (len(runeOfS) - 1) {
			newSlice += string(runeOfS[a])
		} else if a == 0 {
			newSlice += string(runeOfS[a])
			if runeOfS[a+1] == ' ' {
				newSlice += " "
			}
		} else if runeOfS[a] >= 'a' && runeOfS[a] <= 'z' ||
			runeOfS[a] >= 'A' && runeOfS[a] <= 'Z' ||
			runeOfS[a] == '"' ||
			runeOfS[a] >= '0' ||
			runeOfS[a] <= '9' {
			if runeOfS[a-1] == '.' ||
				runeOfS[a-1] == ',' ||
				runeOfS[a-1] == ':' ||
				runeOfS[a-1] == ';' {
				newSlice += " "
			}
			newSlice += string(runeOfS[a])
			if runeOfS[a+1] == ' ' && (a+2) != len(runeOfS) && (runeOfS[a+2] == '.' ||
				runeOfS[a+2] == ',' ||
				runeOfS[a+2] == '!' ||
				runeOfS[a+2] == '?' ||
				runeOfS[a+2] == ':' ||
				runeOfS[a+2] == ';' ||
				runeOfS[a+2] == '\'') {
				a++

			}
		} else if runeOfS[a] == '\'' {
			if runeOfS[a-1] >= 'a' && runeOfS[a-1] <= 'z' ||
				runeOfS[a-1] >= 'A' && runeOfS[a-1] <= 'Z' {
				newSlice += string(runeOfS[a])
			} else if runeOfS[a+1] == ' ' {
				newSlice += string(runeOfS[a])
				a++
			}
		} else if runeOfS[a] == ',' ||
			runeOfS[a] == '?' ||
			runeOfS[a] == ':' ||
			runeOfS[a] == ';' {
			newSlice += string(runeOfS[a])
			if runeOfS[a+1] != ' ' {
				newSlice += " "
			}
		} else if runeOfS[a] == '!' {
			if runeOfS[a+1] == '?' {
				newSlice += "!?"
				a++
			}
		} else if runeOfS[a] == '.' {
			newSlice += "."
			if runeOfS[a+1] == ' ' && runeOfS[a+2] == '.' {
				a++
			}
		}
	}
	return newSlice
}
