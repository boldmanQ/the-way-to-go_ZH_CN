package main

import "fmt"


func fib() func() {
    a,b := 1
    a, b = b, b+a
    return f
