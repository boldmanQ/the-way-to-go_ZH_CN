package main

import "fmt"

type Simpler interface {
    Get() int
    Set(num int) Simple
}

type Simple struct {
    value int
}

func (s Simple) Get() int {
    return s.value
}

func (s Simple) Set(num int) Simple{
    s.value = num
    return s
}

func main() {
    a := &Simple{1}
    b := Simpler(a)
    fmt.Println(b.Get())
    fmt.Println(b.Set(2))
    fmt.Println(a)
}
