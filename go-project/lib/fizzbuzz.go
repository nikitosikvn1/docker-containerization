package lib

import "strconv"

func FizzBuzz(query int) string {
	if query%15 == 0 {
		return "fizzbuzz"
	}
	if query%5 == 0 {
		return "buzz"
	}
	if query%3 == 0 {
		return "fizz"
	}
	return strconv.Itoa(query)
}
