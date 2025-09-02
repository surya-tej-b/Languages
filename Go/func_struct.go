package main

import "fmt"

func main() {
	var printValue string = "Hello from main!"
	printMe(printValue)

	var a int = 10
	var b int = 5
	var result, remainder int = intDivision(a, b)
	fmt.Printf("Result: %d, Remainder: %d\n", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(a int, b int) (int, int) {
	var err error
	if b == 0 {
		err = errors.New("division by zero")
		return 0, 0, err
	}
	var result int = a / b
	var remainder int = a % b
	return result, remainder
}
