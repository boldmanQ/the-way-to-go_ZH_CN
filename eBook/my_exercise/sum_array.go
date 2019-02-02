package main
import "fmt"

func main() {
    b := []float32{1.1, 2.2}
    //c := Sum(b)
    fmt.Println(Sum(b))
}

func Sum(a []float32) (sum float32) {
    for _, val := range a {
        sum += val
    }
    return
}
