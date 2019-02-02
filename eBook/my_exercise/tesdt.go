package main
import "fmt"

func main() {
    a := []int{1,2,3}
    s := a[3:3+1]
    fmt.Println(len(s))
}
