package main

import "fmt"
import "sync"

func main() {

	var mutex sync.Mutex
	var count int
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			count++
			mutex.Unlock()
		}()
	}
	wg.Wait() /// use WaitGroups to wait till all goroutines are done ...... useing time.Sleep() will give different results everytime

	fmt.Println("count", count)
}
