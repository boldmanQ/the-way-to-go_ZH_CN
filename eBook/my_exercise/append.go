package main

import "fmt"

func main() {
    var a = []int{1,2,3}
    var factor int = 2
    fmt.Println(Append(a, factor))
}
func Append(x []int, factor int) []int {
    newSlice := make([]int, len(x) * factor)
    n := copy(newSlice, x)
    fmt.Println(n)
    return newSlice
}
