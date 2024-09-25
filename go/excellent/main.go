package main

func main() {
	EvenOrOdd(12)
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}
