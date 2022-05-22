package bubble

import "fmt"

func Swap(slice []int, i int) []int {
	slice[i], slice[i+1] = slice[i+1], slice[i]
	return slice
}

func BubbleSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice = Swap(slice, j)
			}
		}
	}
	return slice
}

func AskForInputAndSort() {
	var slice []int
	for {
		var input int
		fmt.Print("Enter a number: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		if input == 0 {
			break
		}
		slice = append(slice, input)
	}
	fmt.Println(BubbleSort(slice))
}
