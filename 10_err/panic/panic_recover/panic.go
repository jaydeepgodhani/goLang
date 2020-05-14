package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}
func a() {
	defer recovery()
	fmt.Println("Inside A")
	b() //If the function b() was called in the same goroutine then the panic would have been recovered
	// or if the funxtion go b() was called then Panic
	time.Sleep(1 * time.Second)
}
func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
func fullName(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

//runtime panic
func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}
func arun() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	// Recover works only when it is called from the same goroutine. It's not possible to recover from a panic that has happened in a different goroutine
	a()
	fmt.Println("normally returned from main")

	//we can also recover from RunTime panic
	arun()
	fmt.Println("normally returned from main arun")

	//Getting stack trace after recover
	//If we recover a panic, we loose the stack trace about the panic. Even in the program above after recovery, we lost the stack trace.
	//There is a way to print the stack trace using the PrintStack function of the Debug package
}
