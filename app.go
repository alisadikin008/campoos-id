package main

import (
	"fmt"
	"learn-golang/helper"
)

var months, weeks = helper.GetMonths(), helper.GetWeeks()

func main() {
	// get all months
	for key, value := range months {
		fmt.Println(key, value)
	}
	fmt.Println("")
	//get all weeks
	for key, value := range weeks {
		fmt.Println(key, value)
	}
}
