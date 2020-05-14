package main

import (
	"fmt"
	"reflect"
)

/*
The reflect package implements run-time reflection in Go. The reflect package helps to identify the underlying concrete type and the value of a interface{} variable. This is exactly what we need. The createQuery function takes a interface{} argument and the query needs to be created based on the concrete type and value of the interface{} argument. This is exactly what the reflect package helps in doing
*/
type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	//Type represents the actual type of the interface{}, in this case main.Order and Kind represents the specific kind of the type. In this case, it's a struct.
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)

}
func createQuerydetail(q interface{}) {
	if reflect.TypeOf(q).Kind() == reflect.Struct { //reflect.ValueOf(q).Kind() also works here
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField()) //returns Number of fields
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)
	createQuerydetail(o)

	//Int and String methods
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)
}
