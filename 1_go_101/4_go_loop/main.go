package main

import "fmt"

func main() {
	// 1. Basic for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 2. For loop with condition
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// 3. For loop with range
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// 4. Nested for loop
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

	// 5. Infinite loop with break
	count := 0
	for {
		count++
		if count > 5 {
			break
		}
		fmt.Println("Count:", count)
	}

	// 6. Using continue to skip an iteration
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println("Odd number:", i)
	}

	// 7. If statement
	if sum > 10 {
		fmt.Println("Sum is greater than 10")
	} else {
		fmt.Println("Sum is 10 or less")
	}

	// 8 If with shorthand declaration
	if sum3 := sum + 5; sum3 > 10 {
		fmt.Println("Sum3 is greater than 10:", sum3)
	} else {
		fmt.Println("Sum3 is 10 or less:", sum3)
	}

	// 9. Switch statement
	switch sum % 2 {
	case 0:
		fmt.Println("Sum is even")
	case 1:
		fmt.Println("Sum is odd")
	default:
		fmt.Println("Unexpected value")
	}

	// 10. Switch with fallthrough
	switch {
	case sum < 10:
		fmt.Println("Sum is less than 10")
	case sum < 20:
		fmt.Println("Sum is less than 20")
		fallthrough // Continue to next case
	case sum < 30:
		fmt.Println("Sum is less than 30")
	default:
		fmt.Println("Sum is 30 or more")
	}
}
