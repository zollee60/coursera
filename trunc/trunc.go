package trunc

import "fmt"

func AskForFloatAndTruncate() {
	var input float64
	fmt.Print("Enter a float: ")
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		return
	}
	fmt.Println("Truncated to int:", int(input))
}
