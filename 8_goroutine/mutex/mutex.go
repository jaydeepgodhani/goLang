package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // ch <- true
	x = x + 1
	m.Unlock() // <- ch  -->  will also work if channel is of capacity 1
	wg.Done()
}
func main() {
	fmt.Println("MUTEX")
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
	//It is important to pass the address of the mutex in line no. 18. If the mutex is passed by value instead of passing the address, each Goroutine will have its own copy of the mutex and the race condition will still occur
	//can also be done using channel
	//In general use channels when Goroutines need to communicate with each other and mutexes when only one Goroutine should access the critical section of code
	//In the case of the problem which we solved above, I would prefer to use mutex since this problem does not require any communication between the goroutines. Hence mutex would be a natural fit
}
