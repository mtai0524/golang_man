package main

import "fmt"

func main() {
	fmt.Print("Hello, World!\n")
	fmt.Print(sum(1, 2))
	fmt.Print("\n")
	fmt.Print(sumSlice([]int{1, 2, 3, 4, 5}))
}

// Hàm tính tổng 2 số nguyên
func sum(a, b int) int {
	return a + b
}

// Hàm tính tổng một mảng (slice) các số nguyên
func sumSlice(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
