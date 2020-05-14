package main

import (
	"5_pointer_method_struct/structs/computer"
	"fmt"
)

func main() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	//spec.model = "dell" --> gives error bcs it's not an exported field
	fmt.Println("Spec:", spec)
}
