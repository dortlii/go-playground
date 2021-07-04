package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3}
	array[0] = 5
	fmt.Printf("%v\n", array)

	fmt.Printf("Number of max items: #%v\n", len(array))
}
