//Contains code used in demostration and implements a
//Fizzbuzz solution in go
package main

// Returns fizz if the number is evenly divisible by 3
//
// buzz if the number is evenly divisible by 5
//
// fizzbuzz if the number is evenly divisible by both
func FizzBuzz(i int64) string {
	s := ""
	if i%3 == 0 {
		s = s + "fizz"
	}
	if i%5 == 0 {
		s = s + "buzz"
	}
	return s
}
