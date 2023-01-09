package main

import "fmt"

func CountBits(number uint) int {
	count := 0
	for number != 0 {
		count += int(number) & 1
		number >>= 1
	}
	return count
}

func main() {
	n := 1234
	fmt.Println(CountBits(uint(n)))
}
