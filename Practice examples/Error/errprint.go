package main

import (
	"log"
	"os"
	//"fmt"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//	fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)     //stops the execution then even defer doesent run
		//		log.Panicln(err)
		//		panic(err)
	}
}
