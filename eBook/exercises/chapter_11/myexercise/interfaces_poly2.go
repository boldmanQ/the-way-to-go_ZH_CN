package main

import "math"
import "fmt"

type Shaper interface {
	Area() float32
    Girth() float32
}
type innerInerface struct {
    Shaper
}
type Square struct {
	side float32
}
type Rectangle struct {
	length, width float32
}
type Circle struct {
    banjing float32
}

func (caculater innerInerface) Explain() {
    fmt.Println(caculater.Area())
}
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (sq *Square) Girth() float32 {
	return 2 * sq.side + 2 * sq.side
}
func (cir Circle) Area() float32 {
    return math.Pi * cir.banjing * cir.banjing
}
func (cir Circle) Girth() float32 {
    return  2 * math.Pi * cir.banjing
}
func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {

	//r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	//fmt.Println("Looping through shapes for area ...")
	//	fmt.Println("Shape details: ", shapes[n])
	//	fmt.Println("Area of this shape is: ", shapes[n].Area())
	//	//fmt.Println("Area of this shape is: ", shapes[n].Girth())
	//}
    g := Circle{1}
	shapes := []Shaper{q, g}
	for n, _ := range shapes {
        fmt.Println(shapes[n].Girth())
        //fmt.Println(shapes)
    }
    c := innerInerface{Shaper{}}
    c.Explain()
}
