// Реализуйте функцию для разблокировки мьютекса с помощью defer
package main

import (
	"fmt"
	"sync"
)

const threads = 1000

func main() {
	var (
		count int
		mutex = sync.Mutex{}
		wg    sync.WaitGroup
	)
	wg.Add(threads)

	for i := 0; i < threads; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Completed", count)
}
