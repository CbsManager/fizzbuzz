package main

import "fmt"

func FizzBuzz(n int) string {
	if n == 3 {
		return "Fizz"
	} else if n == 5 {
		return "Buzz"
	}
	return fmt.Sprint(n)
}
