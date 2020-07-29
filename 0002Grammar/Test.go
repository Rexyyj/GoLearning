package main

import (
	"fmt"
	"math"
)

var ( //used to define the global variables
	a int
	b bool
)

func main() {

	fmt.Println("Test " + "Grammar.")
	/*
	 *Defining the variables
	 */
	var temp1, temp2 int
	temp1 = 1
	var temp3 = "hello"
	temp4 := false //left side shouldn't have defined variable,can only used in func

	fmt.Println(temp1, temp2, temp3, temp4, a, b)
	fmt.Println(&temp1)

	temp2, _ = temp1, temp2 //drop the temp2 0
	fmt.Println(temp1, temp2, temp3, temp4, a, b)
	//----------------------------------------------------------------------------------------------------------
	/*
	 *Defining the constants
	 */
	const c1, c2 = 10, 11
	fmt.Println(c1, c2)

	const (
		c3 = iota
		c4 = iota
		c5 = iota //iota increase one everytime it is called
	)
	fmt.Println(c3, c4, c5)
	//-----------------------------------------------------------------------------------------------------------------
	/*
	 *Special switch
	 * No need for break in go
	 */

	//type switch
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
	fmt.Println()

	//fallfhtough
	//directly run the next case no matter the condition is false or true
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
	//----------------------------------------------------------------------------------
	/*
	 *Select
	 * each case must be a communication
	 * if multiple cases are satisfied, only one will be runned
	 * if no case communication and no default, select will be block until one case communication apperas
	 */

	var c6, c7 chan int
	var i1, i2 int
	select {
	case i1 = <-c6:
		fmt.Printf("received ", i1, " from c1\n")
	case c7 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	default:
		fmt.Printf("no communication\n")
	}
	//------------------------------------------------------------------------------------
	/*
	 * Use range in for loop
	 *
	 */
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println("Adding index", i)
	}
	fmt.Println("sum:", sum)

	//------------------------------------------------------------------------------------
	/*
	 *Functions
	 * Function can return multiple values like python
	 */
	fmt.Println(max(c1, c2))

	//Another type of defining the function
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	//------------------------------------------------------------------------------------
	/*
	 *Structure
	 *
	 */
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookID: 6495407})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"}) //ignore the var of zero initial value

	var Book1 = Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	fmt.Println(Book1.bookID)
}

func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}
