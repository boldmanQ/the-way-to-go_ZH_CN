package main
import "fmt"
import "unicode/utf8"


func main() {
    a := "asSASA ddd dsjkdsjs dk"
    a = "1"
    b := "asSASA ddd dsjkdsjsこん dk"
    fmt.Printf("字节数：a %d\n", len(a))
    fmt.Printf("字节数：b %d\n", len(b))
    fmt.Printf("字符数：a %d\n", utf8.RuneCountInString(a))
    fmt.Printf("字符数：b %d\n", utf8.RuneCountInString(b))
}
