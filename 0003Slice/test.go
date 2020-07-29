package main

import "fmt"

/*
 *Slice structure:    var slice1 []type = make([]type, len)
 *					make([]T, length, capacity)//capacity is not necessary
 */

func main() {

	var numbers = make([]int, 3, 5)
	printSlice(numbers)
	// Before initial the slice, it is nil
	var nums []int
	if nums == nil {
		fmt.Println("切片是空的")
	}
	printSlice(nums)
	nums = make([]int, 3, 10)
	if nums == nil {
		fmt.Println("切片是空的")
	}

	//cut the slice with [lower-bound:upper-bound]
	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers2)
	fmt.Println("numbers2 ==", numbers2)
	fmt.Println("numbers2[1:4] ==", numbers2[1:4]) //print m to n (not include n)
	fmt.Println("numbers2[:3] ==", numbers2[:3])   // print form 0 to n (not include n)
	fmt.Println("numbers2[4:] ==", numbers2[4:])   // print form n to end (not include n)

	//add new element to slice with append
	numbers2 = append(numbers2, 9, 10, 11)
	printSlice(numbers2)

	//copy the slice
	numbersCopy := make([]int, len(numbers2), (cap(numbers2))*2)
	copy(numbersCopy, numbers2)
	printSlice(numbers2)
	printSlice(numbersCopy)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
