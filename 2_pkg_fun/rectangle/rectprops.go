package rectangle

import "math"
import "fmt"

/*Capitalize first letter of variable and function name
It means Exported Names
Only exported functions and variables can be accessed from other packages
Every package can contain init() function with no parameters and no return type
The init function cannot be called explicitly in our source code

1.  Package level variables are initialised first
2.  init function is called next. A package can have multiple init functions (either in a single 	file or distributed across multiple files) and they are called in the order in which they are 	presented to the compiler.

*/

func init() {  
    fmt.Println("rectangle package initialized")
}

func Area(len, wid float64) float64 {  
    area := len * wid
    return area
}

func Diagonal(len, wid float64) float64 {  
    diagonal := math.Sqrt((len * len) + (wid * wid))
    return diagonal
}