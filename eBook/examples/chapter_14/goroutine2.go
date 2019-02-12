package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

    time.Sleep(4e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
         fmt.Println("sleep 5e8")
	    time.Sleep(5e8)
	}
}

// Washington Tripoli London Beijing Tokio
