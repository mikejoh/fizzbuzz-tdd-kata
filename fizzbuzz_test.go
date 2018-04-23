package fizzbuzz

/*
Fizzbuzz TDD kata.

Return fizz 	= divisible by 3
Return buzz 	= divisible by 5
Return fizzbuzz = Both divisible by 3 and 5
*/

import (
	"testing"
)

/*
Helper functions
*/
func isEquals(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("%s is not %s", actual, expected)
	}
}

/*
Tests
*/
func TestNormalCases(t *testing.T) {
	isEquals(t, FizzBuzz(1), "1")
	isEquals(t, FizzBuzz(2), "2")
	isEquals(t, FizzBuzz(4), "4")
}

func TestThatItReturnsFizzWhenDevicibleBy3(t *testing.T) {
	isEquals(t, FizzBuzz(3), "fizz")
}

func TestThatItReturnsBuzzWhenDevicibleBy5(t *testing.T) {
	isEquals(t, FizzBuzz(5), "buzz")
}

func TestThatItReturnsFizzbuzzWhenDevicibleBy15(t *testing.T) {
	isEquals(t, FizzBuzz(15), "fizzbuzz")
}
