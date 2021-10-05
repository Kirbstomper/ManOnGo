package main

import "testing"

//fizzbuzz should return fizz if evenly divisible by 3
func TestFizzBuzz_returnsFizz(t *testing.T) {
	s := fizzBuzz(3)
	if s != "fizz" {
		t.Fail()
	}

}

//fizzbuzz should return buzz if evenly divisible by 5
func TestFizzBuzz_returnsBuzz(t *testing.T) {
	s := fizzBuzz(5)
	if s != "buzz" {
		t.Fail()
	}
}

//fizzBuzz should return FizzBuzz if evenly divisible by both
func TestFizzBuzz_returnsFizzBuzz(t *testing.T) {
	s := fizzBuzz(15)
	if s != "fizzbuzz" {
		t.Fail()
	}

}
