package main

import "fmt"

const PI = 3.14

func main() {
	test6()
}

func test6() {
	mySlice1 := []int{}
	mySlice1 = append(mySlice1, 1)
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
	fmt.Println(mySlice1)

	mySlice2 := []string{"Go", "Slices", "Are", "Powerful"}
	mySlice2 = append(mySlice2, "OK")
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
	fmt.Println(mySlice2)
}

func test5() {
	prices := [3]int{10, 20, 30}
	prices[2] = 50
	fmt.Println(prices)
}

func test4() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}

func test3() {
	var i, j string = "Hello", "World"
	fmt.Print(i, " ", j)
}

func test0() {
	var student1 string = "John"
	var student2 = "Jane"
	x := 2 // type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	// this is a comment
	fmt.Println("Hello World!")
}

func test1() {
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func test2() {
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
