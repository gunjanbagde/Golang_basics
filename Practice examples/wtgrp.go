package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup

func main(){
		fmt.Println("OS\t\t",runtime.GOOS)
		fmt.Println("Architecture\t",runtime.GOARCH)
		fmt.Println("CPUs\t\t",runtime.NumCPU())
		fmt.Println("Goroutines\t",runtime.NumGoroutine())

		wg.Add(1)
		go one()
		two()
		fmt.Println("CPUs\t\t", runtime.NumCPU())
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
		wg.Wait()
}

func one() {
	for i := 0; i < 10; i++ {
		fmt.Println("one:", i)
	}
	wg.Done()
}

func two() {
	for i := 0; i < 10; i++ {
		fmt.Println("two:", i)
	}
}
