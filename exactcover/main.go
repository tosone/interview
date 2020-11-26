package main

import (
	"fmt"
)

func tester(input [9]int) bool {
	return false
}

func main() {
	var input = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(tester(input)) // expect: true
}
