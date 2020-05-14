package main

import (
	"fmt"
)

func change(val *int) {
	*val = 55
}
func hello() *int {
	i := 5
	return &i
}
func modify(a *[3]int) { //for slice --> (a []int)
	(*a)[0] = 9 //or a[0]=9
}

//It is possible to declare structures without declaring a new type and these type of structures are called anonymous structures

func main() {
	fmt.Println("Pointers are different from C/C++ pointers")
	//Go does not support pointer arithmetic
	a := 234
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}
	size := new(int)
	fmt.Printf("value in size is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	*size++
	change(b)
	fmt.Println("New size value is", *size)
	fmt.Println("After change of B is", *b)
	fmt.Println("Original value of A is", a)
	//Return pointer of local variable from function
	/*
		The behavior of this code is undefined in programming languages such as C and C++ as the variable i goes out of scope once the function hello returns. But in the case of Go, the compiler does a escape analysis and allocates i on the heap as the address escapes the local scope
	*/
	db := hello()
	fmt.Printf("Value of DB is %d and type of DB is %T\n", *db, db)
	//Do not pass a pointer to an array as a argument to a function. Use slice instead
	arr := [3]int{3, 6, 5}
	modify(&arr) //use slice instead --> modify(arr[:])
	fmt.Println("After modify", arr)
	//NO ARITHMETIC ON POINTER like C/C++
	//p=&b
	//p++ gives error

	//**********STRUCTURE***********
	type employee struct { //named strucutre
		firstname, lastname string // can be written as separately like usual
		age, salary         int
	}
	emp1 := employee{
		firstname: "jd",
		salary:    35,
		age:       12,
		lastname:  "dj",
	} //order is not necessary here AND if brace is on new line then comma is necessary otherwise not
	emp2 := employee{"guy", "etr", 15, 27} //order must be necessary obviously
	fmt.Println("both employees are", emp1, emp2)
	//ANONYMOUS STRUCTURE
	//this structure is called anonymous because it only creates a new struct variable emp3 and does not define any new struct type.
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{"Andreah", "Nikola", 31, 5000} // same rules as above (comma vado rule ane lenghty declaration vado rule)
	fmt.Println("Employee 3", emp3)

	var emp4 employee //zero valued structure
	fmt.Println("Employee 4", emp4)
	//It is also possible to specify values for some fields and ignore the rest (rest are assigned as "zero values")
	fmt.Println("First Name:", emp1.firstname)
	fmt.Println("Last Name:", emp1.lastname)
	fmt.Println("Age:", emp1.age)
	fmt.Printf("Salary: $%d", emp1.salary)
	//pointer to structure
	emp8 := &employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp8).firstname)
	fmt.Println("Age:", (*emp8).age)
	//The language gives us the option to use emp8.firstName instead of the explicit dereference (*emp8).firstName to access the firstName field
	//Anonymous fields
	type Person struct {
		string
		int
	}
	p := Person{"Naveen", 50}
	fmt.Println(p)
	//Even though an anonymous fields does not have a name, by default the name of a anonymous field is the name of its type
	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)
	//NESTED STRUCTS
	type address struct {
		city, state string
	}
	type person struct {
		name string
		age  int
		adrs address
	}
	var p3 person
	p3.name = "Naveen"
	p3.age = 50
	p3.adrs = address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("P nested", p3)
	//Promoted fields
	/*
		Fields that belong to a anonymous struct field in a structure are called promoted fields since they can be accessed as if they belong to the structure which holds the anonymous struct field

		type person struct {
		name string
		age  int
		address //if this is an anonymous field then we can directly access p.city, p.state as if 		//they were directly member of person itself
		}
	*/
	//EXPORTED STRUCTS
	/*
		If a struct type starts with a capital letter, then it is a exported type and it can be accessed from other packages. Similarly if the fields of a structure start with caps, they can be accessed from other packages

		refer struct directory
	*/
	//Structs are value types and are comparable if each of their fields are comparable
	type name struct {
		firstName string
		lastName  string
	}
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}
}
