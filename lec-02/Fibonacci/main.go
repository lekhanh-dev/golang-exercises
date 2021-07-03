package main

import "fmt"

func Fibonacci(n int) {
	f1 := 1
	f2 := 1
	f3 := 0
	i := 3
	fmt.Print(f1)
	fmt.Print(" ")
	fmt.Print(f2)
	fmt.Print(" ")
	for i <= n {
		f3 = f1 + f2
		fmt.Print(f3)
		fmt.Print(" ")
		f1 = f2
		f2 = f3
		i++
	}
}
func main() {
	Fibonacci(7)
}
