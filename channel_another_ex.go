package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork(d time.Duration, resch chan string) {
	fmt.Println("starting the work")
	time.Sleep(d)
	fmt.Println("work done")
	resch <- fmt.Sprintf("The result of the work done is %d", rand.Intn(100))
	wg.Done()
}

var wg *sync.WaitGroup

func main() {
	fmt.Println("Hello, 世界")
	start := time.Now()
	resch := make(chan string)
	wg = &sync.WaitGroup{}
	wg.Add(2)
	go doWork(time.Second*2, resch)
	go doWork(time.Second*4, resch)

	go func() {
		for result := range resch {
			fmt.Printf("result fetched from channel %s \n", result)
		}
		fmt.Printf("work has taken %v seconds \n", time.Since(start))
	}()
	wg.Wait()
	close(resch)
	time.Sleep(time.Second)

}
