package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = f(slice[i])
	}
}

func mapArray(f func(a int) int, array [3]int) {
	for i := 0; i < len(array); i++ {
		array[i] = f(array[i])
	}
}


func main() {
	intsSlice :=[]int{1,2,3}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice) // prints [2,3,4]

	intsArray:=[3]int{1,2,3}
	mapArray(addOne, intsArray)
	fmt.Println(intsArray) // prints [1,2,3]
}
