package main

import "fmt"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, x := range ints {
		if x%2 == 0 {
			fmt.Println(x, " is even")
		} else {
			fmt.Println(x, " is odd")
		}
	}
}
