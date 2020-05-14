package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/gobuffalo/packr"
)

/*
Three ways

    Using absolute file path
    Passing the file path as a command line flag
    Bundling the text file along with the binary

*/
func main() {
	data, err := ioutil.ReadFile("test.txt")
	//absolute path for this file is  -- /home/jordan/go/src/13_file/test.txt
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Printf("type of data is %T\n", data)
	fmt.Println("Contents of file:", string(data))
	//The flag package has a String function. This function accepts 3 arguments. The first is the name of the flag, second is the default value and the third is a short description of the flag
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()                            //flag.Parse() should be called before any flag is accessed by the program.
	fmt.Println("value of fpath is", *fptr) //without command line arg default value is test.txt
	//if uses -- go install 13_file , then from the ../bin folder we can run 13_file with command line args -fpath=/home/jordan/go/src/13_file
	box := packr.NewBox("/13_file")
	data2 := box.String("test.txt") //this command read latest modified file from the location
	fmt.Println("Contents of file:", data2)
	//packr install -v filehandling   --> this command pack text file into binary
	//reaminings are - read line by line, read in chunks, scanner, reader
}
