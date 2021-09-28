package main

import (
	"fmt"
	"time"
)

func main() {
	var start string
	var end string
	var check string

	fmt.Println(`enter start time in "hh:mm:ss" format`)
	fmt.Scanln(&start)
	fmt.Println(`enter end time in "hh:mm:ss" format`)
	fmt.Scanln(&end)
	fmt.Println(`enter check time in "hh:mm:ss" format`)
	fmt.Scanln(&check)

	layout := "15:04:05"
	stime, _ := time.Parse(layout, start)
	etime, _ := time.Parse(layout, end)
	ctime, _ := time.Parse(layout, check)

	//fmt.Println(stime,"\n",etime,"\n",ctime)

	if ctime.After(stime) && ctime.Before(etime) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
