package main

import (
    "fmt"
    "math"
)

func main() {
    var arr1 = []int{2,4,6,7,5,4}
    fmt.Println(minSlice(arr1))
    fmt.Println(maxSlice(arr1))
}

func minSlice(arr []int) (min int) {
    min = math.MaxInt32
    for _, val := range arr {
        if min > val {
            min = val
        }
    }
    return
}


func maxSlice(arr []int) (max int) {
    max = math.MinInt32
    for _, val := range arr {
        if max < val {
            max = val
        }
    }
    return
}
