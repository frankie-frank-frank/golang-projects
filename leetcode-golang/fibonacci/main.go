package main

import (
    //"fmt" // Use fmt.Println(...) to debug your code
// Fix me!
)

func Factorial(n int) int {
    if n == 0 || n == 1 { return 1 }
    factorialNum := n * Factorial(n-1)
	return factorialNum
}