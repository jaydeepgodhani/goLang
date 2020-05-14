package main

import (
	"fmt"
)

type add func(a int, b int) int

func simple(a func(a, b int) int) {
	//simple function take function as argument and returned nothing
	fmt.Println(a(60, 7))
}

func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
func main() {
	//A language which supports first class functions allows functions to be assigned to variables,passed as arguments to other functions and returned from other functions. Go has support for first class functions.
	a := func() {
		fmt.Println("hello world first class function")
	}
	//These kind of functions are called anonymous functions since they do not have a name
	a()
	fmt.Printf("%T", a)
	//It is also possible to call a anonymous function without assigning it to a variable
	func() {
		fmt.Println("hello world first class function")
	}()
	//It is also possible to pass arguments to anonymous functions just like any other function
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")

	//Just like we define our own struct types, it is possible to define our own function types
	var a2 add = func(a int, b int) int {
		return a + b
	}
	s := a2(5, 6)
	fmt.Println("Sum", s)

	/*a Higher Order function which does at least one of the following
	takes one or more functions as arguments
	returns a function as its result
	*/
	f := func(a, b int) int {
		return a + b
	}
	simple(f)

	s2 := simple2()
	fmt.Println(s2(16, 7))

	//Closure
	//Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function
	outside := 07
	func() {
		fmt.Println("outside variable is =", outside)
	}() //the anonymous function accesses the variable a which is present outside its body, Hence this anonymous function is a closure

	/*
		In the program above, the function appendStr returns a closure. This closure is bound to the variable t. Let's understand what this means.
		The variables a and b declared in line nos. 17, 18 are closures and they are bound to their own value of t.
		We first call a with the parameter World. Now the value of a's version of t becomes Hello World.
		In line no. 20 we call b with the parameter Everyone. Since b is bound to its own variable t, b's version of t has a initial value of Hello again. Hence after this function call, the value of b's version of t becomes Hello Everyone
	*/
	ax := appendStr()
	bx := appendStr()
	fmt.Println(ax("World"))
	fmt.Println(bx("Everyone"))

	fmt.Println(ax("Gopher"))
	fmt.Println(bx("!"))
}
