package main

import (
	"fmt"
)

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

//map function
func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}
func main() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	//Let's say we want to find all students from India. This can be done easily by changing the function parameter to the filter function
	f := filter(s, func(s student) bool {
		if s.grade == "B" { // if s.country == "India"
			return true
		}
		return false
	}) //filter returns slice so f become slice
	fmt.Println(f)

	//MAP FUNCTION
	//functions which operate on every element of a collection are called map functions
	am := []int{5, 6, 7, 8, 9}
	rm := iMap(am, func(n int) int {
		return n * 5
	}) //iMap returns slice so rm become slice
	fmt.Println(rm)
}
