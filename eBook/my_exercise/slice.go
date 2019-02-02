package main

import "fmt"

var x [5]int
//var b = []int{2,3,5,7,11}
var b = x[:]
var c = b[1:3]
func main() {
    x[0] = 2
    x[1] = 3
    x[2] = 5
    x[3] = 7
    x[4] = 11
    fmt.Printf("%v %T",x, x)
    fmt.Println("\n")
    fmt.Printf("%v %T",b, b)
    fmt.Println("\n")
    fmt.Printf("%v %T",c,c)
    fmt.Println("\n")
    fmt.Println(len(c), cap(c))
    x[2] = 11
    fmt.Println("\n")
    fmt.Printf("%v %T",c,c)
}
