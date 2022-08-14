package main

// Example of a recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}
func main() {
	println(factorial(5))
}
