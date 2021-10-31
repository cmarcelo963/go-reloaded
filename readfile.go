package gr

import (
	"fmt"
	"os"
)

func ReadFile() string {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Invalid file. Please try again.")
	}
	fmt.Println("Formatting following text:")
	fmt.Println(string(file))
	return string(file)
}
