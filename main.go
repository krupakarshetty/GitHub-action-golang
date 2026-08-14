package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello from GitHub Actions CI/CD!")
	fmt.Println(Add(2, 3))
}
