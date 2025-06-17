package main

import (
	"fmt"
)

func main() {
	// 1. Basic Hello World
	fmt.Println("Hello World")

	// 2. Variable declaration
	var a int = 5
	b := 10
	fmt.Printf("My num is: %d\n", a)
	fmt.Printf("My second num is: %d\n", b)

	// 3. Constant declaration
	const c int = 15
	fmt.Printf("My constant num is: %d\n", c)

	// 4. Data types
	var d float64 = 3.14
	var e string = "Hello"
	var f bool = true
	fmt.Printf("My float num is: %f\n", d)
	fmt.Printf("My string is: %s\n", e)
	fmt.Printf("My boolean is: %t\n", f)
	y := "Hello"
	var g string = y + " world"
	println(g)

	// 4. Array and slice
	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{6, 7, 8, 9, 10}
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
	slice = append(slice, 11) // Append to slice
	fmt.Println("Updated Slice:", slice)
	// arr = append(arr, 6)
	// iterate through array
	for i, v := range arr {
		fmt.Printf("Array element at index %d: %d\n", i, v)
	}

	// 5. Map
	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("Map:", myMap)
	// Iterate through map
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
