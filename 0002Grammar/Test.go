package main

import "fmt"

var ( //used to define the global variables
	a int
	b bool
)

func main() {

	fmt.Println("Test " + "Grammar.")

	//Defining the variables
	var temp1, temp2 int
	temp1 = 1
	var temp3 = "hello"
	temp4 := false //left side shouldn't have defined variable,can only used in func

	fmt.Println(temp1, temp2, temp3, temp4, a, b)
	fmt.Println(&temp1)

	temp2, _ = temp1, temp2 //drop the temp2 0
	fmt.Println(temp1, temp2, temp3, temp4, a, b)

}
