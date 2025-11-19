package main
import ("math"
		"fmt")

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rect := Rectangle{height:5, width:4}
	circle := Circle{radius:2}

	fmt.Println("Rectangle area:", calculateArea(rect))
	fmt.Println("Circle area:", calculateArea(circle))
	
}