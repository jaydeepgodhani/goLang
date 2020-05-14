package main

import (
	"fmt"
)

type myint int

type Employee struct {
	name     string
	salary   int
	currency string
	age      int
}

//use alias instead of just plain int
func (a myint) add(b myint) myint {
	return a + b
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

//value receiver
func (e Employee) changeName(newName string) {
	e.name = newName
}

//pointer receiver
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}
func main() {
	fmt.Println("METHODS")
	//SIGNATURE : func (t Type) methodName(params)
	//To define a method on a type, the definition of the receiver type and the definition of the method should be present in the same package
	//struct is user define type and that's why they worked here
	//but what about int,float type? bcs they're defined in diff package
	//workaround for that is to crerate alias for int
	//A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type
	//receiver can be accessed within the method.
	emp1 := Employee{"sam", 5909, "$", 32}
	emp1.displaySalary() //Calling displaySalary() method of Employee type***********
	//Same behaviour can be achieved using functions also
	//so WHY METHODS????
	/*
		Go is not a pure object-oriented programming language and it does not support classes. Hence methods on types are a way to achieve behavior similar to classes. Methods allow a logical grouping of behavior related to a type similar to classes. In the above sample program, all behaviors related to the Employee type can be grouped by creating methods using Employee receiver type. For example, we can add methods like calculatePension, calculateLeaves and so on.

		Methods with the same name can be defined on different types whereas functions with the same names are not allowed. Let's assume that we have a Square and Circle structure. It's possible to define a method named Area on both Square and Circle
	*/
	//Pointer Receivers vs Value Receivers
	//The difference between value and pointer receiver is, changes made inside a method with a pointer receiver is visible to the caller whereas this is not the case in value receiver
	e := Employee{"heya", 23345, "ruty", 23}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew") //passing value
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).changeAge(51) //passing pointer
	e.changeAge(50)    //passing pointer differently, here just receiver must be pointer
	fmt.Printf("\nEmployee age after change: %d", e.age)
	//When to use what
	//Pointers receivers can also be used in places where it's expensive to copy a data structure
	//Methods belonging to anonymous fields of a struct can be called as if they belong to the structure where the anonymous field is defined.

	//Value receivers in methods vs Value arguments in functions
	/*
		When a function has a value argument, it will accept only a value argument.
		When a method has a value receiver, it will accept both pointer and value receivers
	*/

	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()

	p := &r
	/*
	   compilation error, cannot use p (type *rectangle) as type rectangle
	   in argument to area
	*/
	//area(p)
	//here compiler converted below line into (*p).area() bcs there is a value receiver
	p.area() //calling value receiver with a pointer
	/*
		Similar to value arguments, functions with pointer arguments will accept only pointers whereas methods with pointer receivers will accept both pointer and value receiver.
	*/
	perimeter(p)
	p.perimeter()

	/*
	   cannot use r (type rectangle) as type *rectangle in argument to perimeter
	*/
	//perimeter(r)

	//interpreted by the language as (&r).perimeter()
	r.perimeter() //calling pointer receiver with a value
	//for int,string,float built-in types
	num1 := myint(5)
	num2 := myint(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}
