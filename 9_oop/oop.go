package main

import (
	"9_oop/employee"
	"fmt"
)

//Composition
type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName()) // originally it's p.author.fullName(), but due to embedded struct in other struct we can directly call it as if they were part of it post itself
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

//polymorphism
type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}
func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

func main() {
	e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
	//when we define the employee struct with zero values it displays all zero values
	//Java has constructer for this type of situation
	e2 := employee.New("jaydeep", "Adolf", 70, 20) //this is the only way to create a employee now if we delete all Employee code from employee.go package file
	e2.LeavesRemaining()
	//Go does not support inheritance, however it does support composition. The generic definition of composition is "put together". One example of composition is a car. A car is composed of wheels, engine and various other parts.
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	post1.details()
	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()
	//Polymorphism in Go is achieved with the help of interfaces
	//A variable of type interface can hold any value which implements the interface. This property of interfaces is used to achieve polymorphism in Go
	//POLYMORPHISM
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}
