package main
import(
    "fmt"
)
func main(){
	fmt.Println("ABOUT CONDITIONS SWITCH,FOR")
	var a = 24
	if a%2==0{
		fmt.Println("even");
	} else { //else must start after this closing brace of IF why?
		//because GO automatically insert semicolon after cloasing brace } if that is the last token
		fmt.Println("odd");
	}
	/*
	if num := 10; num % 2 == 0 { //checks if number is even
        fmt.Println(num,"is even") 
    }  else {
        fmt.Println(num,"is odd")
	}
	visibility of num is restricted to IF ELSE block only
	*/
	//only one loop FOR (no while , no do while)
	//all three conditions are optional
	for i := 1; i <= 20; i++ { //i cant be accessed outside of loop
		if i>10{
			break
		}
		if i%2==0{
			continue
		}
        fmt.Printf("%d ",i)
	}
	outer:  //label for breaking this point
    for i := 0; i < 3; i++ {
        for j := 1; j < 4; j++ {
            fmt.Printf("i = %d , j = %d\n", i, j)
            if i == j {
                break outer //break this point
            }
        }
	}
	i := 0
    for i <= 10 { //semicolons are ommitted and only condition is present, WHILE loop equivalent
        fmt.Printf("%d ", i)
        i += 2
	}
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
        fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
	//In Go the control comes out of the switch statement immediately after a case is executed
	finger := 4
    switch finger { // switch finger:=8; finger {  --> this also works and finger variable cant be accessed outside of switch
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")
	default: //default case
        fmt.Println("incorrect finger number")
	}
	//multiple declaration on case
	letter := "i"
    switch letter {
    case "a", "e", "i", "o", "u": //multiple expressions in case
        fmt.Println("vowel")
    default:
        fmt.Println("not a vowel")
	}
	//expressionless switch == ifelse ladder
	num := 75
    switch { // expression is omitted
    case num >= 0 && num <= 50:
        fmt.Println("num is greater than 0 and less than 50")
    case num >= 51 && num <= 100:
        fmt.Println("num is greater than 51 and less than 100")
    case num >= 101:
        fmt.Println("num is greater than 100")
	}
	//check whether the input number is lesser than 50, 100 or 200 (multiple case check using FALLTHROUGH)
	switch num := 75; { //num is not a constant
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough //print next case's stmt(without checking case) even if this case held true
    case num < 20: //fallthrough should be the last stmt
		fmt.Printf("%d is lesser than 200", num)
	case num < 30: //not executing this stmt bcs no fallthrough in above case
        fmt.Printf("%d is lan 200", num)
	}
	fmt.Println("\nABOUT ARRAYS and SLICES")
	//Arrays in Go are value types and not reference types
	//when they are assigned to a new variable, a copy of the original array is assigned to the new variable
	var aa[3] int //The size of the array is a part of the type eg. [3]int, [5]int
	ab := [3]int{12} //other two values are automatically assigned as zero
	ac := [...]int{2,3,7,9,3,7,1,6} //auto assign length
	fmt.Println(aa)
	fmt.Println(ab)
	fmt.Println(ac)
	anew := ac
	anew[2]=564
	sum:=int(0)
	fmt.Println(anew,"and length is",len(anew))
	//RANGED FOR LOOP used for iterating array, must use two variable for index and value
	for i, v := range anew {//range returns both the index and value
		//to ignore any variable use _ operator like --> for _,v :=range anew
        fmt.Printf("%d the element of anew is %d\n", i, v)
        sum += v
    }
	fmt.Println("\nsum of all elements of anew",sum)
	//slice also can have multidiemsions but NO SIZE specified just the [][] is enough
	amulti := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. why? bcs of THAT semicolon rule
	}
	fmt.Println(amulti)
	printarray(amulti)
	

	//SLICES YO SLICES YO SLICES

	//Any modifications done to the slice will be reflected in the underlying array
	/*
	The length of the slice is the number of elements in the slice. The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created
	*/
	var b [] int = ac[1:4] // index 1 to index 3 that means 3,7,9
	fmt.Println("slice b is ",b)
	c := []int{6, 7, 8} //creates an array and returns a slice reference to C
	fmt.Println("slice c ",c)

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before",darr)
    for i := range dslice { // or for i,_ := range dslice
        dslice[i]++
    }
	fmt.Println("array after",darr)
	
	numa := [3]int{78, 79 ,80}
    nums1 := numa[:] //creates a slice which contains all elements of the array
    nums2 := numa[:]
    fmt.Println("array before change 1",numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1", numa)
    nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
	
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6
	//A slice can be re-sliced upto its capacity. Anything beyond that will cause the program to throw a run time error
	fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
	
	/*
	func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity. The capacity parameter is optional and defaults to the length. The make function creates an array and returns a slice reference to it.
	*/
	i2 := make([]int, 5, 5)
	fmt.Println(i2)

	//Appending to the slice
	cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
    cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	
	var names []string //zero value of a slice is nil
    if names == nil {
        fmt.Println("slice is nil going to append") //append is variadic funtion
        names = append(names, "John", "Sebastian", "Vinay")
        fmt.Println("names contents:",names)
	}
	//It is also possible to append one slice to another using the ... operator
	veggies := []string{"potatoes","tomatoes","brinjal"}
    fruits := []string{"oranges","apples"}
    food := append(veggies, fruits...)
	fmt.Println("food:",food)
	//passing slice to a function
	//PASS BY REFERENCE
	nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside
	
	//array cant be garbage collected if it is referenced by slice
	//we can copy that slice to another slice and original array can garbage collected
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2]
    countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	fmt.Println(countriesCpy)

	//VARIADIC FUNCTION
	find(89, 89, 90, 95) //or find(89, []int{89, 90, 95})
	xyt := []int{56,67,78,89}
	find(67, xyt...) //tis is how we pass slice to the variadic function
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)
}
func printarray(a [3][2]string) {  //parameter array, why? bcz size specified
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
	}
}
func subtactOne(numbers []int) {  //parameter slice, why? no size specified
    for i := range numbers {
        numbers[i] -= 2
    }

}
func find(num int, nums ...int) {  //these nums are converted to slice, compiler converts into slice and passes it into find function
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}

// ******ORIGINAL FILE*******


// package main

// import (
// 	"fmt"
// )

// /*
// Arrays are value types. Changes made inside a function are not
// visible to the caller
// */
// func changeLocal(num [5]int) {
// 	num[0] = 55
// 	fmt.Println("inside function ", num)

// }

// /*
// iterating multidimensional arrays
// */
// func printarray(a [3][2]string) {
// 	for _, v1 := range a {
// 		for _, v2 := range v1 {
// 			fmt.Printf("%s ", v2)
// 		}
// 		fmt.Printf("\n")
// 	}
// }

// /*
// Changes made to a slice inside a function are visible
// to the caller
// */
// func subtactOne(numbers []int) {
// 	for i := range numbers {
// 		numbers[i] -= 2
// 	}

// }

// /*
// use copy function to copy the contents from one slice to another
// helful to optimize memory since the original array can now be garbage collected
// */
// func countries() []string {
// 	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
// 	neededCountries := countries[:len(countries)-2]
// 	countriesCpy := make([]string, len(neededCountries))
// 	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
// 	return countriesCpy
// }

// func main() {
// 	/*
// 	array declaration
// 	*/
// 	fmt.Println("array declaration")
// 	var a [3]int //int array with length 3
// 	fmt.Println(a)

// 	var b [3]int //int array with length 3
// 	b[0] = 12    // array index starts at 0
// 	b[1] = 78
// 	b[2] = 50
// 	fmt.Println(b)

// 	/*
// 	short hand declaration to create array
// 	*/
// 	fmt.Println("\nshort hand declaration")
// 	c := [3]int{12, 78, 50}
// 	fmt.Println(c)

// 	/*
// 	... syntactic sugar to make the compiler determine the length
// 	*/
// 	fmt.Println("\nsyntactic sugar to determine length")
// 	d := [...]int{12, 78, 50}
// 	fmt.Println(d)

// 	/*
// 		  e := [3]int{5, 78, 8}
// 		  var f [5]int
// 			e = f //not possible since [3]int and [5]int are distinct types
// 	*/

// 	/*
// 	arrays are value types
// 	*/
// 	fmt.Println("\nArrays are value types")
// 	g := [...]string{"USA", "China", "India", "Germany", "France"}
// 	h := g // a copy of a is assigned to b
// 	g[0] = "Singapore"
// 	fmt.Println("g is ", g)
// 	fmt.Println("h is ", h)

// 	/*
// 	Arrays are value types. Changes made in a function are not
// 	visible to the caller
// 	*/
// 	fmt.Println("\nChanges made to an array inside a function are not visible to the caller")
// 	num := [...]int{5, 6, 7, 8, 8}
// 	fmt.Println("before passing to function ", num)
// 	changeLocal(num) //num is passed by value
// 	fmt.Println("after passing to function ", num)

// 	/*
// 	length of an array
// 	*/
// 	fmt.Println("\nlength of an array")
// 	i := [...]float64{67.7, 89.8, 21, 78}
// 	fmt.Println("length of a is", len(i))

// 	/*
// 	iterating an array using for loop
// 	*/
// 	fmt.Println("\niterating array using for loop")
// 	farray := [...]float64{67.7, 89.8, 21, 78}
// 	for i := 0; i < len(a); i++ {
// 		fmt.Printf("%d th element of a is %.2f\n", i, farray[i])
// 	}

// 	/*
// 	iterating an array using for range loop
// 	*/
// 	fmt.Println("\niterating array using for range loop")
// 	for i, v := range farray { //range returns both the index and value
// 		fmt.Printf("%d the element of a is %.2f\n", i, v)
// 	}

// 	/*
// 	2d arrays
// 	*/
// 	fmt.Println("\ndeclaring 2d arrays")
// 	animals := [3][2]string{
// 		{"lion", "tiger"},
// 		{"cat", "dog"},
// 		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
// 	}
// 	printarray(animals)
// 	var company [3][2]string
// 	company[0][0] = "apple"
// 	company[0][1] = "samsung"
// 	company[1][0] = "microsoft"
// 	company[1][1] = "google"
// 	company[2][0] = "AT&T"
// 	company[2][1] = "T-Mobile"
// 	fmt.Printf("\n")
// 	printarray(company)

// 	/*
// 	slice declaration
// 	*/
// 	fmt.Println("\nSlice declaration")
// 	as := [5]int{76, 77, 78, 79, 80}
// 	var slice1 []int = as[1:4] //creates a slice from a[1] to a[3]
// 	fmt.Println(slice1)

// 	slice2 := []int{6, 7, 8} //creates and array and returns a slice reference
// 	fmt.Println(slice2)

// 	/*
// 	modifying a slice, modifies the underlying array
// 	*/
// 	fmt.Println("\nmodifying a slice, modifies the underlying array")
// 	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
// 	dslice := darr[2:5]
// 	fmt.Println("array before", darr)
// 	for i := range dslice {
// 		dslice[i]++
// 	}
// 	fmt.Println("array after", darr)

// 	/*
// 	When a number of slices share the same underlying array,
// 	the changes that each one makes will be reflected in the array.
// 	*/
// 	fmt.Println("\nWhen slices share the same underlying array, changes each one makes is reflected in the array")
// 	numa := [3]int{78, 79, 80}
// 	nums1 := numa[:] //creates a slice which contains all elements of the array
// 	nums2 := numa[:]
// 	fmt.Println("array before change", numa)
// 	nums1[0] = 100
// 	fmt.Println("array after modification to slice nums1", numa)
// 	nums2[1] = 101
// 	fmt.Println("array after modification to slice nums2", numa)

// 	/*
// 	length and capacity of slice
// 	*/
// 	fmt.Println("\nlength and capacity of slice")
// 	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
// 	fruitslice := fruitarray[1:3]
// 	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6

// 	/*
// 	a slice can be re-sliced upto its capacity
// 	*/
// 	fmt.Println("\n\nreslicing a slice")
// 	vegarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
// 	vegslice := vegarray[1:3]
// 	fmt.Printf("length of slice %d capacity %d\n", len(vegslice), cap(vegslice)) //length of is 2 and capacity is 6
// 	vegslice = vegslice[:cap(fruitslice)]                                    //re-slicing furitslice till its capacity
// 	fmt.Println("After re-slicing length is", len(vegslice), "and capacity is", cap(vegslice))

// 	/*
// 	declaring a slice using make
// 	*/
// 	fmt.Println("\ndeclaring a slice using make")
// 	mkslice := make([]int, 5, 5)
// 	fmt.Println(mkslice)

// 	/*
// 	appending to slice
// 	*/
// 	fmt.Println("\nappending to a slice")
// 	cars := []string{"Ferrari", "Honda", "Ford"}
// 	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
// 	cars = append(cars, "Toyota")
// 	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

// 	/*
// 	appending to nil slice
// 	*/
// 	fmt.Println("\nappending to a nil slice")
// 	var names []string //zero value of a slice is nil
// 	if names == nil {
// 		fmt.Println("slice is nil going to append")
// 		names = append(names, "John", "Sebastian", "Vinay")
// 		fmt.Println("names contents:", names)
// 	}

// 	/*
// 	append one slice to another
// 	*/
// 	fmt.Println("\nappending one slice to another")
// 	veggies := []string{"potatoes", "tomatoes", "brinjal"}
// 	fruits := []string{"oranges", "apples"}
// 	food := append(veggies, fruits...)
// 	fmt.Println("food:", food)
	
// 	/*
// 	changes made to a slice inside a function are visible to the caller
// 	*/
// 	fmt.Println("\nchanges made to a slice inside a function are visible to the caller")
// 	nos := []int{8, 7, 6}
// 	fmt.Println("slice before function call", nos)
// 	subtactOne(nos)                               //function modifies the slice
// 	fmt.Println("slice after function call", nos) //modifications are visible out

// 	/*
// 	multidimensional slices
// 	*/
// 	fmt.Println("\nMultidimensional slices")
// 	pls := [][]string{
// 		{"C", "C++"},
// 		{"JavaScript"},
// 		{"Go", "Rust"},
// 	}
// 	for _, v1 := range pls {
// 		for _, v2 := range v1 {
// 			fmt.Printf("%s ", v2)
// 		}
// 		fmt.Printf("\n")
// 	}

// 	/*
// 		memory optimization using copy
// 	*/
// 	fmt.Println("\nMemory optimization using copy")
// 	countriesNeeded := countries()
// 	fmt.Println(countriesNeeded)

// }
