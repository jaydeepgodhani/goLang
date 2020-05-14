package main

import (
	"fmt"
)

//interface definition
//used on methods only (not function)
type VowelsFinder interface {
	FindVowels() []rune
}
type SalaryCalculator interface {
	CalculateSalary() int
}
type Tester interface {
	Test()
}
type Permanent struct { //this type implements salarycalc interface bcs of line 26
	empId    int
	basicpay int
	pf       int
}
type Contract struct { ////this type implements salarycalc interface bcs of line 31
	empId    int
	basicpay int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

//by declaring these two methods above both Permanent and Contraact implement salarycalculator interface
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

type MyString string ////this type implements VowelsFinder interface bcs of line 48
type MyFloat float64

//MyString implements VowelsFinder
//here type "MyString" implements vowelsfinder :(
func (ms MyString) FindVowels() []rune { //bcs it's method string cantbe used as parameter so custom Mystring is used
	//Explicitly declaration of implement interface is not needed in go and go interfaces are implemented implicitly if a type contains all the methods declared in the interface.
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
	//MyString implements vowelsfinder since it implements findvowel method
}
func (m MyFloat) Test() { //MyFloat implemetns Tester interface
	fmt.Println(m)
}
func describe(t Tester) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

//for empty interface
func desc(i interface{}) { // for empty interface write interface{}
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

//for type assertion
func assert(i interface{}) {
	s := i.(int) //get the underlying int value from i
	fmt.Println(s)
}

//assert that work for all type
func assert2(i interface{}) {
	s, ok := i.(int) //get the underlying int value from i
	fmt.Println(s, ok)
	/*
		If the concrete type of i is T then v will have the underlying value of i and ok will be true.
		If the concrete type of i is not T then ok will be false and v will have the zero value of type T and the program will not panic
	*/
}

//TYPE SWITCH
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

//Indirct inheritance
type SalaryCalculator2 interface {
	DisplaySalary2()
}

type LeaveCalculator2 interface {
	CalculateLeavesLeft2() int
}

type EmployeeOperations2 interface {
	SalaryCalculator2
	LeaveCalculator2
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary2() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft2() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	fmt.Println("ABOUT INTERFACE")
	//an interface is a set of method signatures
	name := MyString("Sam Anderson") //name is of type mystring
	var v VowelsFinder
	v = name                                    // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels()) // name.FindVowels also works

	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

	//Interface internal representation

	/*An interface can be thought of as being represented internally by a tuple (type, value). type is the underlying concrete type of the interface and value holds the value of the concrete type.*/
	var t Tester
	f := MyFloat(89.7)
	t = f //the concrete type of t is MyFloat and the value of t is 89.7
	describe(t)
	t.Test()
	/*
		An interface which has zero methods is called empty interface. It is represented as interface{}. Since the empty interface has zero methods, all types implement the empty interface.
	*/
	s := "Hello World"
	desc(s)
	i := 55
	desc(i)
	strt := struct {
		name string
	}{"Naveen"}
	desc(strt)

	//Type Assertion
	//Type assertion is used extract the underlying value of the interface.

	/*
	   i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T
	*/
	var s2 interface{} = 56 //if var s2 interface{} = "paul" then error bcs not an int ,line 77
	assert(s2)
	var s5 interface{} = 56
	assert2(s5)
	var s7 interface{} = "Steven Paul"
	assert2(s7)

	//type switch
	findType("Naveen")
	findType(77)
	findType(89.98)
	//A type can implement more than one interface
	//Embedding interfaces

	//Although go does not offer inheritance, it is possible to create a new interfaces by embedding other interfaces.
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmployeeOperations2 = e
	empOp.DisplaySalary2()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft2())
	//The zero value of a interface is nil. A nil interface has both its underlying value and as well as concrete type as nil
	/*
			type Describer interface {
				Describe()
			}
			func main() {
		    	var d1 Describer
		    	d1.Describe() //error here bcs no underlying value and type of interface
			}
	*/
}
