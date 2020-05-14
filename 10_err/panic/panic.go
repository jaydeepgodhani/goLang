package main

import (
	"fmt"
)
/*
One important factor is that you should avoid panic and recover and use errors where ever possible. Only in cases where the program just cannot continue execution should a panic and recover mechanism be used
There are two valid use cases for panic.

    An unrecoverable error where the program cannot simply continue its execution.
    One example would be a web server which fails to bind to the required port. In this case it's reasonable to panic as there is nothing else to do if the port binding itself fails.

    A programmer error.
    Let's say we have a method which accepts a pointer as a parameter and someone calls this method using nil as argument. In this case we can panic as it's a programmer error to call a method with nil argument which was expecting a valid pointer.

*/
func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil") //program stops here and print stacktrace
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	//DEFER while panicing
	/*
		When a function encounters a panic, its execution is stopped, any deferred functions are executed and then the control returns to its caller. This process continues until all the functions of the current goroutine have returned at which point the program prints the panic message, followed by the stack trace and then terminates
	*/
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
	//recover is a builtin function which is used to regain control of a panicking goroutine.
	//Recover is useful only when called inside deferred functions. Executing a call to recover inside a deferred function stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic. If recover is called outside the deferred function, it will not stop a panicking sequence.
}
