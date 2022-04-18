package main

import (
	"fmt"
	"time"
)

func main() {
	total := 0
	var workers = make(chan struct{}, 1)

	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}
		go func(job int) {
			defer func() {
				<-workers
			}()
			total++
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(total)
}
