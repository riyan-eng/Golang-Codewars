package main

import (
	"fmt"
	"math/bits"
)

func CountBits(number uint) int {
	count := 0
	for number != 0 {
		count += int(number) & 1
		number >>= 1
	}
	return count
}

func CountBits2(number uint) int {
	return bits.OnesCount(number)

}

func main() {
	n := 1234
	fmt.Println(CountBits(uint(n)))
	fmt.Println(CountBits2(uint(n)))
}
