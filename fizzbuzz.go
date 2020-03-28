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
	result := ""
	if i%3 == 0 {result += "Fizz"}
	if i%5 == 0 {result += "Buzz"}
	if result =="" {result = strconv.Itoa(i)}
	fmt.Print(i)
		return result 
} 