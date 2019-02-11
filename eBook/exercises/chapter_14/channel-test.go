package main

import (
	"fmt"
    //"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func sender(ch chan int) {
    ch <- 2
}

func main() {
	out := make(chan int, 1)
    go sender(out)
	go f1(out)
    fmt.Println(<-out)
    out <- 3
    fmt.Println(<-out)
}

