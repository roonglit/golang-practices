package main

import (
	"fmt"
	inner "golang101/an_inner_module"
)

// 6.
type Point struct {
	x int
	y int
}

func (p *Point) move(x, y int) {
	p.x += x
	p.y += y
}

func (p *Point) display() string {
	return fmt.Sprintf("{x: %d, y: %d}", p.x, p.y)
}

// 7.
type Car struct {
	Color string
	Point // Anonymous type
}

// 8.
func (c *Car) display() string {
	return fmt.Sprintf("{Color: %s}, Location: %s", c.Color, c.Point.display())
}

func main() {
	// 1. Basic Hello World
	fmt.Println("Hello World")

	// 2. Basic calling function
	fmt.Printf("My num is: %d\n", add3(3, 4, 5))

	// 3. Calling func from another file, using "go run *.go"
	x := add4(5, 6, 7, 8)
	fmt.Printf("My second num is %d\n", x)

	// 4. Calling func from a module
	// Need to initiaalize "go mod init"
	// Then create an_inner_module module inside
	// Then refer to that module here
	inner.Show2Nums(x, 4)

	// 5. Using := is a shorthand to assign type automatically
	y := "Hello"
	var g string = y + " world"
	println(g)

	// 6. Define new Point struct
	p := Point{x: 5, y: 8}
	p.move(3, 3)
	println(p.display())

	// 7. Define Car struct
	car := Car{Color: "red", Point: Point{x: 5, y: 3}}
	car.move(3, 4)
	println(car.display())

	// 8. override display function
	println(car.display())
}

func add3(a, b, c int) int {
	return a + b + c
}
