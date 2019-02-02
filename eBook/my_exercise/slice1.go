package main

import "fmt"

func main() {
    s:= make([]byte, 5)
    s = s[2:4]
    fmt.Println(len(s), cap(s))
    s1 := []byte{'p', 'o', 'e', 'm'}
    s2 := s1[2:]
    fmt.Printf("%v", s2)
    s2[1] = 't'
    fmt.Printf("%v", s2)
}
