package main

import "fmt"

func generate() chan int {
    ch := make(chan int)
    go func() {
        for i := 2; ; i++ {
            ch <- i
            }
    }()
    return ch
}

//func filter(int chan int, prime int) chan int {
//    out := make(chan int)
//    go func() {
//        for {
//            if i := <- in; i%prime != 0 {
//                out <- i
//            }
//        }
//    }()
//    return out
//}

func main() {
    ch := generate()
    for {
        fmt.Println(<-ch)
    }
}
