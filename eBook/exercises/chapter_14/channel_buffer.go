package main

import "fmt"
import "time"

func main() {
	c := make(chan int, 1)
	go func() {
		time.Sleep(5 * 1e8)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
    time.Sleep(1e9)
}

/* Output:
sending 10
sent 10   // prints immediately
no further output, because main() then stops
*/
