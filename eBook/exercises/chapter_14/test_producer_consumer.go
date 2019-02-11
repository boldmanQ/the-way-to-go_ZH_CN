package main

import "fmt"

func sender(ch chan int) {
    for i:= 0; i < 10; i++ {
        ch <- i * 10
    }
    close(ch)
}

func receiver(ch chan int, ch2 chan bool) {
    for num:= range ch {
        fmt.Println(num)
    }
    ch2 <- true
}

func main() {
    num := make(chan int)
    done := make(chan bool)
    go sender(num)
    go receiver(num,done)
    <-done
}
