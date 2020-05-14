package main
/*
Sometimes we need to import a package just to make sure the initialisation takes place even though we do not need to use any function or variable from the package
we might need to ensure that the init function of the rectangle package is called even though we do not use that package anywhere in our code
*/
import(
	"fmt"
	"geomatry/rectangle"
	//	_"geomatry/rectangle" --> by using this we can initialize rectangle package and we should remove method calling of retanlge package.
	"log"
	"math"
)
/*
    The imported packages are first initialised. Hence rectangle package is initialised first.
    Package level variables rectLen and rectWidth are initialised next.
    init function is called.
    main function is called at last
*/

var _ = math.Min //error silencer, if commented out then error
//it is recommended to write error silencers in the package level just after the import statement

var rectLen, rectWidth float64 = -6, 7

func init() {  
    println("main package initialized")
    if rectLen < 0 {
        log.Fatal("length is less than zero")
    }
    if rectWidth < 0 {
        log.Fatal("width is less than zero")
    }
}

func main(){
    fmt.Println("Geometrical shape properties")
        /*Area function of rectangle package used
        */
    fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
        /*Diagonal function of rectangle package used
        */
    fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))
}