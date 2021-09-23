package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c   //returns true if a channel is closed and false other way around

	fmt.Println(v, ok)

	v, ok = <-c

	fmt.Println(v, ok)
}
