package main

import (
	"fmt"

	inner "golang101/an_inner_module"
)

func main() {
	// 1. Basic calling function
	fmt.Printf("My num is: %d\n", add3(3, 4, 5))

	// 2. Calling func from another file, using "go run *.go"
	x := add4(5, 6, 7, 8)
	fmt.Printf("My second num is %d\n", x)

	// 3. Calling func from a module
	// Need to initiaalize "go mod init"
	// Then create an_inner_module module inside
	// Then refer to that module here
	inner.Show2Nums(x, 4)

	// 4. Multiple results from a function
	a, b := swap("Hello", "World")
	fmt.Printf("Swapped values: %s, %s\n", a, b)

	// 5. Named return values
	fmt.Println(split(17))
}

func add3(a, b, c int) int {
	return a + b + c
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
