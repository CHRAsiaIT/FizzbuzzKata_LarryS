package main

import "testing"

//fizz tests
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

//fizz tests end

//buzz tests
func TestFizzbuzz_Buzz5( t *testing.T) {
	//Arrange
	expected := "Buzz"
	//Act
	result := fizzbuzz(5)
	//result
	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
		
	}
}

func TestFizzbuzz_Buzz10( t *testing.T) {
	//Arrange
	expected := "Buzz"
	//Act
	result := fizzbuzz(10)
	//result
	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}
//buzz tests end

//fizzbuzz tests
func TestFizzbuzz_FizzBuzz15( t *testing.T) {
	//Arrange
	expected := "FizzBuzz"
	//Act
	result := fizzbuzz(15)
	//result
	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
		
	}
}

func TestFizzbuzz_FizzBuzz45( t *testing.T) {
	//Arrange
	expected := "FizzBuzz"
	//Act
	result := fizzbuzz(45)
	//result
	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
		
	}
}
//fizzbuzz tests end