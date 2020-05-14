package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

//these wg.Done() calls happen just before the area method returns. wg.Done() should be called before the method returns irrespective of the path the code flow takes and hence these calls can be effectively replaced by a single defer call
func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done() // this will ensure that wg.Done() call happen just before return executes
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		//wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		//wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	//wg.Done()
}

//Defer is used in places where a function call should be executed irrespective of the code flow. Lets understand this with the example of a program which makes use of WaitGroup
func main() {
	//in case of exception also defer will be called
	fmt.Println("DEFER USE")
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
