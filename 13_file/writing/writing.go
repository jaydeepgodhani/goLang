package main

import (
	"fmt"
	"os"
)

func main() {
	//If a file with that name already exists, then the create function truncates the file
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Hello World")
	//This method returns the number of bytes written and error if any
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//WRITING BYTES
	fbyte, err := os.Create("bytes")
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 108, 101}
	n2, _ := fbyte.Write(d2)
	// if err != nil {
	//     fmt.Println(err)
	//     f.Close()
	//     return
	// }
	fmt.Println(n2, "bytes written successfully")
	err = fbyte.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//WRITING LINE BY LINE
	fline, _ := os.Create("lines")
	d3 := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}
	//we iterate through the array using a for range loop and use the Fprintln function to write the lines to a file. The Fprintln function takes a io.writer as parameter and appends a new line
	for _, v := range d3 {
		fmt.Fprintln(fline, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = fline.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//APPEND
	//The file has to be opened in append and write only mode. These flags are passed parameters are passed to the Open function
	fappend, _ := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0644)
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(fappend, newLine)
	if err != nil {
		fmt.Println(err)
		fappend.Close()
		return
	}
	err = fappend.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
