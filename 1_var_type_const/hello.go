package main

import(
	"fmt"
	"unsafe" //sizeof
	"math"
)

func main() {
	var age int
	//EACH VARIABLE OR PACKAGE MUST INCLUDE IN GO FILE
	//or we can declare as
	//var age int = 29
	//var age = 29  --> automatically declare age as Integer (Type Inference)
	var wd, hg = 100, 50
	//multiple type inference
	//var wd, hg int (default int is 0)
	fmt.Println("my age is", age)
	fmt.Println("width and height is", wd ,hg)
	//Multiple different declaration
	var(
		name="naveen"
		agee=29
		height int
	)
	/*
	Short Hand Declaration requires all the variable initialized
	name,age,height := "naveen",25,25
	Short hand syntax can only be used when at least one of the variables in the left side of := is newly declared
	Runtime declaration is also okay
	c := math.Min(a,b)
	*/
	fmt.Println("random",name,agee,height)
	//Formatted print
	fmt.Printf("type of age is %T and size is %d",age,unsafe.Sizeof(age))
	//Since Go is stringly typed language, variable assigned as Int cant be assigned as String again, even int+float is not allowed(without type casting)
	i := 23
	j := 46.567
	k := i + int(j)
	fmt.Println("\nvalue of k is",k)

	//const aa = math.Sqrt(4) //not allowed because const value must be known at compile time
	var aa = math.Sqrt(4)
	fmt.Println("value of aa is",aa)
	const hellu = "Hello World"
	var firstt = true
	//const hellu does not have any type
	fmt.Println("value of hellu is",hellu,firstt) //type of hellu is UNTYPED
	//following code creates typed constant
	const typedhello string = "Helludlrow"
	//Interesting
	/*
	var defaultName = "Sam" //allowed  --> type of defaultName is string , "Sam" is untyped constant
    type myString string			   --> custom type as string
	var customName myString = "Sam" //allowed --> type of customName is myString
	//Sam is untyped it can be assigned to any string variable
	customName = defaultName //not allowed --> assignment of string to customName is not allowed
	..
	true and false is also untyped, so rules for bool are same as string
	*/
	//const another int = 5  -->this does have type of int const
	/*
	const a = 5
    var intVar int = a
    var int32Var int32 = a
    var float64Var float64 = a
	var complex64Var complex64 = a
	all valid
	*/
	var a = 5.9/8
	fmt.Printf("a's type %T value %v",a, a)
	var total = calc(20,5)
	fmt.Println("\ntotal price is",total)
	area,peri:=mulret(3,4) //for multi-return all returned var should be caught
	fmt.Println("\narea and parimeter is",area,peri)
	a1,p1:=rectProps(4,5)
	fmt.Println("\narea and parimeter is",a1,p1)
	// _ is a blank identifier
	//what if we only want to use area but not a perimeter
	aa2 , _ :=rectProps(4,5)
	fmt.Println("\narea and BLANK IDEN is",aa2)
}
func calc(q,a int) int{
	return q*a
}
func mulret(h,w int)(int ,int){ //for multiple return () is necessary for return type
	var area = h*w
	var peri = 2*(h+w)
	return area,peri
}
func rectProps(length, width float64)(area, perimeter float64) {  //named return
    area = length * width //same name as return name
    perimeter = (length + width) * 2 //same name as return name
    return //no explicit return value
}