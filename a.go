package gr

func A(s string) string {
	char := []rune(s)
	toReturn := ""

	for a := 0; a < len(s); a++ {
		if char[a] == '.' || char[a] == '?' || char[a] == '!' {
			toReturn += string(char[a])
		} else if char[a] != 'a' && char[a] != 'A' {
			toReturn += string(char[a])
		} else if char[a] == 'a' {
			if char[a+1] == ' ' {
				if char[a+2] == 'a' ||
					char[a+2] == 'A' ||
					char[a+2] == 'e' ||
					char[a+2] == 'E' ||
					char[a+2] == 'i' ||
					char[a+2] == 'I' ||
					char[a+2] == 'o' ||
					char[a+2] == 'O' ||
					char[a+2] == 'u' ||
					char[a+2] == 'U' ||
					char[a+2] == 'h' ||
					char[a+2] == 'H' {
					toReturn += "an"
				}
			} else {
				toReturn += string(char[a])
			}
		} else if char[a] == 'A' {
			if char[a+1] == ' ' {
				if char[a+2] == 'a' ||
					char[a+2] == 'A' ||
					char[a+2] == 'e' ||
					char[a+2] == 'E' ||
					char[a+2] == 'i' ||
					char[a+2] == 'I' ||
					char[a+2] == 'o' ||
					char[a+2] == 'O' ||
					char[a+2] == 'u' ||
					char[a+2] == 'U' ||
					char[a+2] == 'h' ||
					char[a+2] == 'H' {
					toReturn += "An"
				}
			} else {
				toReturn += string(char[a])
			}
		}
	}
	return toReturn
}
