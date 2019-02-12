package main

import "fmt"
import "time"

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
    time.Sleep(2e9)
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}

// Washington Tripoli London Beijing Tokio
