package employee

import (
	"fmt"
)

type Employee struct { //can accessed from other package because capitalized word
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}
type employee struct { //prevent access from other package because uncapitalized word
	firstName   string //all fields are also unexported
	lastName    string
	totalLeaves int
	leavesTaken int
}

//this is function which is Exported bcs of capitalized first letter
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}
func (e Employee) LeavesRemaining() { //this is a method (not a function)
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
func (e employee) LeavesRemaining() { //this is a method (not a function)
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
