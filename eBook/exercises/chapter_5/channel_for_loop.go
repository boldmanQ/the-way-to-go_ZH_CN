package main

import "fmt"

func main() {
    ch1 := make(chan int)
    go tel(ch1)
    for {
        value, open := <-ch1
        if !open {
            break
        }
        fmt.Println(value)
        //value := <-ch1
        //fmt.Println(value)
    }
}

func tel(ch chan <- int) {
    for i := 0; i < 15; i++ {
        ch <- i
    }
    close(ch)
}
