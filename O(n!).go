//It's a simple factorial O(n!) complexity

package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	fmt.Println(n)
	return n * factorial(n-1)
}

func main() {
	number := 4
	result := factorial(number)
	fmt.Printf("Факториал числа %d равен %d\n", number, result)
}
