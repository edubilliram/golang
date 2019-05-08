package main

import "fmt"
import "time"
import "sync/atomic"
import "sync"

func main() {

	var count uint64
	wg := sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			for {

				atomic.AddUint64(&count, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	wg.Wait()

	countFinal := atomic.LoadUint64(&count)
	fmt.Println("count:", countFinal)
}

/*
dont use Sleep, it wont give the expected result all the time
may you have loaded the values before all the goroutines were done
using atomic is one way to synchronize the values between multiple go routines

otherwise we can use Mutex or Channels for synchronization

in this example atomic is efficient way to synchronize ….. as it’s more for counting

*/
