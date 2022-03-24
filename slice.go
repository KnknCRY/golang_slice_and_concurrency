package main

import "fmt"

// golang的append，底層實現概念
// reference: https://go.dev/blog/slices
func Append(slice []int, elements ...int) []int {
	n := len(slice)
	total := len(slice) + len(elements)
	if total > cap(slice) {
		// Reallocate. Grow to 1.5 times the new size, so we can still grow.
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice // @sliceDemo5()
	}
	slice = slice[:total] // @sliceDemo4()
	copy(slice[n:], elements)
	return slice
}

// 最難，用到append的底層概念
func sliceDemo1() {
	// case 1
	a := make([]int, 0, 10)

	// case2
	// a := make([]int, 0, 2)

	// fmt.Println(&a[0])
	// a = a[0:5]
	// a = a[3:5]
	// fmt.Println(&a[0])

	b := append(a, 1, 2, 3)
	c := append(a, 4, 5, 6)
	fmt.Println(a, b, c)
	fmt.Println(len(a), cap(a), &a)
	fmt.Println(len(b), cap(b), &b)
	fmt.Println(len(c), cap(c), &c)
}

func passForDemo2(arr []int) {
	// arr跟 a是不同的slice，但指向同一個underlying array
	// 並非arr指向 a這個slice

	// case 1
	// arr跟 a指向同一個underlying array
	// arr[0] = 4

	// case 2
	// arr指向一個全新的underlying array
	arr = []int{4, 5, 6}
}

// slice參數傳遞概念
func sliceDemo2() {
	a := []int{1, 2, 3}
	passForDemo2(a)
	fmt.Println(a)
}

// slice賦值概念
func sliceDemo3() {
	// a b是不同slice但指向同一個underlying array
	a := []int{1, 2, 3}
	b := a
	b[0] = 0
	fmt.Println(a, b)

	// output:
	// [0 2 3] [0 2 3]
}

// [:]slice賦值概念
func sliceDemo4() {
	// a b是不同slice但指向同一個underlying array
	a := []int{1, 2, 3} // create a new slice and point to a new underlying array
	b := a[1:2]         // create a new slice but point to the underlying array of a
	a[1] = 8
	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(b), cap(b), b)

	// output:
	// 3 3 [1 0 3]
	// 1 2 [0]
}

// slice賦值概念
func sliceDemo5() {
	// 直接改變a的len跟cap，也改變了underlying array
	a := []int{1, 2, 3}
	b := make([]int, 5, 5)

	a = b
	b[2] = 8

	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(b), cap(b), b)

	// output:
	// 5 5 [0 0 8 0 0]
	// 5 5 [0 0 8 0 0]
}

// slice伸縮概念，但沒事不要用伸縮，用append就好
func sliceDemo6() {
	a := make([]int, 3, 5)
	a[2] = 8
	fmt.Println(len(a), cap(a), a)

	a = a[:len(a)+1] // 對a的len做伸縮，並指向同一個underlying array
	fmt.Println(len(a), cap(a), a)

	// output:
	// 3 5 [0 0 8]
	// 4 5 [0 0 8 0]
}
