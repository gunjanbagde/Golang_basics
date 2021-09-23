package main

import(
	"fmt"
	"time"
)

func Sendvalue (c chan int){
	fmt.Println("Executing goroutine")
	time.Sleep(1* time.Second)
	c <- 6
	c <- 3
	fmt.Println("Finished Goroutine")

}


func main(){
	fmt.Println("Go Channels")

	v:= make(chan int ,2) //bufferd channels
	go Sendvalue(v)
	go Sendvalue(v)

	rcvvalue:= <- v
	rcvvalue2:= <- v
	fmt.Println(rcvvalue)
	fmt.Println(rcvvalue2)

	time.Sleep(1* time.Second)


//unidirectional	
// func main() {
// cs := make(chan<- int) // :- Send only
// cr := make(<-chan int) // :- receive only

// 	go func() {
// 		cs <- 42
// 	}()
// 	fmt.Println(<-cs)

// 	fmt.Printf("------\n")
// 	fmt.Printf("cs\t%T\n", cs)

// general to specific assigns buy doesnt work the other way around
// cr = c
// cs = c


}