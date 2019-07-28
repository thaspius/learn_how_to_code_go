package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting main...")

	test1()

	x := 15

	fmt.Println(x)

	x = 22

	fmt.Println(x)

	fmt.Println("Existing.")
}

func test1() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Line ", i)
		}
	}
}
