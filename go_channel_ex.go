package main

import (
	"fmt"
)

func sum(arr []int, ch chan int) {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	ch <- sum
}

func main() {
	arr := []int{2, 4, 8, -7, 2, 0}
	c := make(chan int)
	go sum(arr[:len(arr)/2], c)
	go sum(arr[len(arr)/2:], c)
	x, y := <-c, <-c
	fmt.Println("sum of the first half of the array is : ", x, "\nsum for second half of the array is : ", y)
}

// sum of the first half of the array is :  -5 
// sum for second half of the array is :  14

// Program exited.
