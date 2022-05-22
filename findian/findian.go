package findian

import (
	"fmt"
	"strings"
)

func AskForString() string {
	var input string
	fmt.Print("Enter a string: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return ""
	}
	return strings.ToLower(input)
}

func SearchRunes() {
	input := AskForString()
	found := strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a")
	if found {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
