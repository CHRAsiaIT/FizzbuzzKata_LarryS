package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Please Enter an Integer Number: ")
	var input int
	fmt.Scanf("%d", &input)

	fizzbuzz(input)
}

func fizzbuzz(i int) string {
	if i % 15 == 0{
		fmt.Print("FizzBuzz")
		return "FizzBuzz"
	} else if i % 3 == 0 {
		fmt.Print("Fizz")
		return "Fizz"
	} else if i % 5 == 0 {
		fmt.Print("Buzz")
		return "Buzz"
	} else {
		fmt.Print(i)
		return strconv.Itoa(i)
	}
		
} 