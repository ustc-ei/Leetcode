package main

import (
	"fmt"
	"leetcode/test"
)

func main() {
	a := "abc"
	fmt.Printf("%T", a[0])
	fmt.Println(test.Add(1, 2))
	fmt.Println(test.Sub(2, 3))
	fmt.Println(test.Mul(3, 4))
	fmt.Println(test.Div(2, 4))
}
