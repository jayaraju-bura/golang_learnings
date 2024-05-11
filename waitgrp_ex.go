package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(d time.Duration, wg *sync.WaitGroup) {
	fmt.Println("starting the work")
	time.Sleep(d)
	fmt.Println("work done")
	wg.Done()
}

func main() {
	fmt.Println("Hello, 世界")
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go doWork(time.Second*2, &wg)
	go doWork(time.Second*4, &wg)
	wg.Wait()
	fmt.Printf("work has taken %v seconds \n", time.Since(start))
}
