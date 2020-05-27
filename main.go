package main

import (
	. "fmt"
	"strings"
	"unsafe"
)

func main() {
	stringCompression("aabcccccaaa")
}

// Take the string in, iterate through it,
// store each unique elements in a hashmap,
// Iterate through the hashmap and print the elements with their respective counts


func stringCompression(s string) {
	hashMap := make(map[string]int)
	for i := 0; i < len(s); i++{
		hashMap[string(s[i])] += 1
	}

	Println(hashMap)
}

func urlIfy(s string, l int) string{
	newString := ""
	for i := 0; i < l; i++ {
		if string(s[i]) == " " {
			s2 := newString + "%20"
			newString = s2
			continue
		}
		s2 := newString + string(s[i])
		newString = s2
	}
	return newString
}
// Iterate through the first string,
// Store each character in a hashmap,
// Iterate through the second string,
// If an element has 0,
// Increment by 1
func oneOrZeroEditAway(s1 string, s2 string) bool{
	var count int
	hashMap1 := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		hashMap1[string(s1[i])] += 1
	}

	for i := 0; i < len(s2); i++ {
		if hashMap1[string(s2[i])] == 0 {
			count+= 1
		}
	}
	if count > 1 {
		return false

	} else {
		return true
	}

}
func permutationOfAnother(s1 string, s2 string) bool {
	// Get the first string and put all the elements in a hash map
	// Iterate through the second string and see if it has an element in the first
	hashMap1 := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		hashMap1[string(s1[i])] += 1
	}

	for j := 0; j < len(s2); j++ {
		if hashMap1[string(s2[j])] == 0 {
			return false
			break
		}
	}

	return true
}

func isUnique(s string) bool {

	s = strings.ToLower(s)
	size := len(s)
	elements := make(map[string]int)

	for i := 0; i < size; i++ {
		elements[string(s[i])] += 1
	}

	for i := 0; i < size; i++ {
		if elements[string(s[i])] > 1 {
			return false
			break
		}
	}

	return true
}

// Input : [1, 1, 2, 3, 3, 4, 4, 8, 8]
// Output

// Iterate through it, get the

func singleElement(nums []int) {

	// binary search
	low, high := 0, len(nums)-1
	for i := low; i < high-1; i++ {
		mid := (low + high) / 2
		isMidOdd := mid%2 == 0
		areAdjNumbersMatching := nums[mid] == nums[mid+1]
		if isMidOdd && !areAdjNumbersMatching {
			low = mid
		} else if isMidOdd && areAdjNumbersMatching {
			high = mid
		} else if !isMidOdd && !areAdjNumbersMatching {
			high = mid
		} else if isMidOdd && areAdjNumbersMatching {
			low = mid
		}

	}
	if low%2 == 1 {
		Println(nums[high])
	} else if low%2 == 0 {
		Println(nums[low])
	}
}

func myShuffle(arr string) string {
	//ar := []rune(arr)
	// iterate through the letters,
	// for each letter,
	// print a number

	num := make(map[int]string)
	letters := make(map[int]string)
	for i := 0; i < len(arr)/2; i++ {
		num[i] = string(arr[i])
	}
	for j := len(arr) / 2; j <= len(arr)-1; j++ {
		letters[j] = string(arr[j])
	}

	var s string
	var a string

	for j := range num {
		s += num[j]
		a += letters[j+len(arr)/2]
	}
	Println(s)
	Println(a)
	return arr

}
func shuffle(arr string) string {
	ar := []rune(arr)
	n := len(ar) / 2
	count := 0
	k := 1
	var temp rune
	for i := 1; i < n; i = i + 2 {
		k = i
		temp = ar[i]
		for true {
			k = (2 * k) % (2*n - 1)
			temp, ar[k] = ar[k], temp
			count++
			if i == k {
				break
			}
		}
		if count == (2*n - 2) {
			break
		}
	}
	return string(ar)
}

// This works if the elements are different. i.e. there are no repeating values

func realSwap(arr []int) []int {
	size := len(arr)

	swappable := make(map[int]int)

	var new []int

	for i := 0; i < size; i = i + 4 {
		swappable[i] = arr[i]
	}
	for j := 0; j < size; j++ {
		if arr[j] == arr[size-1] {
			new = append(new, arr[size-1])
			continue
		}
		if arr[j] == swappable[j] {
			a, b := swap(arr[j], arr[j+1])
			new = append(new, a)
			new = append(new, b)
			j += 1
			continue
		}
		new = append(new, arr[j])
	}
	Println(swappable)
	return new
}

func funcQuad(list []int) {
	for i := 1; i < len(list); i++ {
		for j := 1; j <= len(list); j++ {
			Println(i, j)
		}
	}
}

func sumOfNumbers(n int) int {

	return n * (n - 1) / 2
}

func printAlgoExpert(n int) {
	for i := 0; i < n; i++ {
		Println("algaExpert!")
	}
}

// Space complexity is O(1)

// Array sequences
// Array of 6 characters
// Each character takes up a  CELL
// INDEX describes location

func loopAndShowSize(n int) {
	var data []int
	for i := 0; i < n; i++ {
		a := len(data)
		b := unsafe.Sizeof(data)
		Printf("Length: %d; Size in bytes: %d", a, b)
		Println()
		data = append(data, n)
	}

}

func reverseString(s string) {
	for i := len(s) - 1; i >= 0; i-- {
		Println(string(s[i]))
	}
}

func reverseArray(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		Println(arr[i])
	}
}

func reverseWhenMiddle(arr []int) {
	size := len(arr)
	middle := int(size / 2)
	otherHalf := arr[middle:size]
	for i := 0; i < size; i++ {
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

func swapWhenEven(values []int) {
	var swappable []int
	for i := 0; i < len(values); i = i + 4 {
		swappable = append(swappable, i)
	}
	Println(swappable)
	for j := 0; j < len(values); j++ {

		if isset(values, j) && isset(swappable, j) {
			Println("Hey")
			swap(values[j], values[j+1])
			j += 1

		} else {
			Println("Others")
			Println(j + 1)
		}
	}
}

func isset(arr []int, index int) bool {
	return len(arr) > index
}

func swap(a int, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}
