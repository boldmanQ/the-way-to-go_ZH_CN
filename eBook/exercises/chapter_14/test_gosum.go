package main

import "fmt"

func sum(x,y int, c chan int) {
    c <- x+y
}

func main() {
    x:=1
    y:=2
    ch := make(chan int)
    go sum(x,y,ch)
    fmt.Println(<-ch)
}
