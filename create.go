package gr

import (
	"fmt"
	"os"
)

func CreateFile(s string) {
	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(s)
		fmt.Println("Done: " + s)
	}
	file.Close()
}
