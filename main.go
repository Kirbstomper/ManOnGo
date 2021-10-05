package main

// Returns fizz if the number is evenly divisible by 3
// buzz if the number is evenly divisible by 5
// fizzbuzz if the number is evenly divisible by both
func fizzBuzz(i int) string {
	s := ""
	if i%3 == 0 {
		s = s + "fizz"
	}
	if i%5 == 0 {
		s = s + "buzz"
	}
	return s
}

func main() {

}
