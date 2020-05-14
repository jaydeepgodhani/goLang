package main

import (
	"fmt"
	"math"
	"os"
)

/*
Errors in Go are plain old values. Errors are represented using the built-in error type.
Just like any other built in type such as int, float64, ... error values can be stored in variables, returned from functions and so on
*/
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		////This function formats the error according to a format specifier and returns a string as value that satisfies error
		//return 0, errors.New("Area calculation failed, radius is less than zero")
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}
func main() {
	//If a function or method returns an error, then by convention it has to be the last value returned from the function. Hence the Open function returns err as the last value

	//The idiomatic way of handling error in Go is to compare the returned error to nil. A nil value indicates that no error has occurred and a non nil value indicates the presence of an error

	/*
		error is an interface type with the following definition
		type error interface {
		    Error() string
		}
		It contains a single method with signature Error() string. Any type which implements this interface can be used as an error
	*/
	//original code below is to use "f" in place of _, and uncomment return and Println
	fmt.Println("ERROR HANDLING")
	_, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		//return
	}
	//fmt.Println(f.Name(),"opened successfully")
	fmt.Println("opened successfully")

	//CUSTOM ERROR
	//The simplest way to create a custom error is to use the New function of the errors package
	//The New function takes a string parameter, creates a value of type errorString using that parameter and returns the address of it. Thus a new error is created and returned
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}
