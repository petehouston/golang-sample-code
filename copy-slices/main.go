package main

import "fmt"

func srcSizeIsLessThanDestSize() {
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, 5)

	count := copy(dest, src)

	fmt.Println("Source size is less than Destination size")
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dest: %v\n", dest)
	fmt.Printf("Number of copied elements: %d\n", count)
}

func srcSizeIsEqualToDestSize() {
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, 10)

	count := copy(dest, src)

	fmt.Println("Source size is equal to Destination size")
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dest: %v\n", dest)
	fmt.Printf("Number of copied elements: %d\n", count)
}

func srcSizeIsGreaterToDestSize() {
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, 2)

	count := copy(dest, src)

	fmt.Println("Source size is greater than Destination size")
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dest: %v\n", dest)
	fmt.Printf("Number of copied elements: %d\n", count)
}

func checkCopy() {
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, 5)

	count := copy(dest, src)

	fmt.Printf("src: %v\n", src)
	fmt.Printf("dest: %v\n", dest)
	fmt.Printf("Number of copied elements: %d\n", count)
	fmt.Println("Update dest elements")
	dest[0] = 100
	dest[1] = 200
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dest: %v\n", dest)
}

func main() {
	srcSizeIsEqualToDestSize()
	srcSizeIsLessThanDestSize()
	srcSizeIsGreaterToDestSize()
	checkCopy()
}
