package main

import "fmt"

type test func() interface{}

func main() {

	var fn test
	var err error

	// fn := func() interface{} {
	// 	return nil
	// }

	fmt.Println("fn:", fn == nil)

	fmt.Println("err:", err == nil)

	// Also Test constant bools vs inline
	// func(true)
	// func(construe)
}
