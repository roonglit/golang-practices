package main

import (
	"fmt"
)

// 1.
type Point struct {
	x int
	y int
}

// 2.
func (p *Point) move(x, y int) {
	p.x += x
	p.y += y
}

func (p *Point) display() string {
	return fmt.Sprintf("{x: %d, y: %d}", p.x, p.y)
}

// 3.
type Car struct {
	Color string
	Point // Anonymous type
}

// 4.
// func (c *Car) display() string {
// 	return fmt.Sprintf("{Color: %s}, Location: %s", c.Color, c.Point.display())
// }

// 5. Interface
type Displayable interface {
	display() string
}

func main() {
	// 1. Define new Point struct
	p := Point{x: 5, y: 8}

	// 2. Calling struct method
	p.move(3, 3)
	println(p.display())

	// 3. Anonymous struct
	car := Car{Color: "red", Point: Point{x: 5, y: 3}}
	car.move(3, 4)
	println(car.display())

	// 4. override display function
	// uncomment the display method in Car struct to use it
	println(car.display())

	// 5. Using interface
	var d Displayable = &car
	println(d.display())
}
