package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Create("names.txt")
	f, err = os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	d1:=[]byte("hello\n World \n")

	err = os.WriteFile("names.txt", d1, 0644)

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}