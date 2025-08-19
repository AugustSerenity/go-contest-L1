package main

import (
	"fmt"
	"time"
)

func mySleep(n time.Duration) {
	end := time.Now().Add(n)
	for time.Now().Before(end) {

	}
}

func main() {
	var n time.Duration
	fmt.Println("enter time in seconds")
	fmt.Scan(&n)

	timeNow := time.Now()

	mySleep(n * time.Second)

	timeAfter := time.Since(timeNow)
	fmt.Println(timeAfter)
}
