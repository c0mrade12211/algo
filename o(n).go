// This is a simple example demonstrating the O(n) time complexity. 
// The program's time and memory usage depend on the input values. 

package main 
 
import "fmt" 
 
func linearSearch(arr []int, target int) bool { 
    for _, num := range arr { 
        if num == target { 
            return true 
        } 
    } 
    return false 
} 
 
func main() { 
    numbers := []int{1, 2, 3, 4, 5, 6} 
    target := 4 
    found := linearSearch(numbers, target) 
    fmt.Println("Element found:", found) 
}
//next o(log n), O(n!)
