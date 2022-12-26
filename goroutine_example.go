package main

import (
	"fmt"
	"time"
)

func fun(str string) {
	for i := 0; i < 2; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str, " : ", i)
	}
}

func main() {
	fun("direct")
	go fun("goroutine")
	go func(msg string) {
		time.Sleep(2 * time.Second)
		fmt.Println("===>>", msg)
	}("going")
	time.Sleep(10 * time.Second)
	fmt.Println("Hello, World")
}


// direct  :  0
// direct  :  1
// goroutine  :  0
// ===>> going
// goroutine  :  1
// Hello, World

// Program exited.
