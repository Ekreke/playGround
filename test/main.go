package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	input(slice...)

	username := "devtest"
	fmt.Println(username[:3])
}
func input(ints ...int) {
	for _, v := range ints {
		fmt.Println(v)
	}
}
