package manongo

import "testing"

//fizzbuzz should return fizz if evenly divisible by 3
func TestFizzBuzz_returnsFizz(t *testing.T) {
	numbers := []int64{3, 6, 9, 12, 18, 21, 24, 27, 33, 36, 39}
	for _, num := range numbers {
		s := FizzBuzz(num)
		if s != "fizz" {
			t.Fail()
		}
	}

}

//fizzbuzz should return buzz if evenly divisible by 5
func TestFizzBuzz_returnsBuzz(t *testing.T) {
	numbers := []int64{5, 10, 20, 25, 35, 40, 50, 55, 65}
	for _, num := range numbers {
		s := FizzBuzz(num)
		if s != "buzz" {
			t.Fail()
		}
	}

}

//fizzBuzz should return FizzBuzz if evenly divisible by both
func TestFizzBuzz_returnsFizzBuzz(t *testing.T) {
	numbers := []int64{15, 30, 45, 60, 75, 90, 105, 120, 135}
	for _, num := range numbers {
		s := FizzBuzz(num)
		if s != "fizzbuzz" {
			t.Fail()
		}
	}
}

//fizzbuzz should return an empty string when a number is neither evenly divisible
//by 3 or 5
func TestFizzBuzz_returnsEmptyString(t *testing.T) {
	numbers := []int64{1, 2, 4, 7, 8, 101, 202, 304}
	for _, num := range numbers {
		s := FizzBuzz(num)
		if s != "" {
			t.Fail()
		}
	}

}
