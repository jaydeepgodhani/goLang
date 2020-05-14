package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func main() {
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
	/*
		The main function creates the data channel from which random numbers are read from and written to in line no. 41. The done channel in line no. 42 is used by the consume goroutine to notify main that it is done with its task. The wg waitgroup in line no. 43 is used to wait for all the 100 goroutines to finish generating random numbers.

		The for loop in line no. 44 creates 100 goroutines. The goroutine call in line no. 49 calls wait() on the waitgroup to wait for all 100 goroutines to finish creating random numbers. After that it closes the channel. Once the channel is closed and the consume goroutine has finished writing all generated random numbers to the file, it writes true to the done channel in line no. 37 and the main goroutine is unblocked and prints File written successfully
	*/
}
