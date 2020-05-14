package computer

type Spec struct { //exported struct bcz of Capitalize first letter
	Maker string //exported field
	model string //unexported field
	Price int    //exported field
}
