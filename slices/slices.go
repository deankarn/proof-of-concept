package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}

	fmt.Println(s)

	alterSlice(s)
	fmt.Println(s)
}

func alterSlice(s []int) {

	s[0] = 4
	s[1] = 3
	s[2] = 2
	s[3] = 1

	s = append(s, 5)

	fmt.Println(s)
}
