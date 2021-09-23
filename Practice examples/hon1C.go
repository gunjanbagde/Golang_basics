package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup

func main(){
	fmt.Println("GOroutines ", runtime.NumGoroutine())
	fmt.Println("ÏN main function")
	wg.Add(2)
	go one()
	go two()
	fmt.Println("GOroutines ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("GOroutines ", runtime.NumGoroutine())
}

func one(){
	fmt.Println("first function")
	wg.Done()
}


func two(){
	fmt.Println("second function")
	wg.Done()
}