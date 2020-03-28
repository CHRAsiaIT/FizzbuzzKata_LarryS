package main

import "testing"

//test fizzbuzz function
func TestFizzbuzz_Fizz3( t *testing.T) {
		//Arrange
		expected := "Fizz"
		//Act
		result := fizzbuzz(3)
		//result
		if result != expected {
			t.Errorf("expected %q but got %q", expected, result)
		}
}

func TestFizzbuzz_Fizz6( t *testing.T) {
	//Arrange
	expected := "Fizz"
	//Act
	result := fizzbuzz(6)
	//result
	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}