package main
import(
	"fmt"
	"unicode/utf8"
)
func printBytes(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%c ",s[i])
    }
}

func mutate(s []rune) string {  
	s[0] = 'a'
	s = append(s,'h','g','f','d')
    return string(s)
}
func main(){
	var personSalary map[string]int //here personSalary is nil
	personSalary = make(map[string]int) // {"steve":2300, "jamie":4239} at the time of initialization
	fmt.Println("hoya")
	personSalary["steve"] = 12000
    personSalary["jamie"] = 15000
    personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)
	//retrieval
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])
	fmt.Println("Salary of joe is", personSalary["joe"]) //ZERo value of TYPE INT is returned
	//In reality two values returned from map[key], where first one is value second is present bit
	newEmp:="joe"
	value, okornot := personSalary[newEmp]
	if okornot == true {
        fmt.Println("Salary of", newEmp, "is", value)
    } else {
        fmt.Println(newEmp,"not found")
	}
	//range for loop for iteration over map
	//ORDER IS NOT FIXED NOT NOT NOT NOT NOT NOT FIXED
	fmt.Println("All items of a map")
    for key, value := range personSalary {
        fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
	//Deletion of the element by key
	delete(personSalary, "steve")
	fmt.Println("length of the map is ",len(personSalary),"map after deletion", personSalary)
	/*
	Similar to slices, maps are reference types. When a map is assigned to a new variable, they both point to the same internal data structure. Hence changes made in one will reflect in the other.
	*/
	newPerson := personSalary
	fmt.Println(newPerson)
	// here both pointing to the same map
	// map cant be  compared with ==
	// The == can be only used to check if a map is nil.
	// For comparing you must have to compare one by one all the elements

	//**********STRING****************

	name := "Hello World" //strings are stored bytewise
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
    fmt.Printf("\n")
    name = "Señor" // here ~n occupied two bytes in UTF8 so it displayed incorrectly babes
    printBytes(name)
    fmt.Printf("\n")
	printChars(name)
	//Rune saves us here
	/*
	A rune is a builtin type in Go and it's the alias of int32. rune represents a Unicode code point in Go. It does not matter how many bytes the code point occupies, it can be represented by a rune
	*/
	charss := []rune(name) //perfect container for any character in Go
    for i:= 0; i < len(charss); i++ {
        fmt.Printf("%c ",charss[i])
	}
	//range is everywhere :)
	for index, charss := range name {
        fmt.Printf("%c starts at byte %d\n", charss, index)
	}
	//Constructing string from slice of bytes
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9} // {67, 97, 102, 195, 169} also works same
    str := string(byteSlice)
	fmt.Println(str)
	//From rune
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str2 := string(runeSlice)
	fmt.Println(str2)
	
	//Length of the string
	fmt.Printf("length of %s is %d\n", str2, utf8.RuneCountInString(str2))
	fmt.Printf("length of %s is %d\n", "Pets", utf8.RuneCountInString("Pets"))
	fmt.Printf("length of %s is %d\n", "jaydeep", len("jaydeep")) //for simplicity len would be suffice
	//Strings are immutable in Go. Once a string is created it's not possible to change it
	// str2[2]='a' throws error
	//Workaround is to convert it to slice of rune
	h := "hello"
    fmt.Println(mutate([]rune(h)))
}