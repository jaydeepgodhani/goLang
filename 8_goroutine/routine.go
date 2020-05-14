package main

import (
	"fmt"
	"sync"
	"time"
)

func hello() {
	fmt.Println("in hello func")
}
func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func hello2(done chan bool) {
	fmt.Println("\nhello goroutine hello2")
	done <- true
}
func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

//for unidirectional channel
func sendData(sendch chan<- int) { //send only channel
	sendch <- 10
}
func sendData2(sendch chan<- int) {
	sendch <- 20
}

//closing channel
func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

//waitgroup
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

//select statement
func server1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from server2"
}

//for default case
func process2(ch3 chan string) {
	time.Sleep(1000 * time.Millisecond)
	ch3 <- "process successful"
}

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func main() { //main function is running its own goroutine (main goroutine)
	go hello() //here hello function starts its own goroutine, ans concurrently with main function
	fmt.Println("main function")
	//Two main properties of GoRoutine
	/*
		When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the Goroutine call and any return values from the Goroutine are ignored.

		The main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.
	*/
	/*
	 After the call to go hello(), the control returned immediately to the next line of code without waiting for the hello goroutine to finish and printed main function.
	 Then the main Goroutine terminated since there is no other code to execute and hence the hello Goroutine did not get a chance to run.
	*/
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("another main function")
	//Channels can be used to block the main Goroutine until all other Goroutines finish their execution
	//starting multiple goroutines
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
	//Each channel has a type associated with it. This type is the type of data that the channel is allowed to transport
	//The zero value of a channel is nil. nil channels are not of any use and hence the channel has to be defined using make similar to maps and slices.
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
	//or just  -->  a := make(chan int)
	//Sends and receives are blocking by default
	/*
		When a data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel. Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.
		This property of channels is what helps Goroutines communicate effectively without the use of explicit locks or conditional variables
	*/
	done := make(chan bool)
	go hello2(done)
	<-done //we're receiving data from done channel
	//This line of code is blocking which means that until some Goroutine writes data to the done channel, the control will not move to the next line of code
	fmt.Println("main channel function")

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
	//all channels are created so far is bidirectional

	sendch := make(chan<- int) //send only channel
	go sendData(sendch)
	fmt.Println(sendch) // Println(<-sendch) will give error, bcs cant read from send only channel
	//All is well but what is the point of writing to a send only channel if it cannot be read from!
	/*
		This is where channel conversion comes into use. It is possible to convert a bidirectional channel to a send only or receive only channel but not the vice versa.
	*/
	chnl := make(chan int)
	go sendData2(chnl)
	fmt.Println(<-chnl)
	//Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel
	//Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed.
	ch := make(chan int)
	go producer(ch)
	for { //for v := range ch  -->  this also works
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
	/*
		this is also deadlock, think why
			asd := make(chan int)
			asd <- 3
			b := <-asd
			fmt.Printf("%T", b)
	*/
	// Till here all the channels are unbuffered
	// Sends and receives to an unbuffered channel are blocking
	// Sends to a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty
	// Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer
	//capacity in the above syntax should be greater than 0 for a channel to have a buffer. The capacity for an unbuffered channel is 0 by default and hence we omitted the capacity previously

	ch2 := make(chan string, 4)
	ch2 <- "naveen"
	ch2 <- "paul"
	//ch2 <- "anything"  -->  this would panic goprogram
	//capacity is CAPACITY and length is number of elements stored in it
	fmt.Printf("length is %d and capacity is %d \n", len(ch2), cap(ch2))
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Printf("after length is %d and capacity is %d", len(ch2), cap(ch2))
	//A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is blocked until all Goroutines finish executing.

	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
	/*
		WaitGroup is a struct type and we are creating a zero value variable of type WaitGroup in line no.18. The way WaitGroup works is by using a counter. When we call Add on the WaitGroup and pass it an int, the WaitGroup's counter is incremented by the value passed to Add. The way to decrement the counter is by calling Done() method on the WaitGroup. The Wait() method blocks the Goroutine in which it's called until the counter becomes zero.
	*/
	//It is important to pass the address of wg. If the address is not passed, then each Goroutine will have its own copy of the WaitGroup and main will not be notified when they finish executing.
	//WORKER POOL
	//a worker pool is a collection of threads which are waiting for tasks to be assigned to them. Once they finish the task assigned, they make themselves available again for the next task

	//**********SELECT***********
	/*
		The select statement is used to choose from multiple send/receive channel operations. The select statement blocks until one of the send/receive operation is ready. If multiple operations are ready, one of them is chosen at random. The syntax is similar to switch except that each of the case statement will be a channel operation
	*/
	//This way we can send the same request to multiple servers and return the quickest response to the user
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
	//default case
	ch3 := make(chan string)
	go process2(ch3)
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case v := <-ch3:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
	//the default case will be executed even if the select has only nil channels
	//When multiple cases in a select statement are ready, one of them will be executed at random
	/*
		func main(){
			select{}
		}
		this program gives error bcs
		the select statement will block until one of its cases is executed. In this case the select statement doesn't have any cases and hence it will block forever resulting in a deadlock
		fatal error: all goroutines are asleep - deadlock!
	*/
}
