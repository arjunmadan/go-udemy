package main

import "fmt"

func main() {
	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index := range intSlice {
		if intSlice[index]%2 == 0 {
			fmt.Println(intSlice[index], " is even")
		} else {
			fmt.Println(intSlice[index], " is odd")
		}
	}
}
