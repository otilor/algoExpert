package main

import (
	. "fmt"
	"unsafe"
)


func main() {
	//swapOddValues([]int  {12, 3, 323, 31,67, 70, 45})
	realSwap := realSwap([]int {12, 3, 323, 31, 67, 70, 45, 2, 1, 12})
	Println(realSwap)
}

func realSwap(arr []int) []int{
	size := len(arr)

	swappable := make(map[int]int)

	var new []int

	for i:= 0; i < size; i = i + 4{
		swappable[i] = arr[i]
	}
	for j := 0; j < size; j++ {
		if arr[j] == arr[size-1]{
			new = append(new, arr[size-1])
			continue
		}
		if arr[j] == swappable[j]{
			a, b := swap(arr[j], arr[j+1])
			new = append(new, a, b)
			j+=1
			continue
		}
		new = append(new, arr[j])
	}
	return new
}

func funcQuad(list []int){
	for i :=1; i < len(list); i++{
		for j:= 1 ; j <= len(list); j++ {
			Println(i, j)
		}
	}
}


func sumOfNumbers(n int) int{

	return n * (n - 1) / 2
}

func printAlgoExpert(n int){
	for i := 0; i < n; i++ {
		Println("algaExpert!")
	}
}
// Space complexity is O(1)


// Array sequences
// Array of 6 characters
// Each character takes up a  CELL
// INDEX describes location


func loopAndShowSize(n  int){
	var data []int
	for i:= 0; i < n; i++ {
		a := len(data)
		b := unsafe.Sizeof(data)
		Printf("Length: %d; Size in bytes: %d", a, b)
		Println()
		data = append(data, n)
	}

}

func reverseString(s string) {
	for i:=len(s)-1; i >= 0; i--{
		Println(string(s[i]))
	}
}


func reverseArray(arr []int) {
	for i:= len(arr)-1; i >= 0; i--{
		Println(arr[i])
	}
}

func reverseWhenMiddle(arr []int) {
	size := len(arr)
	middle := int (size / 2)
	otherHalf := arr[middle:size]
	for i:= 0; i < size; i++ {
		if i == middle {
			reverseArray(otherHalf)
			break
		}

		Println(arr[i])
	}
}

func swapOddValues(arr []int) {
	// loop through the array
	// For example,
	// {1, 2, 3, 4} should give

	// {2, 1, 3, 4}

	// {1, 2, 3, 4, 5, 6} becomes
	// {2, 1, 3, 4, 6, 5}

}


func swapWhenEven(values []int){
	var swappable []int
	for i := 0; i < len(values); i = i + 4{
		swappable = append(swappable, i)
	}
	Println(swappable)
	for j := 0; j < len(values); j++{

		if isset(values, j) && isset(swappable, j){
			Println("Hey")
			swap(values[j], values[j+1])
			j += 1

		}else {
			Println("Others")
			Println(j+1)
		}
	}
}

func isset(arr []int, index int) bool{
	return len(arr) > index
}



func swap (a int, b int) (int, int){
	temp := a
	a = b
	b = temp
	return a, b
}