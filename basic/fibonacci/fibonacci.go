package main

import (
	"fmt"
)

func calculateFib(n int) ([]int, error) {
	// Fibonacci series: 0,1,1,2,3,5,8...
	s := make([]int, n)

	if n == 0 {
		return s, nil
	}

	s[0] = 0
	if n == 1 {
		return s, nil
	}

	s[1] = 1
	if n == 2 {
		return s, nil
	}

	for i := 3; i <= n; i++ {
		s[i-1] = s[i-2] + s[i-3]
	}

	return s, nil
}

func main() {
	var input int
	fmt.Print("Enter the count of numbers in a fibonacci series to print: ")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Printf("\nCouldn't get the input from the user. Error: %s", err.Error())
	}

	if output, err := calculateFib(input); err != nil {
		fmt.Printf("\nCouldn't calculate the fibonacci series. Error: %s", err.Error())
	} else {
		fmt.Println(output)
	}
}
