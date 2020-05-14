package main

import (
	"fmt"
)

func finished() {
	fmt.Println("Finished finding largest")
}

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}
func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
	//defer function call happen here
}
func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}
func main() {
	fmt.Println("DEFER and ERROR")
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)

	//defer for method
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ") // if this was the last line of code then output will be "welcome john smith"

	//The arguments of a deferred function are evaluated when the defer statement is executed and not when the actual function call is done.
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)
	//here defer function takes value of A as 5 not 10

	//When a function has multiple defer calls, they are added on to a stack and executed in Last In First Out (LIFO) order.
	name := "Naveen"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
